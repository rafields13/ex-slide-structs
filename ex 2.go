package main

import "fmt"

type Livro struct {
	Titulo        string
	Autor         string
	AnoPublicacao int
}

func info(l Livro) {
	fmt.Println(l.Titulo)
	fmt.Println(l.Autor)
	fmt.Println(l.AnoPublicacao)
}

func main() {
	l := Livro{Titulo: "O homem mais rico da Babilônia", Autor: "George S. Clason", AnoPublicacao: 1926}
	info(l)
}
