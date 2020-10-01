package models

type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResponse struct {
	Status      bool         `json:"status"`
	Validations []Validation `json:"validations"`
}
