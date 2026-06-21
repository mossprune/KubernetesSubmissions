package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New().String()

	for {
		fmt.Printf("%v: %s\n", time.Now().UTC(), id)
		time.Sleep(5 * time.Second)
	}
}
