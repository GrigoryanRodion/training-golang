// Minesweeper game in terminal version-2.0
/*
*	Version 2.0
*	- перенесли ('захардкодені') розміри поля до окремої структури даних 'Game'
*	- додана можливість задавати розмір поля і кількість мін у ньому
* 	- позбавилися від меж '-1' у масиві board[][]int
*	- позбавилися від масиву з типом boolean, використовуючи лише масив типу integer
*	- markBoardHorizontally() i markBoardVertically() - недороблені
* 	- неправильна робота openNeighboringCells() при ширині, або довжині - 2 і при великій довжині, або ширині
 */

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var row []int = []int{-1, 0, 1, -1, 1, -1, 0, 1}
var col []int = []int{-1, -1, -1, 0, 0, 1, 1, 1} // Масиви рядків і стовпців, за допомогою яких будемо рахувати кількість сусідів
var opennedCells int                             // Загальна кількість відкритих комірок у полі. Якщо ігорьок відкриє всі клітини, крім бімб, то він виграє

type Game struct {
	boardHeight int
	boardWidth  int
	mines       int
	board       [][]int
}

func getUserAnswer(limit int) int { // Отримуємо необхідне число (залежно від контексту) від користувача
	var userAnswer string
	fmt.Scan(&userAnswer)

	value, errValue := strconv.Atoi(userAnswer)

	if errValue != nil {
		fmt.Print("\nПомилка! Невірно вказане число. Необхідно вказувати числовий тип даних. Повторіть спробу ще раз: ")
		return getUserAnswer(limit)
	} else if value < limit {
		fmt.Printf("\nПомилка! Число не може бути менше за %v. Повторіть спробу ще раз: ", limit)
		return getUserAnswer(limit)
	} else {
		return value
	}
}

func getNumberMines(g *Game) int { // Отримуємо к-ть мін, яку хоче стоврити користувач
	userAnswer := getUserAnswer(1)

	if userAnswer >= (g.boardWidth * g.boardHeight) {
		fmt.Print("\nПомилка! Кількість бімб не може перевищувати, або відповідати розміру поля. Повторіть спробу ще раз: ")
		return getNumberMines(g)
	} else {
		return userAnswer
	}
}

func (g *Game) createBoard() { // Створюємо поле гри
	fmt.Println("---------- New Game ----------")
	fmt.Print("\nВведіть необхідно висоту поля: ")
	g.boardHeight = getUserAnswer(2)
	fmt.Print("\nВведіть необхідно ширину поля: ")
	g.boardWidth = getUserAnswer(2)
	fmt.Print("\nВведіть к-ть бімб на початку гри: ")
	g.mines = getNumberMines(g)
	fmt.Print(g.mines)

	g.board = make([][]int, g.boardHeight) // Створюємо ширину поля (поле board)
	for i := 0; i < g.boardHeight; i++ {   // Створюємо висоту поля (поле board)
		g.board[i] = make([]int, g.boardWidth)
	}
}

func getColAndRow(g *Game) (int, int) { // Отримуємо значення сповпця і рядка від користувача
	fmt.Print("\nВведіть номер стовпця: ")
	col := getUserAnswer(1) - 1
	fmt.Print("\nВведіть номер рядка: ")
	row := getUserAnswer(1) - 1

	if col >= g.boardWidth || row >= g.boardHeight {
		g.showBoard()
		fmt.Println("\nПомилка! Колона, або рядок виходять за межі поля. Спробуйте ввести їх знову.")
		return getColAndRow(g)
	} else {
		return col, row
	}
}

func checkIndexesBoard(x, y, i int, g *Game) bool { // Перевіряємо індекс масиву на можливий вихід за його межі
	if x+row[i] > -1 && y+col[i] > -1 && x+row[i] < g.boardHeight && y+col[i] < g.boardWidth {
		return true
	} else {
		return false
	}
}

func (g *Game) fillingBoard() { // Заповнюємо рандомно поле
	for i, x, y := 0, 0, 0; i < g.mines; { // Створюємо у полі міни (9) у випадковому подрядку
		x, y = rand.Intn(g.boardHeight), rand.Intn(g.boardWidth)

		if g.board[x][y] != 9 { // Якщо поле в цій клітинці не має міни (0), то створюємо її;
			g.board[x][y] = 9 // інакше, якщо міна вже є, цикл не виконається. Це необхідня для заповнення поле мінаму у РІЗНИХ клітинках
			i++

			for j := 0; j < 8; j++ {
				if x+row[j] > -1 && y+col[j] > -1 && x+row[j] < g.boardHeight && y+col[j] < g.boardWidth && g.board[x+row[j]][y+col[j]] != 9 {
					g.board[x+row[j]][y+col[j]]++
				}
			}
		}
	}
}

