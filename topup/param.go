package topup

// TopUpParam contains data parameters to request TopUp Prepaid Products
type TopUpParam struct {
	Username    string `json:"username" validate:"required"`
	RefID       string `json:"ref_id" validate:"required"`
	CustomerID  string `json:"customer_id" validate:"required"`
	ProductCode string `json:"product_code" validate:"required"`
	Sign        string `json:"sign" validate:"required"`
}
