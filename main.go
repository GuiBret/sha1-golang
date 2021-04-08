package main

import "fmt"

func main() {
	fmt.Println("Entrez votre texte")
	var baseStr string = ""

	fmt.Scan(&baseStr)

	fmt.Printf("%s\n", baseStr)
}
