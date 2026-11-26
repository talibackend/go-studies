package main

import "fmt";
import "time";

func main(){
	var limit int = 1000000;
	var testSlice = []int{};
	var testSlice2 = make([]int, 0, limit);

	fmt.Printf("Time spent when preallocation is disabled is %v\n", timeLoops(testSlice, limit))
	fmt.Printf("Time spent when preallocation is enabled is %v\n", timeLoops(testSlice2, limit))
}

func timeLoops(slice []int, n int) time.Duration {
	var starTime = time.Now();

	for len(slice) < n {
		slice = append(slice, 1);
	}

	return time.Since(starTime);
}