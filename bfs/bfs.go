package bfs

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Bfs() {
	filename := "/home/zimbu/Projekte/BFS/Bestand"

	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		os.Exit(1)
	}
}
