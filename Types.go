package main
import "fmt"

func add(x, y int) int {
	return  x + y
}

func multiple(a,b string) (string,string){
	return a,b
}

func main() {
	num1,num2  := 20,20
	var resultInt = add(num1,num2)
	fmt.Println(resultInt)
	fmt.Println(add(num1,num2))
	w1,w2 := "Hey,","There"
	fmt.Println(multiple(w1,w2))

}
