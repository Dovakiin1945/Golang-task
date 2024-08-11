package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type Cocktail struct {
	ID              uuid.UUID
	Name            string
	FirstComponent  Component
	SecondComponent Component
	ThirdComponent  Component
	FourthComponent Component
	FifthComponent  Component
}

type Component struct {
	ID    uuid.UUID
	Name  string
	Cold  bool
	Warm  bool
	Hot   bool
	Slice bool
}

func NewCocktail(name string, first, second, third, fourth, fifth Component) Cocktail {
	return Cocktail{

		ID:              uuid.New(),
		Name:            name,
		FirstComponent:  first,
		SecondComponent: second,
		ThirdComponent:  third,
		FourthComponent: fourth,
		FifthComponent:  fifth,
	}
}

func NewComponent(name string, cold, warm, hot, slice bool) Component {
	return Component{
		ID:    uuid.New(),
		Name:  name,
		Cold:  cold,
		Warm:  warm,
		Hot:   hot,
		Slice: slice,
	}
}

func main() {

	text := "Hello Gold!"
	file, err := os.Create("hello.txt")

	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

	fmt.Println("Done.")



	fmt.Println("Выберите желаемое дейстиве")
	fmt.Println("1-Выбор коктейля по ингредиентам")
	fmt.Println("2-Выбор коктейля по названию")
	fmt.Println("3-Добавить свой коктейль")
	var i int
	fmt.Scan(i)
	switch i {
	case 1: fmt.Println("Zero")
	case 2: fmt.Println("One")
	case 3: fmt.Println("Two")
	case 4: fmt.Println("Three")







	// Создание компонентов для коктейлей
	tequila := NewComponent("Tequila", false, false, true, false)
	limeJuice := NewComponent("Lime Juice", true, false, false, false)
	tripleSec := NewComponent("Triple Sec", true, false, false, false)
	salt := NewComponent("Salt", false, false, false, false)
	zeroComponent := NewComponent("Капелька любви", false, false, false, false)

	// Создание коктейлей
	margarita := NewCocktail("Margarita", tequila, limeJuice, tripleSec, salt, zeroComponent)
	martini := NewCocktail("Martini", zeroComponent, zeroComponent, zeroComponent, zeroComponent, zeroComponent)
	mojito := NewCocktail("Mojito", zeroComponent, zeroComponent, zeroComponent, zeroComponent, zeroComponent)

	// Вывод компонента первого коктейля для проверки
	fmt.Println(margarita.FirstComponent)
	fmt.Println(martini.FirstComponent)
	fmt.Println(mojito.FirstComponent)

}

/*	NewCocktail("Margarita", )*/

// БД Коктейлей
/*Margarita := NewCocktail("Tequila", 123, "Triple Sec", "Salt", "0")
Martini := NewCocktail("Gin", "Vermouth", "Lemon", "0", "0")
Mojito := NewCocktail("White Rum", "Lime Juice", " Syrup", "Mint", "Soda")
// Что-бы не ругалась программа
fmt.Println(Margarita.FirstComponent)
fmt.Println(Martini.FirstComponent)
fmt.Println(Mojito.FirstComponent)*/
