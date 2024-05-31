package logic

import (
	"context"
	"encoding/json"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/threading"
	"os"
	"qqcc/apps/dump/rpc/domain"
	"strings"
	"time"

	"qqcc/apps/dump/rpc/internal/svc"
	"qqcc/apps/dump/rpc/types/dump"

	"github.com/zeromicro/go-zero/core/logx"
)

/**
https://pwmzlkcu3p.feishu.cn/docx/O1u3d9pqWo4sZTx6NTxcOHhknmd
*/

type ParserExcelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewParserExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParserExcelLogic {
	return &ParserExcelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ParserExcelLogic) ParserExcel(in *dump.ParserRequest) (*dump.ParserResponse, error) {
	filePath := in.GetFilepath()
	_, err := os.Stat(filePath)
	if err != nil {
		return &dump.ParserResponse{}, err
	}

	// 1. 解析Excel
	f, err := excelize.OpenFile(filePath)
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			return
		}
	}()
	if err != nil {
		return &dump.ParserResponse{}, err
	}
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		//p.lg.Error("获取工作表的行信息失败", logger.Error(err), logger.String("文件路径：", filePath))
		return &dump.ParserResponse{}, err
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
		// 2. 写入kafka消息队列
		threading.GoSafe(func() {
			company := l.parserTitle(titleMap, row)
			data, err := json.Marshal(company)
			if err != nil {
				l.Logger.Errorf("[Kafka Dump] 序列化消息失败: %+v error: %v", company, err)
				return
			}
			err = l.svcCtx.KqPusherClient.Push(string(data))
			if err != nil {
				//	能怎么办？只能记录日志啦，还能怎么办
				l.Logger.Errorf("[Kafka Dump] kafka发送消息失败: %s error: %v", data, err)
			}
		})
	}
	return &dump.ParserResponse{}, nil
}

func (p *ParserExcelLogic) parserTitle(titleMap map[string]int, row []string) domain.EnterpriseBasicDM {
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

func (l *ParserExcelLogic) parseDate(str string) time.Time {
	if len(str) == 0 {
		return time.Time{}
	}
	parse, err := time.Parse("2024-05-10", str)
	if err != nil {
		return time.Time{}
	}
	return parse
}
func (l *ParserExcelLogic) cleanTitle(str string) string {
	if len(str) == 0 || strings.TrimSpace(str) == "-" {
		return ""
	}
	return str
}
