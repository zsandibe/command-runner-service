package entity

type Command struct {
	Id          int64  `json:"id" db:"id"`
	Script      string `json:"script" db:"script"`
	Description string `json:"description,omitempty" db:"description"`
}
