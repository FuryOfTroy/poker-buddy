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

func (c *Card) GetValueString() string {
	return ValueToString(c.value)
}

func (c *Card) GetSuitString() string {
	return SuitToString(c.suit)
}

// ValueToString returns the string representation of a card's value
func ValueToString(value int) string {
	if value < 0 || value > 14 {
		panic(fmt.Errorf("invalid value %d", value))
	}
	switch value {
	case 14:
		return "A"
	case 13:
		return "K"
	case 12:
		return "Q"
	case 11:
		return "J"
	case 10:
		return "T"
	default:
		return strconv.Itoa(value)
	}
}

// StrToValue returns a card's value from the string representation provided
func StringToValue(valueStr string) int {
	switch valueStr {
	case "A", "a":
		return 14
	case "K", "k":
		return 13
	case "Q", "q":
		return 12
	case "J", "j":
		return 11
	case "T", "t":
		return 10
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(err)
	}

	if value < 0 || value > 14 {
		panic(fmt.Errorf("invalid value string %s", valueStr))
	}

	return value
}

// SuitToString returns the fancy string representation of a card's suit
func SuitToString(suit int) string {
	switch suit {
	case 1:
		return "♥️"
	case 2:
		return "♦️"
	case 3:
		return "♣️"
	case 4:
		return "♠️"
	}
	panic(fmt.Errorf("invalid suit %d", suit))
}

// StringToSuit returns a card's suit from the basic or fancy string representation provided
func StringToSuit(suitStr string) int {
	switch suitStr {
	case "♥️", "h":
		return 1
	case "♦️", "d":
		return 2
	case "♣️", "c":
		return 3
	case "♠️", "s":
		return 4
	}
	panic(fmt.Errorf("invalid suit string %s", suitStr))
}

// SuitToStringBasic returns the basic string representation of a card's suit
func SuitToStringBasic(suit int) string {
	switch suit {
	case 1:
		return "h"
	case 2:
		return "d"
	case 3:
		return "c"
	case 4:
		return "s"
	}
	panic(fmt.Errorf("invalid suit %d", suit))
}

// Print returns the string representation of the card
func (c *Card) String() string {
	return ValueToString(c.value) + SuitToString(c.suit)
}
