package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	font := Getfont()
	input := os.Args[1]
	fmt.Print(AsciiArt(input, font))
}
func Getfont() []string {
	fontFile, _ := os.ReadFile(os.Args[2])
	// if len(os.Args) == 3 {
	// 	fontFile, err = os.ReadFile(os.Args[2] + ".txt")
	return strings.Split(string(fontFile), "\n")
}
func RowFinder(b byte) int {
	return (int(b)-32)*9 + 1
}
func AsciiArt(input string, font []string) string {
	var result string
	// split the input by new line whenever it finds \n
	lines := strings.Split(input, "\r")
	// looping through the slice that contains the arguments
	for i := 0; i < len(lines); i++ {
		switch lines[i] {
		case "":
			result += "\n"
		default:
			for j := 0; j < 8; j++ {
				for k := 0; k < len(lines[i]); k++ {
					//fmt.Println(j,k)
					result += (font[RowFinder(lines[i][k])+j])
				}
				result += "\n"
			}
		}
	}
	return result
}
