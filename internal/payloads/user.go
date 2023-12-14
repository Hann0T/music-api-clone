package payloads

type LoginRequest struct {
	Name string `json:"name" validate:"required,email"`
}
