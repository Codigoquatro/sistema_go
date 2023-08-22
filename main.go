package main

import "fmt"

func main() {
	fmt.Println(Ola("Chris"))
}

func Ola(nome string) string {
	return "Olá," + nome
}
