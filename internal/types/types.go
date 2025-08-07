package types

type Student struct {
	Id    int64
	Name  string `validate:"required"` // used with github.com/go-playground/validator/v10
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}
