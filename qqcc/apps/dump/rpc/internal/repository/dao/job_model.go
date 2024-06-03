package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木

type ParseJob struct {
	gorm.Model
	FileName   string  `json:"file_name" gorm:"type:varchar(128);comment:文件名"`
	FilePath   string  `json:"file_path" gorm:"uniqueIndex;type:varchar(128);comment:文件路径,唯一索引"`
	FileSizeMB float64 `json:"file_size_mb" gorm:"type:float;comment:文件大小"`
	Resolved   bool    `json:"resolved" gorm:"comment:是否解析过"`
	Mark       string  `json:"mark" gorm:"type:varchar(128);comment:备注"`
	ParseSec   float64 `json:"parse_sec" gorm:"type:float;comment:解析时长"`
	Rows       int     `json:"rows" gorm:"type:int;comment:行数"`
}

func (p ParseJob) TableName() string {
	return "qqcc_job"
}
