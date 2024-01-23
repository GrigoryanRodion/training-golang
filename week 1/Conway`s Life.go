// Game of life (Conway's life) in terminal

package main 

import (
	"fmt"
	"math/rand"
	"time"
)

var board [52][52]int		// Створення поля 
var rowArround = []int{-1,  0,  1, -1, 1, -1, 0, 1}		
var colArround = []int{-1, -1, -1,  0, 0,  1, 1, 1}		// Сусдні клітинки бомби по стовпцям (colArround) та рядкам (rowArround)

var numberNeighbors int		// Змінна, яка буде зберігати к-ть сусідів відносно клітинки

func fillingBoard(livingCellsAtStart int) {		// Заповнює поле на початку гри
	for i := 0; i < len(board); i++ {		// Створює межі поля (крайні індекси масиву)
		board[0][i] = -1
		board[i][0] = -1
		board[51][i] = -1
		board[i][51] = -1
	}

	for i, x, y := 0, 0, 0; i < livingCellsAtStart; {		// У випадковому порядку заповнюємо поле живими клітинками
		x, y = 1 + rand.Intn(50), 1 + rand.Intn(50)		// 1 + необхідно, щоб ми не виходили за межі поля (наприклад board[0][0]), 20 (а не 21) необхідна для тих же цілей
		
		if board[x][y] != 1 {
			board[x][y] = 1
			i++
		}
	}
}

func showBoard() {		// Виводить у терміналі поле board, який буде перевіряти - жива клітинка (1), або ні (0) і на основі цього виводити необхідний символ для позначення стану (без його справжніх меж "-1")
	fmt.Print("\033[H\033[2J")

	for i := 1; i < len(board) - 1; i++ {
		for j := 1; j < len(board) - 1; j++ {
			if board[i][j] == 1 {
				fmt.Print("@ ")
			} else {
				fmt.Print("- ")
			}
		}

		fmt.Println()
	}
}

func checkRule(i, j int) {		// Правила гри
	numberNeighbors = 0

	for k := 0; k < 8; k++ {		// Рахуємо к-ть живих клітинок, навколо комірки
		if board[i + rowArround[k]][j + colArround[k]] != -1 {
			numberNeighbors += board[i + rowArround[k]][j + colArround[k]]
		}
	}

	if board[i][j] == 0 && numberNeighbors == 3 {		// (Перевірка правил) На основі к-ть живих клітинок вирішуємо - виживе клітинка, або ні
		board[i][j] = 1
	} else if board[i][j] == 1 && (numberNeighbors <= 1 || 4 <= numberNeighbors) {
		board[i][j] = 0
	}
}

func game() {		// Запуск гри
	fillingBoard(200)

	for {
		showBoard()
		time.Sleep(400 * time.Millisecond)
		
		for i := 1; i < len(board) - 1; i++ {		// Перевіряємо кожену клітинку і вирішуємо (checkRule(i, j)) - виживе вона, чи ні
			for j := 1; j < len(board) - 1; j++ {
				checkRule(i, j)
			}
		}
	}
}

func main () {
	game()
}