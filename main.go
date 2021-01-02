package main

import (
	"fmt"
	io "io/ioutil"
	rand "math/rand"
	str "strings"
	"time"
)

var (
	datasetFormed [][]string
	index         int                 = 1
	chain         map[string][]string = make(map[string][]string)
)

func main() {
	rand.Seed(time.Now().Unix())
	dataRaw, _ := io.ReadFile("data.txt")
	dataset := string(dataRaw)
	var dataString []string
	var key string
	for _, oneString := range str.Split(dataset, "\n") {
		dataString = str.Split(oneString, " ")
		dataString = append([]string{"__BEGIN__"}, dataString...)
		dataString = append(dataString, "__END__")
		datasetFormed = append(datasetFormed, dataString)
	}
	for _, oneString := range datasetFormed {
		for _, word := range oneString[index:] {
			key = oneString[index-1]
			chain[key] = append(chain[key], word)
			index++
		}
		index = 1
	}

	word1 := "__BEGIN__"
	message := []string{word1}
	for {
		word2 := chain[word1][rand.Intn(len(chain[word1]))]
		word1 = word2
		if word2 == "__END__" {
			break
		} else {
			message = append(message, word2)
		}
	}
	fmt.Println(str.Join(message[1:], " "))
}
