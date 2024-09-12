package main

import (
	"github.com/k0swe/rigctld-go"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
)

func main() {

	client, err := rigctld.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	setThenReadFreq(client, 7200000)
	setThenReadFreq(client, 14050000)
}

func setThenReadFreq(client rigctld.Client, freq rigctld.Frequency) {
	p := message.NewPrinter(language.English)
	err := client.SetFreq(freq)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Frequency set to %s Hz\n", p.Sprintf("%d", freq))

	freq, err = client.GetFreq()
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Frequency is now %s Hz\n", p.Sprintf("%d", freq))
}
