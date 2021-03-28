package main

import (
	"image/png"
	"bufio"
	"os"
	"fmt"
	"strings"

	"github.com/jakobvarmose/go-qidenticon"
)

var input string
var size int

func main() {
	
	fmt.Print("Enter your string:")
	inputReader := bufio.NewReader(os.Stdin)
        input, _ := inputReader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")  // "\r\n" for Windows, to get the same results as with Linux/macOS.
	fmt.Print("Image size, x = y:")
        fmt.Scanf("%d",&size)

	code := qidenticon.Code(input)
	settings := qidenticon.DefaultSettings()
	img := qidenticon.Render(code, size, settings)
	w, err := os.Create("identicon.png")
	if err != nil {
		panic(err)
	}
	defer w.Close()
	err = png.Encode(w, img)
	if err != nil {
		panic(err)
	}
}
