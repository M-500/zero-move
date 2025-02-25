package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"os"
	"qqcc/apps/dump/rpc/internal/repository/dao"
	"qqcc/apps/dump/rpc/internal/svc"
	"qqcc/apps/dump/rpc/types"
	"qqcc/apps/dump/rpc/types/dump"

	"github.com/zeromicro/go-zero/core/logx"
)

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

// ParserExcel
//
//	@Description: 将文件解析任务传递到kafka，并保存到数据库中
//	@receiver l
//	@param in
//	@return *dump.ParserResponse
//	@return error
func (l *ParserExcelLogic) ParserExcel(in *dump.ParserRequest) (*dump.ParserResponse, error) {
	filePath := in.GetFilepath()
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return &dump.ParserResponse{}, err
	}

	data := dao.ParseJob{
		FileName:   in.Filename,
		FilePath:   in.Filepath,
		FileSizeMB: float64(fileInfo.Size()) / (1024 * 1024),
		Resolved:   false,
		Mark:       "",
		ParseSec:   0,
		Rows:       0,
	}
	// 开启事务，保证消费者一定发送成功，防止重复消费
	//l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
	//
	//}
	// 保存到数据库
	id, err := l.svcCtx.ParserDAO.Insert(l.ctx, data)
	if err != nil {
		return nil, err
	}

	fileMsg := types.FileParserMsg{
		ID:       id,
		FileName: in.Filename,
		FilePath: in.Filepath,
	}
	// 异步插入kafka中
	threading.GoSafe(func() {
		data1, err1 := json.Marshal(fileMsg)
		if err1 != nil {
			l.Logger.Errorf("[Kafka Dump] 序列化消息失败: %+v error: %v", data1, err1)
			return
		}
		err = l.svcCtx.KqPusherClient.Push(string(data1))
		if err != nil {
			// 能怎么办？只能记录日志啦，还能怎么办 重试？ 重试多少次呢？还是说 改成同步，使用本地事务保证一定发送消息成功？
			// 重试 -> 导致重复消费哦！
			l.Logger.Errorf("[Kafka Dump] kafka发送消息失败: %s error: %v", data, err)
		}
	})
	return &dump.ParserResponse{}, nil
}
