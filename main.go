package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Password struct om het gegenereerde wachtwoord en ID bij te houden
type Password struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func main() {
	// Command-line flags voor de lengte van het wachtwoord, het configuratiebestand en de gebruikersinvoer
	lengthPtr := flag.Int("length", 8, "Length of the generated password")
	flag.Parse()

	// Vraag de gebruiker om de gewenste lengte van het wachtwoord
	fmt.Print("Enter the length of the password: ")
	var userInput int
	fmt.Scan(&userInput)

	// Gebruikerinvoer voor wachtwoordlengte overschrijft -length vlag indien opgegeven
	if userInput > 0 {
		*lengthPtr = userInput
	}

	// Genereer wachtwoord
	password := generatePassword(*lengthPtr)

	// Schrijf wachtwoord naar JSON-bestand
	p := &Password{
		Value: password,
	}
	writePasswordToJSONFile(p)

	fmt.Println("Generated password:", password)
}

func generatePassword(length int) string {
	var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var password = make([]byte, length)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		password[i] = characters[rand.Intn(len(characters))]
	}

	return string(password)
}

func writePasswordToJSONFile(p *Password) {
	file, err := os.Create("password.json")
	if err != nil {
		log.Fatal("Error creating JSON file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(p)
	if err != nil {
		log.Fatal("Error writing password to JSON file:", err)
	}
}
