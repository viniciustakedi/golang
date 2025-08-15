package main

import "fmt"

// Demonstre no seu código:
// Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
// Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
type Humanos interface {
	Speak()
}

type Pessoa struct {
	Name string
	Age  int
}

func (p *Pessoa) Speak() {
	fmt.Println("To falanuuu meuuu.")
}

func SaySomething(h Humanos) {
	h.Speak()
}

func main() {
	pessoa := &Pessoa{
		Name: "Vinius Takedi",
		Age:  22,
	}

	pessoa.Speak()
	SaySomething(pessoa)
}
