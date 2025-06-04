package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newdeck() deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var d deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newdeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error Occured :", err)
		os.Exit(1) // Exit the program if there is an error reading the file
	}
	return deck(strings.Split(string(bs), ", "))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	se := rand.New(source)
	for i := range d {
		newPosition := se.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
