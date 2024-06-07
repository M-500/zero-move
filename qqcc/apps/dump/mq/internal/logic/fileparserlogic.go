package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"os"
	"qqcc/apps/dump/domain"
	"qqcc/apps/dump/mq/internal/svc"
	"qqcc/apps/dump/mq/internal/types"
	"qqcc/apps/dump/rpc/types/dump"
	"strings"
	"time"
)

type FileParserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileParserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileParserLogic {

	return &FileParserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileParserLogic) Consume(_, val string) error {
	var msg *types.FileParserMsg
	err := json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}
	return l.ParserFile(l.ctx, msg)
}

func (l *FileParserLogic) ParserFile(ctx context.Context, msg *types.FileParserMsg) error {
	// 1.去数据库查询是否 重复消费了
	res, err := l.svcCtx.DumpRpc.FindParserJobById(ctx, &dump.FindParserJonRequest{Id: msg.ID})
	if err != nil {
		l.Logger.Errorf("调用Dump服务失败", err)
		return err
	}
	if res.Resolved {
		return errors.New("文件已经被解析过了！")
	}
	// 2. 开始解析，并插入数据库中
	_, err = os.Stat(msg.FilePath)
	if os.IsNotExist(err) {
		return errors.New("文件路径不存在")
	}
	f, err := excelize.OpenFile(msg.FilePath)
	sheetName := f.GetSheetName(0)
	// 获取第一个工作表的行数和列数
	rows, err := f.GetRows(sheetName)
	if err != nil {
		logx.Errorf("获取工作表的行信息失败", err)
		return err
	}
	if err != nil {
		return err
	}
	titleMap := make(map[string]int, 20)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			// 第一次进来遍历第一行，就可以获取map表结构
			for ik, cell := range row {
				titleMap[cell] = ik
			}
			continue
		}
		// 解析出来的公司信息
		company := l.parserTitle(titleMap, row)
		fmt.Println(company)
	}
	return nil
}

func (l *FileParserLogic) parseDate(str string) time.Time {
	if len(str) == 0 {
		return time.Time{}
	}
	parse, err := time.Parse("2024-05-10", str)
	if err != nil {
		return time.Time{}
	}
	return parse
}
func (l *FileParserLogic) cleanTitle(str string) string {
	if len(str) == 0 || strings.TrimSpace(str) == "-" {
		return ""
	}
	return str
}
func (p *FileParserLogic) parserTitle(titleMap map[string]int, row []string) domain.EnterpriseBasicDM {
	company := domain.EnterpriseBasicDM{}
	val, ok := titleMap["企业名称"]
	if ok {
		company.CompanyName = p.cleanTitle(row[val])
	}
	val, ok = titleMap["登记状态"]
	if ok {
		company.RegStatus = p.cleanTitle(row[val])
	}
	val, ok = titleMap["法定代表人"]
	if ok {
		company.LegalBoss = p.cleanTitle(row[val])
	}
	val, ok = titleMap["注册资本"]
	if ok {
		company.RegCapital = p.cleanTitle(row[val])
	}
	val, ok = titleMap["所属省份"]
	if ok {
		company.Province = p.cleanTitle(row[val])
	}
	val, ok = titleMap["所属城市"]
	if ok {
		company.City = p.cleanTitle(row[val])
	}
	val, ok = titleMap["所属区县"]
	if ok {
		company.District = p.cleanTitle(row[val])
	}
	val, ok = titleMap["成立日期"]
	if ok {
		company.RegDate = p.parseDate(p.cleanTitle(row[val]))
	}
	val, ok = titleMap["统一社会信用代码"]
	if ok {
		company.SocialCode = p.cleanTitle(row[val])
	}
	val, ok = titleMap["企业注册地址"]
	if ok {
		company.CompanyRegAddr = p.cleanTitle(row[val])
	}
	val, ok = titleMap["电话"]
	if ok {
		company.Phone = p.cleanTitle(row[val])
	}
	val, ok = titleMap["更多电话"]
	//if ok {
	//	company.MorePhone = row[val]
	//}
	val, ok = titleMap["邮箱"]
	if ok {
		company.Email = p.cleanTitle(row[val])
	}
	val, ok = titleMap["更多邮箱"]
	//if ok {
	//	company.MoreEmail = row[val]
	//}
	return company
}
func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewFileParserLogic(ctx, svcCtx)),
	}
}
