package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Hand struct {
	hand string
	bid  int
}

func handType(hand Hand, part2 bool) int {
	m := map[byte]int{}
	for i := range hand.hand {
		m[hand.hand[i]]++
	}

	// sort cards by count
	type card struct {
		c   byte
		cnt int
	}
	cards := []card{}
	for k, v := range m {
		cards = append(cards, card{k, v})
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].cnt > cards[j].cnt
	})

	if part2 {
		cnt := 0
		if len(cards) > 1 {
			for i := range cards {
				if cards[i].c == 'J' {
					cnt = cards[i].cnt
					cards = append(cards[:i], cards[i+1:]...) // remove 'J'
					break
				}
			}
			cards[0].cnt += cnt
		}
	}

	if cards[0].cnt == 5 {
		return 6
	} else if cards[0].cnt == 4 {
		return 5
	} else if cards[0].cnt == 3 && cards[1].cnt == 2 {
		return 4
	} else if cards[0].cnt == 3 && cards[1].cnt == 1 {
		return 3
	} else if cards[0].cnt == 2 && cards[1].cnt == 2 {
		return 2
	} else if cards[0].cnt == 2 && cards[1].cnt == 1 {
		return 1
	}
	return 0
}

func compareHands(h1, h2 Hand, part2 bool) bool {
	score := map[byte]int{
		'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8, '9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0,
	}
	if part2 {
		score = map[byte]int{
			'A': 12, 'K': 11, 'Q': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1, 'J': 0,
		}
	}

	if handType(h1, part2) > handType(h2, part2) {
		return true
	} else if handType(h1, part2) == handType(h2, part2) {
		for k := range h1.hand {
			if score[h1.hand[k]] == score[h2.hand[k]] {
				continue
			}
			return score[h1.hand[k]] > score[h2.hand[k]]
		}
	}
	return false
}

func part1(hands []Hand) {
	s := uint64(0)

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], false)
	})

	for i := range hands {
		s += uint64((len(hands) - i) * hands[i].bid)
	}

	fmt.Printf("%d\n", s)
}

func part2(hands []Hand) {
	s := uint64(0)

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], true)
	})

	for i := range hands {
		s += uint64((len(hands) - i) * hands[i].bid)
	}

	fmt.Printf("%d\n", s)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage: <filename>")
		os.Exit(-1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	defer file.Close()

	hands := []Hand{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var h string
		var b int
		fmt.Sscanf(scanner.Text(), "%s %d", &h, &b)
		hands = append(hands, Hand{h, b})
	}

	part1(hands)
	part2(hands)
}
