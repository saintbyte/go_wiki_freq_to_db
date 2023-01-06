package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Word struct {
	word      string
	frequency int
}

func printHelp() {
	fmt.Println("Usage: command file")
	os.Exit(0)
}

func parseLine(line string) Word {
	pair := strings.SplitN(line, " ", 2)
	word := strings.Trim(pair[0], " ")
	frequency, err := strconv.Atoi(strings.Trim(pair[1], " "))
	if err != nil {
		panic(err)
	}
	w := Word{word: word, frequency: frequency}
	return w
}
func insertToDB(tx pgx.Tx, word Word) error {
	_, err := tx.Exec(context.Background(), "insert into words(word, frequency) values ($1, $2)", word.word, word.frequency)
	if err != nil {
		return err
	}
	return nil
}
func connectToDB() *pgx.Conn {
	var db_url = "postgres://postgres:test1111@127.0.0.1:5432/default"
	if os.Getenv("DATABASE_URL") != "" {
		db_url = os.Getenv("DATABASE_URL")
	}
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
	}
	filePath := os.Args[1]
	fh, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	conn := connectToDB()
	defer conn.Close(context.Background())
	scanner := bufio.NewScanner(fh)
	tx, err := conn.Begin(context.Background())
	for scanner.Scan() {
		w := parseLine(scanner.Text())
		fmt.Println(w)
		insertToDB(tx, w)
	}
	tx.Commit(context.Background())
}
