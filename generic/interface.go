package generic

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type PlayingCard struct {
	Suit string
	Rank string
}

type Deck1 struct {
	cards []interface{}
}

func newPlayingCard(suit string, card string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: card}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

func (d *Deck1) addCard(card interface{}) {
	d.cards = append(d.cards, card)
}

func (d *Deck1) randomCard() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

func newPlayingCardDeck1() *Deck1 {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck1{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.addCard(newPlayingCard(suit, rank))
		}
	}
	return deck
}

func NonGenericInterface() {
	deck := newPlayingCardDeck1()

	fmt.Printf("--- drawing playing card ---\n")
	card := deck.randomCard()
	fmt.Printf("drew card: %s\n", card)

	playingCard, ok := card.(*PlayingCard)
	if !ok {
		fmt.Printf("card received wasn't a playing card!")
		os.Exit(1)
	}
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)
}
