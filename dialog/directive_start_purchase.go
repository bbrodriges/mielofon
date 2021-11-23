package dialog

type PurchaseCurrency string

const (
	PurchaseCurrencyRUB PurchaseCurrency = "RUB"
)

type PurchaseAction string

const (
	PurchaseActionBuy PurchaseAction = "BUY"
)

type ProductNDSType string

const (
	PurchaseNDSNone ProductNDSType = "nds_none"
	PurchaseNDS0    ProductNDSType = "nds_0"
	PurchaseNDS10   ProductNDSType = "nds_10"
	PurchaseNDS20   ProductNDSType = "nds_20"
)

type StartPurchaseDirective struct {
	PurchaseRequestId string            `json:"purchase_request_id"`
	ImageURL          string            `json:"image_url"`
	Caption           string            `json:"caption"`
	Description       string            `json:"description"`
	Currency          PurchaseCurrency  `json:"currency"`
	ActionType        PurchaseAction    `json:"type"`
	Payload           interface{}       `json:"payload"`
	MerchantKey       string            `json:"merchant_key"`
	TestPayment       bool              `json:"test_payment"`
	Products          []PurchaseProduct `json:"products,omitempty"`
}

func (StartPurchaseDirective) Type() DirectiveType {
	return DirectiveStartPurchase
}

type PurchaseProduct struct {
	ProductId string         `json:"product_id"`
	Title     string         `json:"title"`
	UserPrice string         `json:"user_price"`
	Price     string         `json:"price"`
	Quantity  string         `json:"quantity"`
	NdsType   ProductNDSType `json:"nds_type"`
}
