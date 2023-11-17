package generic

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck2[C any] struct {
	cards []C
}

func (d *Deck2[C]) addCard(card C) {
	d.cards = append(d.cards, card)
}

func (d *Deck2[C]) randomCard() C {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

func newPlayingCardDeck2() *Deck2[*PlayingCard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck2[*PlayingCard]{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.addCard(newPlayingCard(suit, rank))
		}
	}
	return deck
}

func GenericAny() {
	deck := newPlayingCardDeck2()

	fmt.Printf("--- drawing playing card ---\n")
	playingCard := deck.randomCard()
	fmt.Printf("drew card: %s\n", playingCard)
	// Code removed
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)
}
