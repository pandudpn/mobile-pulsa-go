package pricelist

type (
	Service int
	Type    int
)

const (
	Prepaid Service = iota
	Postpaid
)

const (
	Data Type = iota
	EToll
	Game
	PLN
	Pulsa
	Voucher
	PDAM
	BPJS
	Internet
	PajakKendaraan
	Finance
	HP
	Estate
	EMoney
	Kereta
	TV
	Airlane
	O2O
	PBB
	Gas
	PajakDaerah
	Pasar
	Retribusi
	Pendidikan
	Asuransi
)

// priceListParam contains parameters for getting Product PriceList
type priceListParam struct {
	Commands string  `json:"commands,omitempty"`
	Status   string  `json:"status" validate:"required"`
	Username string  `json:"username" validate:"required"`
	Sign     string  `json:"sign" validate:"required"`
	Service  Service `json:"-"`
	Type     Type    `json:"-"`
}
