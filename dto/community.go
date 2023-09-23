package dto

type CreateCommunity struct {
	Id   string `validate:"required,min=2"`
	Name string `validate:"required,min=2"`
}

type UpdateCommunity struct {
	Name string `validate:"required,min=2"`
}
