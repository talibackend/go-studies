package main

import "fmt";

func main(){
	var pointer *int32;
	fmt.Println(pointer);
	pointer = new(int32);
	fmt.Println(pointer);
	fmt.Printf("The value at location %v is %v\n", pointer, *pointer);
	*pointer = 70;
	fmt.Printf("The value at location %v is %v\n", pointer, *pointer);

	var num int;
	var pointer_2 = &num;

	fmt.Printf("Value at memory %v is %v\n", pointer_2, num);

	*pointer_2 = 30;

	fmt.Printf("Value at memory %v is %v\n", pointer_2, num);

	sliceCopy();

	var testSlice = [11]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11};
	fmt.Printf("Location of fresh array is %p\n", &testSlice);
	practicalPointers(testSlice);
	fmt.Println(testSlice);
	// fmt.Printf("%p", &result);
}

func sliceCopy(){
	var numArr []int = []int{1, 2, 3, 4, 5, 6};
	fmt.Println(numArr);
	var numArr2 = numArr;
	numArr2[0] = 9075956;
	fmt.Println(numArr);
}

// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this
// I failed this

func practicalPointers(nums [11]uint){
	fmt.Printf("Location of fresh 2 array is %p\n", &nums);

	for i := 0; i < len(nums); i++{
		nums[i] = nums[i] * nums[i];
	}

	fmt.Println(nums);
	fmt.Printf("Location of new array is %p\n", &nums);
}
