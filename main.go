package main

import (
	"fmt"
	"math/rand"
	"strings"
	"bufio"
	"os"
	"strconv"
  "sort"
)

func generateDeck() []string {
	var deck = []string{}
	for i := 0; i < 4; i++ {
		for j := 2; j <= 10; j++ {
			var card string
			switch i {
			case 0:
				card = card + "♤"
			case 1:
				card = card + "♡"
			case 2:
				card = card + "♢"
			case 3:
				card = card + "♧"
			}
			card = card + strconv.Itoa(j)
			deck = append(deck, card)
		}

		for j := 1; j <= 4; j++ {
			var card string

			switch i {
			case 0:
				card = card + "♤"
			case 1:
				card = card + "♡"
			case 2:
				card = card + "♢"
			case 3:
				card = card + "♧"
			}

			switch j {
			case 1:
				card = card + "J"
			case 2:
				card = card + "Q"
			case 3:
				card = card + "K"
			case 4:
				card = card + "A"
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func draw(deck []string, inplay[]string) (d []string, i []string, c string) {
  var card string
  if len(deck) == 0 {
    seperate()
    fmt.Print("Shuffling.....")
    seperate()
    newDeck := generateDeck()
    if len(inplay) != 0 {
      newDeck = diff(newDeck, inplay)
      rand.Shuffle(len(newDeck), func(i, j int) {
        newDeck[i], newDeck[j] = newDeck[j], newDeck[i]
      })
    }
    card = newDeck[len(newDeck) - 1]
    newDeck = newDeck[:len(newDeck) - 1]
    inplay = append(inplay, card)
    return newDeck, inplay, card
  } else {
    card = deck[len(deck) - 1]
    deck = deck[:len(deck) - 1]
    inplay = append(inplay, card)
    return deck, inplay, card
  }
}
func updateTotal(a string, c int) int {
  i := strings.Trim(a, "♤♡♢♧")
  switch(i) {
  case "A":
    if (c + 11 > 21) {
      c = c + 1
    } else {
      c = c + 11
    }
  case "J","Q","K":
    c = c + 10
  default:
    j, err := strconv.Atoi(i)
    if err != nil {
      fmt.Println(err)
    }
    c = c + j
}
  return c
}
func deal(bet float64, d []string) (deck []string, owed float64) {  
  rand.Shuffle(len(d), func(i, j int) {
    d[i], d[j] = d[j], d[i]
  })
  var inplay []string
  playerCount, dealerCount := 0 , 0
  var playerHand []string
  var dealerHand []string

  deck, inplay, card := draw(d, inplay)
  playerHand = append(playerHand, card)
  fmt.Printf("Player is dealt: %v \n", card)
  playerCount = updateTotal(card, playerCount)
   
  seperate()
  
  deck, inplay, card = draw(deck, inplay)
  dealerHand = append(dealerHand, card)
  fmt.Printf("Dealer is dealt: %v \n", card)
  dealerCount = updateTotal(card, dealerCount)
  
  seperate()

  deck, inplay, card = draw(deck, inplay)
  playerHand = append(playerHand, card)
  fmt.Printf("Player is dealt: %v \n", card)
  playerCount = updateTotal(card, playerCount)
  fmt.Printf("Player has a total of %v \n", playerCount)

  seperate()

  deck, inplay, card = draw(deck, inplay)
  dealerHand = append(dealerHand, card)
  fmt.Println("Dealer draws ???")
  dealerCount = updateTotal(card, dealerCount)

    if playerCount == 21{
      fmt.Printf("%v\n", playerHand)
      fmt.Printf("You lucky bastard...\n")
      bet = bet * 3 / 2
      return deck, bet 
    }
  loop := true
  for playerCount <= 21 {
    if loop == false {
      break
    } else {
      seperate()
      fmt.Printf("Dealer has %v #\n", dealerHand[0])
      fmt.Println()
      fmt.Printf("Player has %v || Total: %v\n", playerHand, playerCount)

      seperate()
      
      fmt.Print("Would you like to Hit or Stay(h/s): ")
      reader := bufio.NewReader(os.Stdin)
      char, _, err := reader.ReadRune()
      switch char {
        case 'Q', 'q':
          return
        case 'h', 'H':
          deck, inplay, card = draw(deck, inplay)
          playerHand = append(playerHand, card)
          fmt.Printf("Player is dealt: %v \n", card)
          playerCount = updateTotal(card, playerCount)
          if playerCount > 21 {
          fmt.Printf("Player busts with %v %v!, Only $%v...\n", playerHand, playerCount, bet)
          bet = -bet
          return deck, bet
        }
        case 's','S':
        loop = false
          break
      }
		  if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		  }
    }
  }
  seperate()
  fmt.Printf("Dealer shows %v || Total: %v\n", dealerHand, dealerCount)
  for dealerCount <= 21 {
    if dealerCount >= 17 || dealerCount == 21 {
      fmt.Printf("Dealer stays with %v || Total: %v\n", dealerHand, dealerCount)
      break
    } else if dealerCount < 17{
    deck, inplay, card = draw(deck, inplay)
    dealerHand = append(dealerHand, card)
    fmt.Printf("Dealer is dealt: %v \n", card)
    dealerCount = updateTotal(card, dealerCount)
      if dealerCount > 21 {
        fmt.Printf("Dealer Busts with %v %v! You win $%v!!\n", dealerHand, dealerCount, bet)
        seperate()
        return deck, bet
      }
    }
  }
  if playerCount < dealerCount {
    seperate()
    fmt.Printf("Dealer: %v || Total: %v\n", dealerHand, dealerCount)
    fmt.Printf("Player: %v || Total: %v\n", playerHand, playerCount)
    fmt.Println("Better Luck Next Time!")
    seperate()
    bet = -bet
    return deck, bet
  } else if playerCount == dealerCount {
    seperate()
    fmt.Printf("Dealer: %v || Total: %v\n", dealerHand, dealerCount)
    fmt.Printf("Player: %v || Total: %v\n", playerHand, playerCount)
    fmt.Println("Its a push!")
    return deck, 0 
  } else {
    seperate()
    fmt.Printf("Dealer: %v || Total: %v\n", dealerHand, dealerCount)
    fmt.Printf("Player: %v || Total: %v\n", playerHand, playerCount)
    fmt.Printf("You Win $%v!\n", bet)
    seperate()
    return deck, bet
  }
}

func startGame() {
  balance := 100.00
  var owed float64
  loop := true
  deck := generateDeck()
  for balance > 0 && loop {
    fmt.Printf("Your balance is $%.2f, how much would you like to bet: ", balance)
    reader := bufio.NewReader(os.Stdin)
    string, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("L:69 -> Error reading input, try again.")
    }
    string = strings.TrimSpace(string)
    betAmount, err := strconv.ParseFloat(string, 64)
    if err != nil {
      fmt.Println("U WoT M8!?")
    }
    seperate()
    if betAmount > balance {
      fmt.Println("You don't have that much money...")
    } else if (betAmount == balance) {
      fmt.Println("Scared money don't make money, Good Luck")
      seperate()
      deck, owed = deal(betAmount, deck)
      balance = balance + owed
    } else if (betAmount != 0){
      deck, owed = deal(betAmount, deck)
      balance = balance + owed
    }
	}
	return
}

func seperate() {
  fmt.Println("")
  for i := 0; i < 40; i++ {
    fmt.Print("-")
  }
  fmt.Println("")
  fmt.Println("")

  return
}

func diff(a []string, b []string) []string {
	var m map[string]bool
	m = make(map[string]bool, len(b))
	for _, s := range b {
		m[s] = false
	}
	var diff []string
	for _, s := range a {
		if _, ok := m[s]; !ok {
			diff = append(diff, s)
			continue
		}
		m[s] = true
	}
	sort.Strings(diff)
	return diff
}


func main() {
	var loop int
	for loop != 1 {
    seperate()
		fmt.Print("Would you like to play blackjack?(y/n): ")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		switch char {
		case 'y', 'Y':
			fmt.Println("So you want to gamble huh?")
      seperate()
			startGame()
      fmt.Println("Thanks for playing!")
			return
		case 'n', 'N', 'q', 'Q':
			fmt.Println("So you want to be rich huh?")
			return
		}
	}
}

