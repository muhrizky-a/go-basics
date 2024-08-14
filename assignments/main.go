package main

import (
	"fmt"
	"sort"
	"math"
	"math/rand/v2"
)

func main() {
	fmt.Println("hello world")

	arraySign([]int{2, 1})                    // 1
	arraySign([]int{-2, 1})                   // -1
	arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) // 1
	
	isAnagram("anak", "kana") // true
	isAnagram("anak", "mana") // false
	isAnagram("anagram", "managra") // true

	findTheDifference("abcd", "abcde") // 'e'
	findTheDifference("abcd", "abced") // 'e'
	findTheDifference("", "y")         // 'y'

	canMakeArithmeticProgression([]int{1, 5, 3})    // true; 1, 3, 5 adalah baris aritmatik +2
	canMakeArithmeticProgression([]int{5, 1, 9})   // true; 9, 5, 1 adalah baris aritmatik -4
	canMakeArithmeticProgression([]int{1, 2, 4, 8})// false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code here
	product := 1
	for _, value := range nums {
		product *= value
	}

	if product > 0 {
		return 1 // if positive
	} else if product < 0 {
		return -1 // if negative
	}
	return 0
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// write code here
	if len(s) != len(t) {
        return false;
    }

	byteS := []byte(s)
	byteT := []byte(t)

	sort.Slice(byteS, func( i, j int) bool {
		return byteS[i] < byteS[j]
	})

	sort.Slice(byteT, func( i, j int) bool {
		return byteT[i] < byteT[j]
	})
	
	for i := range byteS {
		if byteS[i] != byteT[i] {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	// write code here
	if len(t) != len(s) + 1 {
		return 0
	}

	byteT := []byte(t)

	sort.Slice(byteT, func( i, j int) bool {
		return byteT[i] < byteT[j]
	})
	
	return byteT[len(t) -1]
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// write code here
	if len(arr) < 2 {
		return false
	}

	sort.Slice(arr, func( i, j int) bool {
		return arr[i] < arr[j]
	})

	diff := 0.0
	for i := range arr {
		if(i == 0){
			diff = math.Abs(float64(arr[1]) - float64(arr[0]))
		} else {
			if math.Abs(float64(arr[i]) - float64(arr[i - 1])) != diff {
				return false
			}
		}
	}
	
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
	newCards := []Card{}
	for symbol := range 4 {
		for number := range 13 {
			newCards = append(
				newCards,
					Card{
					symbol: symbol,
					number: number + 1,
				},
			)
		}
	}

	d.cards = newCards
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here

	cards := d.cards[ : n ]
	return cards
	// return nil
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	cards := d.cards[ len(d.cards) - n : ]
	return cards
	// return nil
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	// write code here
	// https://stackoverflow.com/questions/12264789/shuffle-array-in-go
	for i := len(d.cards) - 1; i >= 0; i-- {
        j := rand.IntN(i + 1)
        d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
    }
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here

	cutCards := d.cards[ : n ]
	restCards := d.cards[ n :]

	newCards := []Card{}
	newCards = append(newCards, restCards...)
	newCards = append(newCards, cutCards...)
	d.cards = newCards
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(5)
	fmt.Println("PeekTop 5")
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println("PeekCardAtIndex index array 10 - 15")
	fmt.Println(deck.PeekCardAtIndex(10).ToString()) // Jack Spade
	fmt.Println(deck.PeekCardAtIndex(11).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // 2 Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 3 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(5)
	fmt.Println("Deck Shuffle")
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	fmt.Println("Deck Cut")
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}