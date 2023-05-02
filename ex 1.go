package main

import "fmt"

type Retangulo struct {
	Largura float64
	Altura  float64
}

func area(r Retangulo) {
	fmt.Println(r.Largura * r.Altura)
}

func main() {
	r := Retangulo{Largura: 5, Altura: 10}
	area(r)
}
