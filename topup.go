package mobilepulsa

// TopUp contains data from response TopUp Prepaid Product
type TopUp struct {
	Data DataTopUp `json:"data"`
}

type DataTopUp struct {
	RefID       string  `json:"ref_id"`
	Status      float64 `json:"status"`
	ProductCode string  `json:"product_code"`
	Price       float64 `json:"price"`
	Message     string  `json:"message"`
	TrID        int     `json:"tr_id"`
	RC          string  `json:"rc"`
}
