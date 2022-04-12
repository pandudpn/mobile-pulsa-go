package telco

// InquiryParam contains data request parameters before customer doing payment
// the response will get `tr_id` for use in Payment
type InquiryParam struct {
	Code     string `json:"code" validate:"required"`
	HP       string `json:"hp" validate:"required"`
	RefID    string `json:"ref_id" validate:"required"`
	Commands string `json:"commands" validate:"required"`
	Username string `json:"username" validate:"required"`
	Sign     string `json:"sign" validate:"required"`
}

// PaymentParam contains data request parameters for payment Telco Postpaid
// TrID is data from response Inquiry
type PaymentParam struct {
	TrID     int    `json:"tr_id" validate:"required"`
	Commands string `json:"commands" validate:"required"`
	Username string `json:"username" validate:"required"`
	Sign     string `json:"sign" validate:"required"`
}
