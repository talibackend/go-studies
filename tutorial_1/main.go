package main

// Just gonna be studying basic syntax

import "fmt";
import "errors";

func main(){
	var intNum int = 900000;
	var floatNum float32 = 70000.9;
	var casted int = int(floatNum)

	fmt.Println("Hello, world...");
	fmt.Println(intNum);
	fmt.Println(floatNum);
	fmt.Println(casted);

	var name string = "Femi"

	fmt.Println(name +  " Fatokun");

	var lie bool = false;

	fmt.Println(lie);

	const dontChange int32 = 6089807;
	
	fmt.Println(dontChange);

	testFunc();

	// var arithFuncCall,  err = arithFunc(5, 6);
	var arithFuncCall,  err = arithFunc(0, 6);

	if(err != nil){
		fmt.Printf("Operation failed with error: %v\n", err);
	}else{
		fmt.Printf("Arithmethic function returned %v\n", arithFuncCall);
	}

	var age int32 = 18;
	// var age string = "Don't wanna disclose";

	switch{
		case age < 18:
			fmt.Printf("You are not an adult\n");
		case age >= 18:
			fmt.Printf("You are an adult!\n");
		default:
			fmt.Printf("Invalid date passed...\n");
	}
}

func testFunc() {
	fmt.Println("Printing works well from another function....");
}

func arithFunc(a int32, b int32) (int32, error) {
	var err error;
	if(a == 0 || b == 0){
		err = errors.New("Both variables must be greater than 0");
		return 0, err;
	}
	return a * b, err;
}