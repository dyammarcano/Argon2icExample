# Example Argon2id

This is an example of how to use Argon2id in Golang.

## Usage

```go 
package main

import (
	"github.com/dyammarcano/argon2"
	"log"
)

func main() {
	password := "mySecretPassword"

	params := argon2.NewArgon2Default()

	for i := 0; i < 5; i++ {
		result := params.HashString(password)
		log.Printf("Argon2id hash: %s, salt: %s verify: %v\n", result[0], result[1], params.Verify(password, []byte(result[0]), []byte(result[1])))
	}
}
```
