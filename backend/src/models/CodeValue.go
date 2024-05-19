package models

type CodeValue struct {
	ID            int    `json:"id" gorm:"uniqueIndex:type_id_code_value_id"`
	CodeTypeID    int    `json:"code_type_id" gorm:"uniqueIndex:type_id_code_value_id"`
	CodeValue     string `json:"code_value"`
	CodeValueDesc string `json:"code_value_desc"`
}
