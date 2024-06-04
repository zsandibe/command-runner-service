package domain

type CreateCommandInput struct {
	Script      string `json:"script" db:"script"`
	Description string `json:"description" db:"description"`
}

type CreateCommandOutput struct {
	Id          string `json:"id" db:"id"`
	Script      string `json:"script" db:"script"`
	Description string `json:"description" db:"description"`
}