func (g *Game) markBoardHorizontally() { // !недороблено // Розмітка поля зверхну, для більш зручної орієнтації по горизонталі
	fmt.Print("      ")
	// for i := 1; i <= len(g.board); i++ {	// Розмітка не парних чисел

	// }

	for i := 1; i <= g.boardWidth; i++ { // Розмітка парних чисел
		if i <= 9 {
			fmt.Printf("%v ", i)
		} else if i > 9 && (i%2) == 0 {
			fmt.Printf("%v  ", i)
		}
	}
	fmt.Println()

	fmt.Print("      ")
	for i := 1; i <= g.boardWidth; i++ {
		if i <= 9 {
			fmt.Print("| ")
		} else if i > 9 && (i%2) == 0 {
			fmt.Print("|   ")
		}
	}
	fmt.Println()
}

func (g *Game) markBoardVertically(i int) { // !недороблено // Розмітка поля збоку, для більш зручної орієнтації по вертикалі
	if i <= 9 {
		fmt.Printf("%v  => ", i)
	} else {
		fmt.Printf("%v => ", i)
	}
}

func (g *Game) showBoard() { // Виводимо у термінал поле
	fmt.Print("\033[H\033[2J")

	g.markBoardHorizontally()

	for i := 0; i < len(g.board); i++ {
		g.markBoardVertically(i + 1)

		for j := 0; j < len(g.board[i]); j++ {
			if g.board[i][j] >= 0 {
				fmt.Print("- ")
			} else if g.board[i][j] == -10 {
				fmt.Print("0 ")
			} else {
				fmt.Printf("%v ", g.board[i][j]*(-1))
			}
		}

		fmt.Println()
	}
}

func (g *Game) openNeighboringCells(x, y int) { // Якщо кормірка, на яку натиснув користувач = 0, то відкриваємо усі сусідні комірки рекурсивно
	for i := 0; i < 8; i++ {
		if checkIndexesBoard(x, y, i, g) && g.board[x+row[i]][y+col[i]] > 0 && g.board[x+row[i]][y+col[i]] != 9 {
			g.board[x+row[i]][y+col[i]] *= -1
			opennedCells++
		} else if checkIndexesBoard(x, y, i, g) && g.board[x+row[i]][y+col[i]] == 0 {
			g.board[x+row[i]][y+col[i]] = -10
			opennedCells++
			g.openNeighboringCells(x+row[i], y+col[i])
		}
	}
}

func (g *Game) runGame() { // Старт гри
	g.showBoard()
	col, row := getColAndRow(g)

	if g.board[row][col] == 9 {
		fmt.Println("\nУпс! Тут була бімба... Кінець гри.")
		return
	} else if g.board[row][col] == 0 {
		g.board[row][col] = -10
		opennedCells++
		g.openNeighboringCells(col, row)
	} else if g.board[row][col] > 0 {
		g.board[row][col] *= -1
		opennedCells++
	}

	if opennedCells == g.boardWidth*g.boardHeight-g.mines { // Умова перемоги користувача
		g.showBoard()
		fmt.Println("\nВітаємо! Ви виграли!")
		return
	}

	g.runGame()
}

func (g *Game) restartGame() {
	var userAnswer string
	fmt.Print("\nХочете спробувати ще раз? (Y/N): ")
	fmt.Scan(&userAnswer)

	switch userAnswer {
	case "Y", "y":
		fmt.Print("\033[H\033[2J")
		g.boardHeight = 0
		g.boardWidth = 0
		g.mines = 0
		g.board = nil
		opennedCells = 0

		g.createBoard()
		g.fillingBoard()
		g.runGame()
		g.restartGame()
	case "N", "n":
		return
	default:
		// fmt.Print("\033[H\033[2J")
		fmt.Print("\nПомилка! Відповідь надайте лише одним символом: Y / N")
		g.restartGame()
	}
}

func main() {
	fmt.Println("\033[H\033[2J========== Minesweeper ==========\n")
	var g Game
	g.createBoard()

	g.fillingBoard()

	g.runGame()

	g.restartGame()
}
