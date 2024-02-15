// Обхід файлової системи

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	// "encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type HashTable struct {
	mtx   sync.Mutex
	wg    sync.WaitGroup
	table map[string]string
}

var myHashTable HashTable

func getUserResponse() string {
	scanner := bufio.NewScanner(os.Stdin)

	done := scanner.Scan()
	if done && scanner.Err() != nil {
		log.Fatal("error scanning text")
	}

	return scanner.Text()
}

func isValidDirectory(directory string) {
	err := os.Chdir(directory)
	if err != nil {
		log.Fatal("no such directory")
	}
}

func getNumberThreads() uint64 {
	response := getUserResponse()
	goroutines, err := strconv.ParseUint(response, 10, 64)

	if err != nil {
		log.Fatal("value is not a type uint8")
	}

	if goroutines == 0 {
		log.Fatal("error, minimum number of goroutines is 1")
	}

	return goroutines
}

func getMD5Hash(text string) string {
	data := []byte(text)
	hash := md5.Sum(data)

	return hex.EncodeToString(hash[:])
}

func walkingDir(directory string) {
	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		myHashTable.table[path] = getMD5Hash(path)

		return nil
	})

	if err != nil {
		log.Fatal(err) // ?error
	}
}

func showHashTable() {
	for i, v := range myHashTable.table {
		fmt.Printf("%s : %s\n", i, v)
	}
}

func main() {
	fmt.Println("Введіть шлях до директорії, на якій необхідно зробити обхід (/home/rodion): ")
	initialDirectory := getUserResponse()
	isValidDirectory(initialDirectory)

	fmt.Print("Введіть к-ть 'goroutines', які ви хочете використовувати: ")
	goroutines := getNumberThreads()

	start := time.Now()

	hash := getMD5Hash(initialDirectory)
	myHashTable.table = map[string]string{
		initialDirectory: hash,
	}

	walkingDir(initialDirectory)

	// showHashTable()
	fmt.Println(goroutines)

	elapsed := time.Since(start)
	log.Printf("%s took", elapsed)

	// Save result in JSON format file

	// fmt.Println("Введіть шлядо до директорії, де вам необхідно зберегти сформовану хеш-таблицю у форматі JSON (/home/rodion): ")
	// pathJSONfile := getUserResponse()
	// isValidDirectory(pathJSONfile)

	// data, err := json.MarshalIndent(myHashTable.table, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// os.Chdir(pathJSONfile)
	// os.WriteFile("Hash-table_JSON", data, 0644)
}
