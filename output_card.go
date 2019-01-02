package mielofon

type CardType string

const (
	TypeBigImageCard  CardType = "BigImage"
	TypeItemsListCard CardType = "ItemsList"
)

type OutputCard interface {
	CardType() CardType
}

var (
	_ OutputCard = BigImageCard{}
	_ OutputCard = ItemsListCard{}
)

type BigImageCard struct {
	Type        CardType `json:"type"`
	ImageID     string   `json:"image_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Button      Button   `json:"button"`
}

func (c BigImageCard) CardType() CardType {
	return TypeBigImageCard
}

type ItemsListCard struct {
	Type   CardType   `json:"type"`
	Header CardHeader `json:"header"`
	Items  []CardItem `json:"items"`
	Footer CardFooter `json:"footer"`
}

func (c ItemsListCard) CardType() CardType {
	return TypeItemsListCard
}

type CardHeader struct {
	Text string `json:"text"`
}

type CardFooter struct {
	Text   string `json:"text"`
	Button Button `json:"button"`
}

type CardItem struct {
	ImageID     string `json:"image_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Button      Button `json:"button"`
}
