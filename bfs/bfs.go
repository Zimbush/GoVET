package bfs

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Bfs() {
	bestand_alt := "/home/zimbu/Projekte/BFS/Bestand_alt"
	bestand_neu := "/home/zimbu/Projekte/BFS/Bestand_neu"

	input_bestand, err := os.Open(bestand_alt)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	output, err := os.Create(bestand_neu)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input_bestand)
	writer := bufio.NewWriter(output)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		writer.WriteString(text)
		writer.WriteString("\n")
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		os.Exit(1)
	}
	err = writer.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
