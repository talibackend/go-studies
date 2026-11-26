package main;

import "fmt";

func main(){
	var intArr [3]uint64 = [3]uint64{1, 2, 3};

	fmt.Println(intArr);
	intArr[2] = 67894;
	fmt.Println(intArr[2]);
	fmt.Println(intArr);
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	slices();
	maps();
	loops();
}

func slices(){
	var testSlice []int64 = []int64{1, 2, 3, 4, 5, 6, 7, 8};
	fmt.Println(testSlice);
	fmt.Printf("The length of the array is %v and the capacity is %v\n", len(testSlice), cap(testSlice));

	testSlice = append(testSlice, 89)
	fmt.Printf("The length of the array is %v and the capacity is %v\n", len(testSlice), cap(testSlice));
}

func maps(){
	var testMap = map[string]uint8 { "Femi" : 24, "Praise" : 23 };
	var age, ok = testMap["Praise"];

	if(ok == false){
		fmt.Println("Record not found...");
	}else{
		fmt.Println(age);
	}
}

func loops(){
	var intArr []int32 = []int32{7, 2, 3, 4, 5, 6, 1, 8, 9, 0}

	fmt.Println(intArr);

	for i:= range intArr{
		fmt.Println(intArr[i]);
	}

	var testMap = map[string]string { "name" : "Femi", "age" : "17", "major" : "cs" }

	for info:= range testMap{
		fmt.Printf("%v: %v\n", info, testMap[info]);
	}

	for i := 0; i < 3; i++{
		fmt.Println(i);
	}
}