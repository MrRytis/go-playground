package main

import "fmt"

type Animal struct {
	Name  string
	Age   int
	voice string
}

func (a *Animal) Voice() string {
	return a.voice
}

type Dog struct {
	Animal
	favoriteToy string
}

func (d *Dog) setFavoriteToy(toy string) {
	d.favoriteToy = toy
}

type Cat struct {
	Animal
	Breed string
}

type AnimalAction interface {
	Voice() string
}

func main() {
	golden := Dog{
		Animal: Animal{
			Name:  "golden",
			Age:   2,
			voice: "woof",
		},
		favoriteToy: "ball",
	}

	golden.setFavoriteToy("stick")

	british := Cat{
		Animal: Animal{
			Name:  "british",
			Age:   3,
			voice: "meow",
		},
		Breed: "british",
	}

	fmt.Printf("golden voice: %s\n", golden.Voice())
	fmt.Printf("british voice: %s\n", british.Voice())

	fmt.Println("----")

	animals := []AnimalAction{&golden, &british}
	for _, animal := range animals {
		fmt.Printf("animal voice: %s\n", animal.Voice())
	}

	fmt.Println("----")
	fmt.Printf("British: %+v\n", british)

	british.Name = "british2"
	british.Age = 4

	fmt.Printf("British: %+v\n", british)
}
