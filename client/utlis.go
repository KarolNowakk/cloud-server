package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return strings.Trim(text, "\n")
}

func trimBytes(bytes []byte) []byte {
	var firstZeroByteIndex int
	firstZeroWasHitted := false

	for i, singleByte := range bytes {
		if singleByte == 0 && !firstZeroWasHitted {
			firstZeroByteIndex = i
			firstZeroWasHitted = true
		}
		if singleByte != 0 {
			firstZeroByteIndex = -1
			firstZeroWasHitted = false
		}
	}

	if firstZeroByteIndex == -1 {
		firstZeroByteIndex = len(bytes)
	}

	bytes = bytes[:firstZeroByteIndex]

	return bytes
}
