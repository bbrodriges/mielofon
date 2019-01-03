package dialog

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
	Type        CardType `json:"type,omitempty"`
	ImageID     string   `json:"image_id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Button      *Button  `json:"button,omitempty"`
}

func (c BigImageCard) CardType() CardType {
	return TypeBigImageCard
}

type ItemsListCard struct {
	Type   CardType    `json:"type,omitempty"`
	Header *CardHeader `json:"header,omitempty"`
	Items  []CardItem  `json:"items,omitempty"`
	Footer *CardFooter `json:"footer,omitempty"`
}

func (c ItemsListCard) CardType() CardType {
	return TypeItemsListCard
}

type CardHeader struct {
	Text string `json:"text,omitempty"`
}

type CardFooter struct {
	Text   string  `json:"text,omitempty"`
	Button *Button `json:"button,omitempty"`
}

type CardItem struct {
	ImageID     string  `json:"image_id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Button      *Button `json:"button,omitempty"`
}
