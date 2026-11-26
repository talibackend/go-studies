package main

import "fmt";

type Human struct{
	name string
	age uint
	height float32
	weight float32
	personality string
}

func (b Human) isAdult() bool {
	if(b.age >= 18){
		return true;
	}

	return false;
}

func main(){
	var femi Human = Human{ name : "Femi", age : 24, height : 1.6, weight : 78.92, personality: "jovial" }

	fmt.Println(femi);
	fmt.Println(femi.isAdult());
}