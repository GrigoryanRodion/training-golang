// walk-file-system-version-2.0
// NO CHEATING IN LIBRARY !!!

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func isValidDirectory(directory string) error {
	err := os.Chdir(directory)

	return err
}

func getNumberGoroutines() uint64 {
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

func walkDir(path string, ch chan string) {
	defer myHashTable.wg.Done()

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range entries {
		filePath := path + "/" + v.Name()
		ch <- filePath

		if v.IsDir() {
			myHashTable.wg.Add(1)
			walkDir(filePath, ch)
		}
	}
}

func walkFileSystem(path string, ch chan string) {
	myHashTable.wg.Add(1)
	walkDir(path, ch)
	myHashTable.wg.Wait()
	close(ch)
}

func showHashTable() {
	for i, v := range myHashTable.table {
		fmt.Printf("%s : %s\n", i, v)
	}
}

func main() {
	fmt.Println("Введіть шлях до директорії, на якій необхідно зробити обхід (/home/rodion): ")
	initialDirectory := getUserResponse()
	err := isValidDirectory(initialDirectory)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Введіть к-ть 'goroutines', які ви хочете використовувати: ")
	goroutines := getNumberGoroutines()

	start := time.Now()

	hash := getMD5Hash(initialDirectory)
	myHashTable.table = map[string]string{
		initialDirectory: hash,
	}

	ch := make(chan string)

	go walkFileSystem(initialDirectory, ch)

	for v := range ch {
		myHashTable.table[v] = getMD5Hash(v)
	}

	fmt.Println(goroutines)

	showHashTable()
	elapsed := time.Since(start)
	log.Printf("%s took", elapsed)

	// Save result in JSON format file

	fmt.Println("Введіть шлях до директорії, де вам необхідно зберегти сформовану хеш-таблицю у форматі JSON (/home/rodion): ")
	pathJSONfile := getUserResponse()

	err = isValidDirectory(pathJSONfile)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(myHashTable.table, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(pathJSONfile)
	os.WriteFile("Hash-table.json", data, 0644)
}
