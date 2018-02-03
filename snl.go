package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var snakes map[int]int
var ladders map[int]int
var players map[int]int
var turn int
var totalPlayers int
var finalPosition = 100

func init() {
	snakesInit()
	laddersInit()
}

func snakesInit() {
	snakes = make(map[int]int)
	snakes[51] = 11
	snakes[56] = 15
	snakes[62] = 57
	snakes[92] = 53
	snakes[98] = 8
}

func laddersInit() {
	ladders = make(map[int]int)
	ladders[2] = 38
	ladders[4] = 14
	ladders[9] = 31
	ladders[33] = 85
	ladders[52] = 88
	ladders[80] = 99
}

func playersInit(numOfPlayers int) {
	players = make(map[int]int)
	for i := 1; i <= numOfPlayers; i++ {
		players[i] = 0
	}
}

func start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of players: ")
	noOfPlayers, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Error reading the number of players")
	}
	numOfPlayers, err := strconv.Atoi(string(noOfPlayers))
	if err != nil {
		panic("Sorry! you need to enter a number and nothing else!!!,")
	}
	turn = 1
	totalPlayers = numOfPlayers
	playersInit(numOfPlayers)
	startGame()
}

func getInputFromUser() string {
	diceReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the dice reading for player %d: ", turn)
	diceReadingAsString, _, _ := diceReader.ReadLine()
	return string(diceReadingAsString)
}

func makeMove(diceReading int) {
	snakeBitMe := false
	ladderFound := false
	currentPosition := players[turn]
	nextPosition := currentPosition + diceReading
	var newPosition int
	if val, ok := snakes[nextPosition]; ok {
		newPosition = val
		snakeBitMe = true
	}
	if snakeBitMe {
		fmt.Printf("Snake bit player %d now he is in position %d\n", turn, newPosition)
	} else if val, ok := ladders[nextPosition]; ok {
		newPosition = val
		ladderFound = true
	} else if nextPosition > finalPosition {
		newPosition = nextPosition
	} else if nextPosition == finalPosition {
		fmt.Printf("Player %d won the Game, Everybody congratulate him\n", turn)
	} else {
		newPosition = nextPosition
	}
	if ladderFound {
		fmt.Printf("Player %d found a ladder now he is in position %d\n", turn, newPosition)
	}
	players[turn] = newPosition
	for player, position := range players {
		fmt.Printf("Player %d is at position: %d\n", player, position)
	}
}

func startGame() {
	for {
		diceReadingAsString := getInputFromUser()
		diceReading, err := strconv.Atoi(diceReadingAsString)
		if err != nil {
			fmt.Println("You should only enter a number not any string")
			continue
		}
		if diceReading < 1 && diceReading > 6 {
			fmt.Println("Stop Cheating! You can enter a number between 1 and 6 only")
		}
		makeMove(diceReading)
		if turn == totalPlayers {
			turn = 1
		} else {
			turn++
		}
	}
}
