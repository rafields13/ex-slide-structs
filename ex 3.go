package main

import "fmt"

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func media(a Aluno) {
	sum := 0.0
	for _, nota := range a.Notas {
		sum += nota
	}
	fmt.Println(sum / float64(len(a.Notas)))
}

func main() {
	a := Aluno{Nome: "Rafael", Idade: 18, Notas: []float64{10, 8, 9}}
	media(a)
}
