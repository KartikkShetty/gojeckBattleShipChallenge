package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	flReader := &fileReaderWriter{}
	flReader.Init("input.txt")
	for {
		var order int

		value, err := flReader.ReadLine()
		if err != nil {
			return
		}
		order, _ = strconv.Atoi(value)

		if order == 0 {
			fmt.Println("Invalid matrix order")
			return
		}

		battlefield1 := &BattleField{}
		battlefield2 := &BattleField{}

		battlefield1.Initialize(order)
		battlefield2.Initialize(order)

		flReader.ReadLine()

		value, _ = flReader.ReadLine()
		battleShip1Positions := strings.Split(value, ",")

		value, _ = flReader.ReadLine()
		battleShip2Positions := strings.Split(value, ",")

		battlefield1.ArrangeShips(battleShip1Positions)
		battlefield2.ArrangeShips(battleShip2Positions)

		flReader.ReadLine()

		value, _ = flReader.ReadLine()
		battleship1Moves := strings.Split(value, ",")

		value, _ = flReader.ReadLine()
		battleship2Moves := strings.Split(value, ",")

		battlefield1.Attack(battleship2Moves)
		battlefield2.Attack(battleship1Moves)
		printOutput(flReader, battlefield1, battlefield2)
	}
}
func printOutput(flReader *fileReaderWriter, battlefield1 *BattleField, battlefield2 *BattleField) {
	output := ""
	output = "Player1 \n"
	output = output + battlefield1.Print()
	output = output + "\n"
	output = output + "Player2 \n"
	output = output + battlefield2.Print()
	output = output + "\n"
	output = output + fmt.Sprintf("P1:%d \n", battlefield2.hits)
	output = output + fmt.Sprintf("P2:%d \n", battlefield1.hits)

	if battlefield2.hits == battlefield1.hits {
		output = output + "It is a draw"
	} else if battlefield2.hits > battlefield1.hits {
		output = output + "P1 wins"
	} else {
		output = output + "P2 wins"
	}
	output = output + "\n"
	flReader.WriteToFile("output.txt", output)
	flReader.ReadLine()

}

type fileReaderWriter struct {
	scanner *bufio.Scanner
}

//Init initializes the fileReader
func (f *fileReaderWriter) Init(fileName string) {
	data, _ := os.Open(fileName)
	f.scanner = bufio.NewScanner(data)

}

//ReadLine reads the next line of input
func (f *fileReaderWriter) ReadLine() (string, error) {

	if !f.scanner.Scan() {
		return "", errors.New("EOF")
	}
	return f.scanner.Text(), nil
}

//WriteToFile writes to file the data string
func (f *fileReaderWriter) WriteToFile(fileName string, data string) {
	fileHandle := f.getOutputWriter(fileName)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, data)
	writer.Flush()
}

func (f *fileReaderWriter) getOutputWriter(fileName string) *os.File {
	var fileHandle *os.File
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fileHandle, _ = os.Create(fileName)
	} else {
		fileHandle, _ = os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0666)
	}
	return fileHandle
}

//BattleField represents a BattleShip
type BattleField struct {
	battleField [][]string
	hits        int
}

//Initialize initializes the BattleField
func (bs *BattleField) Initialize(order int) {
	matrix := make([][]string, order)
	for i := 0; i < order; i++ {
		matrix[i] = make([]string, 0, order)
		arr := make([]string, order)
		for j := 0; j < order; j++ {
			arr[j] = "_"
			matrix[i] = append(matrix[i], arr[j])
		}
	}
	bs.battleField = matrix
}

//ArrangeShips ,arranges ships on the BattleField at the provided coordinates
func (bs *BattleField) ArrangeShips(battleShipPositions []string) {
	for i := 0; i < len(battleShipPositions); i++ {
		r, _ := strconv.Atoi(strings.Split(battleShipPositions[i], ":")[0])
		c, _ := strconv.Atoi(strings.Split(battleShipPositions[i], ":")[1])
		bs.battleField[r][c] = "B"
	}
}

//Attack ,attack the ships on the battle BattleField with the provided coordinates
func (bs *BattleField) Attack(attackPositions []string) {

	hits := 0
	for i := 0; i < len(attackPositions); i++ {
		r, _ := strconv.Atoi(strings.Split(attackPositions[i], ":")[0])
		c, _ := strconv.Atoi(strings.Split(attackPositions[i], ":")[1])
		if bs.battleField[r][c] == "B" {
			bs.battleField[r][c] = "X"
			hits++
		} else {
			bs.battleField[r][c] = "O"
		}
	}
	bs.hits = hits
}

//Print ,prints the current state of the battle field
func (bs *BattleField) Print() string {
	value := ""
	for i := 0; i < len(bs.battleField); i++ {
		for j := 0; j < len(bs.battleField[i]); j++ {
			value = value + fmt.Sprintf(bs.battleField[i][j]+" ")
		}
		value = value + fmt.Sprintf("\n")
	}
	return value
}
