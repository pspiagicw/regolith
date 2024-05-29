package main

import (
	"log"

	"github.com/pspiagicw/regolith"
)

func main() {
	rg, err := regolith.New(&regolith.Config{
		StartWords: []string{"if", "fn", "while"},
		EndWords:   []string{"end"},
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	for true {
		content, err := rg.Input()

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("Content: %v", content)
	}
}
