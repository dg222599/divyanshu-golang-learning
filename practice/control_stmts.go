package practice

import (
	"fmt"
	"runtime"
)

func Control () {

	 //we will be adding if-else , switch  ,for , defer

	 defer fmt.Println("\nand the party has ended and it has ended with Defer!!")

	 //if-else

	 gender := "male"

	 if gender = "female" ; gender == "male" {
		fmt.Println("We have got a Male among ourselves!!")
	 } else {
		fmt.Println("We have got a female among ourselves today!!")
	 }

	 //for,while

	 for i:=1;i<10;i++ {
		if y:=2*i ; y%4==0 {
			fmt.Printf("We have got someone from the family of 4 and he is  %d\n" , y)
		}
	 }


	 //while

	 prev1 :=1
	 prev2 :=2
     fmt.Print(prev1," ",prev2)

	 for prev1 <10 {
		 ans := prev1 + prev2
         fmt.Print(ans," ")
		 prev1 = prev2

		 prev2 = ans

	 }

	


	 //switch

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}


	 
}

