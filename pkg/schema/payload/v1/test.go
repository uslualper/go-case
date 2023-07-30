package v1

type Test struct {
	Message string `json:"message" validate:"required"`
}
