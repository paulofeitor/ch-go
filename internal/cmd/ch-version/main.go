package main

import (
	"fmt"

	"github.com/paulofeitor/ch-go/internal/version"
)

func main() {
	fmt.Println("version", version.Get())
}
