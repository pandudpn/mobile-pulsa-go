package mobilepulsa

// PriceList contains data from MobilePulsa API response of getting price list request
type PriceList struct {
	Data DataPricelist `json:"data"`
}

type DataPricelist struct {
	Pasca     []DataPostpaid `json:"pasca,omitempty"`     // postpaid
	PriceList []DataPrepaid  `json:"pricelist,omitempty"` // prepaid
}

// DataPrepaid contains data from MobilePulsa API Prepaid response
type DataPrepaid struct {
	ProductCode        string `json:"product_code"`
	ProductDescription string `json:"product_description"`
	ProductNominal     string `json:"product_nominal"`
	ProductDetails     string `json:"product_details"`
	ProductPrice       int    `json:"product_price"`
	ProductType        string `json:"product_type"`
	ActivePeriod       string `json:"active_period"`
	Status             string `json:"status"`
	IconUrl            string `json:"icon_url"`
}

// DataPostpaid contains data from MobilePulsa API Postpaid response
type DataPostpaid struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Fee    int    `json:"fee"`
	Komisi int    `json:"komisi"`
	Type   string `json:"type"`
}
