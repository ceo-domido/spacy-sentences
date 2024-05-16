package main

import (
	"fmt"
	"log"
	"spacy/spacy"
)

func main() {
	nlp, err := spacy.LoadModel("en_core_web_sm")
	if err != nil {
		log.Fatalf("Error loading model: %v", err)
	}
	defer nlp.Free()

	text := "This is a test. This is another sentence."
	doc, err := nlp.Parse(text)
	if err != nil {
		log.Fatalf("Error parsing text: %v", err)
	}
	defer doc.Free()

	sentences, err := doc.GetSentences()
	if err != nil {
		log.Fatalf("Error getting sentences: %v", err)
	}

	for i, sentence := range sentences {
		fmt.Printf("Sentence %d: %s\n", i+1, sentence)
	}
}
