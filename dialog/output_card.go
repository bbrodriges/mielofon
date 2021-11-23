package dialog

type CardType string

const (
	CardBigImage     CardType = "BigImage"
	CardItemsList    CardType = "ItemsList"
	CardImageGallery CardType = "ImageGallery"
)

type OutputCard interface {
	Type() CardType
}

type BigImageCard struct {
	CardType    CardType `json:"type,omitempty"`
	ImageID     string   `json:"image_id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Button      *Button  `json:"button,omitempty"`
}

func (BigImageCard) Type() CardType {
	return CardBigImage
}

type ItemsListCard struct {
	CardType CardType    `json:"type,omitempty"`
	Header   *CardHeader `json:"header,omitempty"`
	Items    []CardItem  `json:"items,omitempty"`
	Footer   *CardFooter `json:"footer,omitempty"`
}

func (ItemsListCard) Type() CardType {
	return CardItemsList
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

type ImageGalleryCard struct {
	CardType CardType           `json:"type,omitempty"`
	Items    []ImageGalleryItem `json:"items,omitempty"`
}

func (ImageGalleryCard) Type() CardType {
	return CardImageGallery
}

type ImageGalleryItem struct {
	ImageID string  `json:"image_id,omitempty"`
	Title   string  `json:"title,omitempty"`
	Button  *Button `json:"button,omitempty"`
}
