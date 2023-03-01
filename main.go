package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Person struct {
	Name  string
	Order int
}

func homePage(w http.ResponseWriter, r *http.Request) {
	names := getNames()

	for i, name := range names {
		fmt.Fprintln(w, i, name)
	}

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":7005", nil))
}

func main() {
	handleRequests()
}

func getNames() []string {
	names, err := readLines("names.txt")

	if err != nil {
		log.Fatal(err)
	}

	shuffleSlice(names)

	return names
}

func shuffleSlice(names []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
