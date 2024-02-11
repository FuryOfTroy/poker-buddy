package objects

// Card struct definition
type CardRef struct {
	card  *Card
	avail bool
}

// NewCard is a constructor for creating a new Card instance
func NewCardRef(card *Card, avail bool) *CardRef {
	return &CardRef{card: card, avail: avail}
}

// GetValue returns the cardRef's card
func (c *CardRef) GetCard() *Card {
	return c.card
}

// GetSuit returns the card's suit
func (c *CardRef) GetAvail() bool {
	return c.avail
}
