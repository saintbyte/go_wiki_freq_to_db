package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Word struct {
	word  string
	count int
}

func printHelp() {
	fmt.Println("Usage: command file")
}
func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	filePath := os.Args[1]
	fh, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.SplitN(line, " ", 2)
		fmt.Println(pair[0], "|", pair[1])
		word := strings.Trim(pair[0], " ")
		count, err := strconv.Atoi(strings.Trim(pair[1], " "))
		if err != nil {
			panic(err)
		}
		w := Word{word: word, count: count}
	}
}
