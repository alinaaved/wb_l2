package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func getLines(r io.Reader) []string {
	var lines []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// чтение из файла
func getLinesFile(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	return getLines(file)
}

// чтение из stdin
func getLinesStdin() []string {
	return getLines(os.Stdin)
}
