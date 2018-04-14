package main

import "fmt"

func main(){
	nums :=[]int{0,1,2,3,4,5,6,7,8,9,10}

	for i,_:= range nums{
		if nums[i] %2 == 0 {
		fmt.Printf("%v  is even\n",nums[i])
		} else{ 
			fmt.Printf("%v is odd\n",nums[i])
		}
	}
}
