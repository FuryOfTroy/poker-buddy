package objects

import (
	"fmt"
	"strconv"
)

// Card struct definition
type Card struct {
	value int
	suit  int
}

// NewCard is a constructor for creating a new Card instance
func NewCard(value int, suit int) *Card {
	if value < 2 || value > 14 {
		panic(fmt.Errorf("value not in range 2-14"))
	}
	if suit < 1 || suit > 4 {
		panic(fmt.Errorf("suit not in range 1-4"))
	}
	return &Card{value: value, suit: suit}
}

func (c *Card) GetValue() int {
	return c.value
}

func (c *Card) GetSuit() int {
	return c.suit
}

// GetValueChar returns the character representation of the card's value
func (c *Card) GetValueChar() string {
	switch c.value {
	case 14:
		return "A"
	case 13:
		return "K"
	case 12:
		return "Q"
	case 11:
		return "J"
	default:
		return strconv.Itoa(c.value)
	}
}

// GetSuitChar returns the character representation of the card's suit
func (c *Card) GetSuitChar() string {
	switch c.suit {
	case 1:
		return "♥️"
	case 2:
		return "♦️"
	case 3:
		return "♣️"
	case 4:
		return "♠️"
	}
	return ""
}

// Print returns the string representation of the card
func (c *Card) Print() string {
	return c.GetValueChar() + c.GetSuitChar()
}
