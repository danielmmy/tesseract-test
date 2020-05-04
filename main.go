package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("/home/daniel/Pictures/example_01.png")
	text, _ := client.Text()
	fmt.Println(text)
}
