package main

import "fmt"

type Human struct {
	Name string
	Age  int
	// true - мужской, false - женский
	Gender bool
}

func (h Human) Speak() {
	male := "парень"
	female := "девушка"
	fmt.Printf("Я %s и мне %d лет. Я %s \n", h.Name, h.Age, map[bool]string{true: male, false: female}[h.Gender])
}

func (h Human) Walk() {
	fmt.Println("Ходить круто")
}

type Action struct {
	Human
	Skill string
}

func (a Action) Do() {
	fmt.Printf("%s имеет особый навык -  %s\n", a.Name, a.Skill)
}

func main() {
	a := Action{
		Human: Human{Name: "Ратмир", Age: 20, Gender: true},
		Skill: "рисовать",
	}

	a.Speak()
	a.Walk()

	a.Do()
}
