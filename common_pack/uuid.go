package main

import (
	"fmt"

	"github.com/google/uuid"
)

func GetUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

func main() {
	fmt.Println(GetUUID())
}
