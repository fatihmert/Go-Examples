package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var stringData string = "ABCDEFGHIJKLMNOPRSTUVYZQXW0123456789"

func generateRandomChar() string {
	return string(stringData[seededRand.Intn(len(stringData))])
}

func keygen(stringSize int, partSize int, seperator string) string {
	var result string

	for i := 0; i < partSize; i++ {
		for j := 0; j < stringSize; j++ {
			result += generateRandomChar()
		}
		result += seperator
	}

	return result[:len(result)-1]
}

func main() {
	fmt.Println(keygen(5, 4, "-"))
}
