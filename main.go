package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage("images/cerpen.png")

	text, _ := client.Text()
	fmt.Println("text: ", text)
}
