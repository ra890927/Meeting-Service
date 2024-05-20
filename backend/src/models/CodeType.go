package models

type CodeType struct {
	ID         int         `json:"id"`
	TypeName   string      `json:"type_name" gorm:"unique"`
	TypeDesc   string      `json:"type_desc"`
	CodeValues []CodeValue `json:"code_values"`
}
