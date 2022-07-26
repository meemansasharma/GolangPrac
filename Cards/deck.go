package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func newDeck() deck {

	cards := deck{}

	cardSuite := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNumber := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suite := range cardSuite {
		for _, number := range cardNumber {
			cards = append(cards, number+" of "+suite)
		}
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fileString := strings.Split(string(file), ",")
	return deck(fileString)
}

func (d deck) shuffleDeck() {

	//rand depends on the seed value inside rand; rand not giving randon value due to fixed seed
	//now generate new seed
	//generate new int64 number each time
	source := rand.NewSource(time.Now().UnixNano())
	//get rand with new seed vaue
	r := rand.New(source)
	for i := range d {
		newIndex := r.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
