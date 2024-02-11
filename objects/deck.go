package objects

import (
	"fmt"
	"strings"
)

// GetCardIndex returns the index of the card in the deck
func GetCardIndex(value int, suit int) int {
	return (value-2)*4 + (suit - 1)
}

// Deck struct definition
type Deck struct {
	cardRefs []*CardRef
}

// NewDeck is a constructor for creating a new Deck instance
func NewDeck() *Deck {
	deck := &Deck{
		cardRefs: make([]*CardRef, 0),
	}
	for value := 2; value <= 14; value++ {
		for suit := 1; suit <= 4; suit++ {
			card := NewCard(value, suit)
			cardRef := NewCardRef(card, true)
			deck.cardRefs = append(deck.cardRefs, cardRef)
		}
	}
	return deck
}

// Take removes and returns a card from the deck
func (d *Deck) Take(value int, suit int) *Card {
	return d.TakeIndex(GetCardIndex(value, suit))
}

// Take removes and returns a card from the deck
func (d *Deck) TakeIndex(cardIndex int) *Card {
	cardRef := d.cardRefs[cardIndex]
	if cardRef.avail {
		cardRef.avail = false
		return cardRef.card
	}
	panic(fmt.Errorf("Card not found"))
}

// Take removes and returns a card from the deck
func (d *Deck) TryTake(value int, suit int) *Card {
	return d.TryTakeIndex(GetCardIndex(value, suit))
}

// Take removes and returns a card from the deck
func (d *Deck) TryTakeIndex(cardIndex int) *Card {
	cardRef := d.cardRefs[cardIndex]
	if cardRef.avail {
		cardRef.avail = false
		return cardRef.card
	}
	return nil
}

// Returns cards to the deck
func (d *Deck) ReturnAll(cards []*Card) {
	for _, card := range cards {
		d.Return(card)
	}
}

// Returns a card to the deck
func (d *Deck) Return(card *Card) {
	cardIndex := GetCardIndex(card.value, card.suit)
	cardRef := d.cardRefs[cardIndex]
	if cardRef.avail == false {
		cardRef.avail = true
		return
	}
	panic(fmt.Errorf("Card %s already in deck", card.Print()))
}

// Clone the deck
func (d *Deck) Clone() *Deck {
	cardRefs := make([]*CardRef, 0)
	for _, cardRef := range d.cardRefs {
		cardRefs = append(cardRefs, &CardRef{card: cardRef.GetCard(), avail: cardRef.GetAvail()})
	}
	return &Deck{cardRefs: cardRefs}
}

// Print prints all cards in the deck
func (d *Deck) Print() string {
	var b strings.Builder
	for _, cardRef := range d.cardRefs {
		if cardRef.avail {
			b.WriteString(cardRef.card.Print())
			b.WriteString("\n")
		}
	}

	return b.String()
}
