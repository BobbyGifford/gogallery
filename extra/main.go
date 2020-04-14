package main

import "fmt"

type User struct {
	Name string
}

func main() {
	//t, err := template.ParseFiles("hello.gohtml")
	//if err != nil {
	//	panic(err)
	//}
	//
	//data := User{
	//	Name: "John Smith",
	//}
	//
	//err = t.Execute(os.Stdout, data)
	//if err != nil {
	//	panic(err)
	//}
	var cheeses = []string{"cheddar", "pep pep"}
	cheese(cheeses)
}

//func cheese(cheeses ...string) {
//	fmt.Println(cheeses)
//}

func cheese(cheeses [] string) {
	fmt.Println(cheeses)
}
