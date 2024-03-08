package objects

import (
	"fmt"
	"strings"
)

// GetCardIndex returns the deck index of the card with the provided value and suit
func GetIndex(card *Card) int {
	return GetCardIndex(card.GetValue(), card.GetSuit())
}

// GetCardIndex returns the deck index of the card with the provided value and suit
func GetCardIndex(value int, suit int) int {
	return (value-2)*4 + (suit - 1)
}

// GetCardIndex returns the deck index of the card with the provided value and suit
func GetCardIndexFromName(cardName string) int {
	return GetCardIndex(StringToValue(cardName[0:1]), StringToSuit(cardName[1:2]))
}

// Card struct definition
type cardRef struct {
	card  *Card
	avail bool
}

// NewCard is a constructor for creating a new Card instance
func NewCardRef(card *Card, avail bool) *cardRef {
	return &cardRef{card: card, avail: avail}
}

// GetCard returns the cardRef's card
func (c *cardRef) GetCard() *Card {
	return c.card
}

// GetAvail returns whether the cardRef's card is in the deck
func (c *cardRef) GetAvail() bool {
	return c.avail
}

// Deck struct definition
type Deck struct {
	cardRefs []*cardRef
}

// NewDeck is a constructor for creating a new Deck instance
func NewDeck() *Deck {
	deck := &Deck{
		cardRefs: make([]*cardRef, 0),
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

// Take removes a card from the deck with the value and suit provided
// Panics if the card isn't in the deck
// Returns the card if found
func (d *Deck) Take(value int, suit int) *Card {
	return d.TakeIndex(GetCardIndex(value, suit))
}

// TakeIndex removes a card from the deck with the index provided
// Panics if the card isn't in the deck
// Returns the card if found
func (d *Deck) TakeIndex(cardIndex int) *Card {
	cardRef := d.cardRefs[cardIndex]
	if cardRef.avail {
		cardRef.avail = false
		return cardRef.card
	}
	panic(fmt.Errorf("Card not found"))
}

// TakeName removes a card from the deck with the name provided
// Panics if the card isn't in the deck
// Returns the card if found
func (d *Deck) TakeName(cardName string) *Card {
	return d.TakeIndex(GetCardIndexFromName(cardName))
}

// TryTake removes a card from the deck with the index provided
// Returns nil if the card isn't in the deck
// Returns the card if found
func (d *Deck) TryTake(value int, suit int) *Card {
	return d.TryTakeIndex(GetCardIndex(value, suit))
}

// TryTakeIndex removes a card from the deck with the value and suit provided
// Returns nil if the card isn't in the deck
// Returns the card if found
func (d *Deck) TryTakeIndex(cardIndex int) *Card {
	cardRef := d.cardRefs[cardIndex]
	if cardRef.avail {
		cardRef.avail = false
		return cardRef.card
	}
	return nil
}

// ReturnAll returns all provided cards to the deck
// Panics if any of the cards are already in the deck
func (d *Deck) ReturnAll(cards []*Card) {
	for _, card := range cards {
		d.Return(card)
	}
}

// Return returns provided card to the deck
// Panics if the card is already in the deck
func (d *Deck) Return(card *Card) {
	cardIndex := GetCardIndex(card.value, card.suit)
	cardRef := d.cardRefs[cardIndex]
	if !cardRef.avail {
		cardRef.avail = true
		return
	}
	panic(fmt.Errorf("Card %s already in deck", card))
}

// Clone clones the deck and and each of it's CardRefs, but not the Cards themselves
// Returns the deck clone
func (d *Deck) Clone() *Deck {
	cardRefs := make([]*cardRef, 0)
	for _, cr := range d.cardRefs {
		cardRefs = append(cardRefs, &cardRef{card: cr.GetCard(), avail: cr.GetAvail()})
	}
	return &Deck{cardRefs: cardRefs}
}

// Print prints all cards in the deck
func (d *Deck) String() string {
	var b strings.Builder
	for _, cardRef := range d.cardRefs {
		if cardRef.avail {
			b.WriteString(cardRef.card.String())
			b.WriteString("\n")
		}
	}

	return b.String()
}
