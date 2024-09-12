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
	setThenReadMode(client, "CW", 100)
	setThenReadFreq(client, 14050000)
	setThenReadMode(client, "USB", 3000)
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

func setThenReadMode(client rigctld.Client, mode string, bandpass rigctld.Frequency) {
	err := client.SetMode(mode, bandpass)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Mode and bandpass set to %s, %d\n", mode, bandpass)

	mode, bandpass, err = client.GetMode()
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Mode and bandpass is now %s, %d\n", mode, bandpass)
}
