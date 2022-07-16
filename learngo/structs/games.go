package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	id, price int
	name      string
}

type game struct {
	item
	genre string
}

// encodable and decodable game type
type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

func main() {

	decodeStruct()

	fmt.Println("********************************")

	games := make([]game, 0)

	games = append(games,
		game{
			item: item{
				id:    1,
				price: 50,
				name:  "god of war",
			},
			genre: "action adventure",
		},
		game{
			item: item{
				id:    2,
				price: 30,
				name:  "x-com 2",
			},
			genre: "strategy",
		},
		game{
			item: item{
				id:    3,
				price: 20,
				name:  "minecraft",
			},
			genre: "sandbox",
		},
	)

	gamesMapByID := make(map[int]game)
	for _, g := range games {
		gamesMapByID[g.id] = g
	}

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("-> list   : print all the games")
		fmt.Println("-> id N   : get game by id")
		fmt.Println("-> export   : export the data to json")
		fmt.Println("-> quit   : exit")

		if !input.Scan() {
			break
		}

		cmd := strings.Fields(input.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		case "id": // query by id exercise
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := gamesMapByID[id]
			if !ok {
				fmt.Println("Not found")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)

		case "save": // encoding exercise
			type jsonGame struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Genre string `json:"genre"`
				Price int    `json:"price"`
			}

			// load the data into the encodable game values
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable,
					jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}

			fmt.Println(string(out))
			return
		}

	}

}

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

func decodeStruct() {
	fmt.Println("********* Decoding **********")
	// load the initial data from json
	var decoded []jsonGame
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem:", err)
		return
	}

	// load the data into usual game values
	var games []game
	for _, dg := range decoded {
		games = append(games, game{item{dg.ID, dg.Price, dg.Name}, dg.Genre})
	}

	for _, g := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			g.id, g.name, "("+g.genre+")", g.price)
	}
}
