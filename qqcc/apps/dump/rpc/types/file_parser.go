package types

// @Description
// @Author 代码小学生王木木

type FileParserMsg struct {
	ID       int64  `json:"id"`
	FileName string `json:"file_name" gorm:"type:varchar(128);comment:文件名"`
	FilePath string `json:"file_path" gorm:"uniqueIndex;type:varchar(128);comment:文件路径,唯一索引"`
}
