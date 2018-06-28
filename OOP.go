package main

import "fmt"

type car struct {
	model string
	Type string
	price string

}

func oop()  {

	Car1 := car{
		model: "model:2020",
		Type:"type:kia",
		price:"price:1000000",
	}

	Car2 := car{
		model:"model:2014",
		Type:"type:ford",
		price:"price:1000000",
	}

	Car3 := car{
		model:"model:2017",
		Type:"type:toyota",
		price:"price:1000000",
	}

	fmt.Println(Car1,Car2,Car3)

}
func main() {

	oop()

}
