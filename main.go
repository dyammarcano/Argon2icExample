package main

import (
	"github.com/dyammarcano/argon2/cryptography"
	"log"
	"strings"
)

func main() {
	password := "mySecretPassword"

	params := cryptography.NewArgon2Default()

	for i := 0; i < 5; i++ {
		hash := params.HashString(password)
		values := strings.Split(hash, "$")
		log.Printf("Argon2id hash: %s, salt: %s verify: %v\n", values[0], values[1], params.VerifyString(password, values[0], values[1]))
	}
}
