// Conway's life in terminal version-2.0
/*	
*	Version 2.0
*	- позбавилися від меж '-1' у масиві board[][]int
*	- перенесли ('захардкодені') розміри поля до окремої структури даних 'Game' (boardWidth, boardHeight)
*	- додана можливість задавати розмір поля і кількість клітинок першого покоління
*	- був створений другий двовимірний масив для більш точної перевірки правил
*/		

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var row []int = []int{-1, 0, 1, -1, 1, -1, 0, 1}	// Рядки навколо клітинки
var col []int = []int{-1, -1, -1, 0, 0, 1, 1, 1}	// Колонки навколо клітинки

type Game struct {
	boardHeight int
	boardWidth 	int
	board 		[][]int
	board2		[][]int		// Додати друге поле!
}

func getUserAnswer() int {	// Отримуємо необхідне число (залежно від контексту) від користувача
	var userAnswer string
	fmt.Scan(&userAnswer)

	value, errValue := strconv.Atoi(userAnswer)

	if errValue != nil {
		fmt.Print("\nПомилка! Невірно вказане число. Необхідно вказувати числовий тип даних. Повторіть спробу ще раз: ")
		return getUserAnswer()
	} else if value < 1 {
		fmt.Print("\nПомилка! Число не може бути менше за 1. Повторіть спробу ще раз: ")
		return getUserAnswer()
	} else {
		return value
	}
}

func (g *Game) getFirstGenetarion() int {
	userAnswer := getUserAnswer()

	if userAnswer > (g.boardWidth * g.boardHeight) {
		fmt.Print("\nПомилка! Кількість живих клітинок не може перевищувати розмір поля. Повторіть спробу ще раз: ")
		return g.getFirstGenetarion()
	} else {
		return userAnswer
	}
}

func (g *Game) createBoard()  {		// Створюємо перше і друге поле гри (друге для більш точної перевірки правил)
	fmt.Print("\nВведіть необхідно висоту поля: ")
	g.boardHeight = getUserAnswer()
	fmt.Print("\nВведіть необхідно ширину поля: ")
	g.boardWidth = getUserAnswer()

	g.board = make([][]int, g.boardHeight)	// Створюємо ширину поля (поле board)
	g.board2 = make([][]int, g.boardHeight)
	for i := 0; i < g.boardHeight; i++ {		// Створюємо висоту поля (поле board)
		g.board[i] = make([]int, g.boardWidth)
		g.board2[i] = make([]int, g.boardWidth)
	}
}

func (g *Game) fillingBoard(livingCellsAtStart int) {		// У випадковому порядку заповнюємо поле живими клітинками
	for i, x, y := 0, 0, 0; i < livingCellsAtStart; {
		x, y = rand.Intn(g.boardHeight), rand.Intn(g.boardWidth)

		if g.board[x][y] != 1 {
			g.board[x][y] = 1
			i++
		}
	}

	g.board2 = g.board
}

func (g *Game) showBoard() {		// Виводить у терміналі поле board, який буде перевіряти - жива клітинка (1), або ні (0) і на основі цього виводити необхідний символ для позначення стану
	fmt.Print("\033[H\033[2J")

	for i := 0; i < len(g.board); i++ {
		for j := 0; j < len(g.board[i]); j++ {
			if g.board[i][j] == 1 {
				fmt.Print("@ ")
			} else {
				fmt.Print("- ")
			}
		}

		fmt.Println()
	}
}

func (g *Game) checkRule(i, j int) {	// Правила гри
	numberNeighbors := 0

	for k := 0; k < 8; k++ {	// Рахуємо к-ть живих клітинок, навколо комірки
		if i + row[k] > -1 && j + col[k] > -1 && i + row[k] < len(g.board) && j + col[k] < len(g.board[i]) {	// Перевіряємо, щоб ми не вийшли за межі масиву
			numberNeighbors += g.board[i + row[k]][j + col[k]]
		}
	}

	if g.board[i][j] == 0 && numberNeighbors == 3 {	// Перевірка правил 
		g.board2[i][j] = 1
	} else if g.board[i][j] == 1 && (numberNeighbors <= 1 || 4 <= numberNeighbors) {
		g.board2[i][j] = 0
	}
}

func game(g *Game) {	// Запуск гри
	for {
		g.showBoard()
		time.Sleep(300 * time.Millisecond)

		for i := 0; i < len(g.board); i++ {	// Перевіряємо кожену клітинку і вирішуємо (checkRule(i, j)) - виживе вона, чи ні
			for j := 0; j < len(g.board[i]); j++ {
				g.checkRule(i, j)
			}
		}

		g.board = g.board2
	}
}

func main() {
	fmt.Print("\033[H\033[2J")

	var g Game
	g.createBoard()

	fmt.Print("\nВведіть к-ть живих клітинок на початку гри: ")
	g.fillingBoard(g.getFirstGenetarion())

	game(&g)
}