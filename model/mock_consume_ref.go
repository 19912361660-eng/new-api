package model

import "gorm.io/gorm"

// MockConsumeRef 模拟日志关联标记表
// 只用来记录哪些consume是管理员模拟生成，原生代码完全不读取这张表
type MockConsumeRef struct {
	gorm.Model
	ConsumeID uint64 `json:"consume_id" gorm:"column:consume_id;index;comment:关联原生消费日志ID"`
	BatchID   string `json:"batch_id" gorm:"column:batch_id;index;comment:批次UUID，同一批模拟日志，用来批量删除"`
	MockNote  string `json:"mock_note" gorm:"column:mock_note;type:text;comment:管理员内部备注，不对外展示"`
}

// TableName 指定表名
func (MockConsumeRef) TableName() string {
	return "mock_consume_ref"
}
