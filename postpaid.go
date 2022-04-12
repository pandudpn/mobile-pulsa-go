package mobilepulsa

// Inquiry is response data from Postpaid product inquiry
type Inquiry struct {
	Data DataInquiry `json:"data"`
}

// Payment is response data from Postpaid product purchase
type Payment struct {
	Data DataPayment `json:"data"`
}

type DataInquiry struct {
	TrID         int         `json:"tr_id"`
	TrName       string      `json:"tr_name"`
	Code         string      `json:"code"`
	HP           string      `json:"hp"`
	Period       string      `json:"period"`
	Nominal      float64     `json:"nominal"`
	Admin        float64     `json:"admin"`
	RefID        string      `json:"ref_id"`
	ResponseCode string      `json:"response_code"`
	Message      string      `json:"message"`
	Price        float64     `json:"price"`
	SellingPrice float64     `json:"selling_price"`
	Description  Description `json:"desc"`
}

type DataPayment struct {
	TrID         int         `json:"tr_id"`
	Code         string      `json:"code"`
	Datetime     string      `json:"datetime"`
	HP           string      `json:"hp"`
	TrName       string      `json:"tr_name"`
	Period       string      `json:"period"`
	Nominal      float64     `json:"nominal"`
	Admin        float64     `json:"admin"`
	ResponseCode string      `json:"response_code"`
	Message      string      `json:"message"`
	Price        float64     `json:"price"`
	SellingPrice float64     `json:"selling_price"`
	Balance      float64     `json:"balance"`
	NoRef        string      `json:"noref"`
	RefID        string      `json:"ref_id"`
	Description  Description `json:"desc"`
}

// Description is description charge for customer inquiry
type Description struct {
	KodeArea      string  `json:"kode_area"`
	Divre         string  `json:"divre"`
	Datel         string  `json:"datel"`
	JumlahTagihan int     `json:"jumlah_tagihan"`
	Tagihan       Tagihan `json:"tagihan"`
}

type Tagihan struct {
	Details []DetailTagihan `json:"detail"`
}

// DetailTagihan is data for customer charge monthly
type DetailTagihan struct {
	Periode      string      `json:"periode"`
	NilaiTagihan string      `json:"nilai_tagihan"`
	Admin        string      `json:"admin"`
	Total        interface{} `json:"total"` // can be string or integer or double (float)
}
