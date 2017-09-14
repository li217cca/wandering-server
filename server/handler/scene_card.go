package handler

import (
	"fmt"
	"errors"
)

type Card struct {
	ID 		int `json:"id"`
	Name     string `json:"name"`
	TachieID int    `json:"tachie_id"`
	Type     string `json:"type"`
	Rare     int    `json:"rare"`

	OnHandle func()
}

type Scene struct {
	cards []Card
	cardIDPos int
}

func (s *Scene) GetScene() map[string]interface{} {
	return map[string]interface{}{
		"cards": s.GetCards(),
	}
}

func (s *Scene) GetCard(ID int) (Card, error) {
	for _, card := range s.cards {
		if card.ID == ID {
			return card, error(nil)
		}
	}
	return Card{}, errors.New(fmt.Sprintf("Card ID = %d not found", ID))
}

func (s *Scene) GetCards() []Card {
	return s.cards
}

func (s *Scene) GetLen() int {
	return len(s.cards)
}

func (s *Scene) AddCards(cards ...Card) {
	for _, card := range cards {
		s.cardIDPos ++
		s.cards = append(s.cards, Card{
			s.cardIDPos,
			card.Name,
			card.TachieID,
			card.Type,
			card.Rare,
			card.OnHandle,
		})
	}
}
func (s *Scene) RemoveCard(ID int) {
	for index, card := range s.cards {
		if card.ID == ID {
			s.cards = append(s.cards[:index], s.cards[index+1:]...)
		}
	}
}