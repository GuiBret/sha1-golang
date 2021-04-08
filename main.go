package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Entrez votre texte")
	str := ReadString(os.Stdin)
	fmt.Printf(str)
}

func ReadString(stdin io.Reader) string {

	scanner := bufio.NewReader(stdin)

	str, err := scanner.ReadString('\n')

	if err != nil && err != io.EOF {
		return ""
	}
	if strings.Contains(str, " ") {
		return ""
	}

	return str
}
