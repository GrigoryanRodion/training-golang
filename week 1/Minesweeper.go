// Minesweeper game in terminal

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var board [11][11]int		// Створення поля 
var usedCellsBoard [11][11]bool		// Створення булевого поля, для 
var rowArround, colArround []int		// Масиви рядків і стовпців, за допомогою яких будемо рахувати кількість сусідів
var opennedCells int		// Загальна кількість відкритих комірок у полі. Якщо ігорьок відкриє всі клітини, крім бімб, то він виграє

func fillingBoard(numberMines int) {		// Заповнюємо рандомно поле	 
	rowArround = []int{-1,  0,  1, -1, 1, -1, 0, 1}		
	colArround = []int{-1, -1, -1,  0, 0,  1, 1, 1}		// Сусдні клітинки бомби по стовпцям (colArround) та рядкам (rowArround)

	for i := 0; i < len(board); i++ {		// Заповнюємо края поля [-1], щоб потім ідентифікувати їх
		board[i][0] = -1
		board[0][i] = -1
		board[i][10] = -1
		board[10][i] = -1
	}

	for i, x, y := 0, 0, 0; i < numberMines; {		// Створюємо у полі міни (9) у випадковому подрядку 
		x, y = 1 + rand.Intn(9), 1 + rand.Intn(9)

		if board[x][y] != 9 {		// Якщо поле в цій клітинці не має міни (0), то створюємо її; 
			board[x][y] = 9			// інакше, якщо міна вже є, цикл не виконається. Це необхідня для заповнення поле мінаму у РІЗНИХ клітинках
			i++

			for j := 0; j < 8; j++ {		
				if board[x + rowArround[j]][y + colArround[j]] != 9 && board[x + rowArround[j]][y + colArround[j]] != -1{		// Перебирає усіх сусів бомби, і збільшує їх значення на 1-цю 
					board[x + rowArround[j]][y + colArround[j]]++
				}
			}
		}
	}
}

func showBoard() {			// Створюємо поле для ігорька
	fmt.Print("\033[H\033[2J")
	fmt.Println("   1 2 3 4 5 6 7 8 9\n   _________________")		// Горизонтальне поле 1 - 9, для зручнішої орієнтації

	for i := 1; i < len(usedCellsBoard) - 1; i++ {
		fmt.Printf("%v |", i)		// Вертикальне поле 1 - 9, для зручнішої орієнтації

		for j := 1; j < len(usedCellsBoard) - 1; j++ {
			if usedCellsBoard[i][j] == false {
				fmt.Print("- ")
			} else if usedCellsBoard[i][j] == true {
				fmt.Printf("%v ", board[i][j])
			}
		}

		fmt.Println()
	}
}

func getColAndRow() (int, int) {		// Отримуємо значення ствопця і рядка від ігорька
	var rowStr, colStr string

	for {
		fmt.Print("\nВведіть номер стовпця: ")
		fmt.Scan(&colStr)
		fmt.Print("Введіть номер рядка: ")
		fmt.Scan(&rowStr)

		colInt, errCol := strconv.Atoi(colStr)
		rowInt, errRow := strconv.Atoi(rowStr)

		if errRow != nil || errCol != nil {
			fmt.Println("\nПомилка при конвертуванні стовпця, і/або рядка в числовий тип! Спробуйте ще раз.\n")
		} else if rowInt > 9 || rowInt < 1 || colInt > 9 || colInt < 1 {
			fmt.Println("\nПомилка! Значення стовпця, і/або рядка не має бути меншим за 1-цю, або більшим 9-ти. Спробуйте ще раз.\n")
		} else {	
			return rowInt, colInt;
		}
	}
}

func openNeighboringCells(row, col int) {		// Відкриття сусдніх клітинок, якщо поточна клітника 0
	for i := 0; i < 8; i++ {
		if board[row + rowArround[i]][col + colArround[i]] > 0 && 9 > board[row + rowArround[i]][col + colArround[i]] && !usedCellsBoard[row + rowArround[i]][col + colArround[i]] {		// Це жах x/
			usedCellsBoard[row + rowArround[i]][col + colArround[i]] = true
			opennedCells++
		} else if board[row + rowArround[i]][col + colArround[i]] == 0 && !usedCellsBoard[row + rowArround[i]][col + colArround[i]] {
			usedCellsBoard[row + rowArround[i]][col + colArround[i]] = true
			opennedCells++
			openNeighboringCells(row + rowArround[i], col + colArround[i])
		}
	}
}

func game() {		// Правила і логіка гри
	showBoard()
	row, col := getColAndRow()

	if board[row][col] == 9 {
		fmt.Println("\nУпс! Тут була бімба... Кінець гри.")
		return
	} else if board[row][col] == 0 && !usedCellsBoard[row][col] {
		usedCellsBoard[row][col] = true
		opennedCells++
		openNeighboringCells(row, col)
	} else if !usedCellsBoard[row][col] {
		usedCellsBoard[row][col] = true
		opennedCells++
	}

	if opennedCells == 9 * 9 - 10 {
		fmt.Println("\nВітаємо! Ви виграли!")
		return
	}
	
	game()
}



func restartGame() {		// Перезапуск гри за бажанням ігорька
	var playerAnswer string

	for isContinue := true; isContinue ; {
		fmt.Println("\nХочете спробувати ще раз? (Y/N)")
		fmt.Scan(&playerAnswer)

		switch playerAnswer {
		case "Y", "y":
			for i := 0; i < len(board); i++ {		// Заповняємо поле нулями і usedCellsBoard - false
				for j := 0; j < len(board); j++ {
					board[i][j] = 0
					usedCellsBoard[i][j] = false		
				}
			}

			opennedCells = 0
			fillingBoard(10)
			fmt.Print("\033[H\033[2J")
			fmt.Println("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nНова гра\n")

			game()
		case "N", "n":
			isContinue = false
		default:
			fmt.Println("\nПомилка! Відповідь надайте лише одним символом: Y / N")
			restartGame()
		}
	}
}

func main() {
	fmt.Println("\nПочаток гри")

	fillingBoard(10)
	game()
	restartGame()
}

