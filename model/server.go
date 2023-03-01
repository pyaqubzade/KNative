package model

type ReportRequest struct {
	VIN         string `json:"vin" validate:"required"`
	Email       string `json:"email" validate:"email"`
	PhoneNumber string `json:"phoneNumber" validate:"e164"`
}

type Data struct {
	Value string `json:"value"`
}
