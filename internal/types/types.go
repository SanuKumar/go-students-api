package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" validate:"required"` // used with github.com/go-playground/validator/v10
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}
