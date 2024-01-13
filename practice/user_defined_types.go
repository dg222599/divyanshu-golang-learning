package practice

import (
	"fmt"
)


type Person struct {
	Name string
	Age int
}


func ExtraTypes () {
	  
	 i,j:=42,21


	 pointerI := &i

	 fmt.Print(i," ",j," ",pointerI,"   ",*pointerI)

	//  *pointerI = 34

	//  fmt.Print(i," ",j," ",pointerI,"   ",*pointerI)


	 person := Person{"Div",23}

	 person.Name = "Var"

	 pointerPerson := &person

	 pointerPerson.Age = 34

	 fmt.Println(person)


	 // slices and arrays
   
	var arr[2]int = [2]int{1,2}
    // arr[0] = 1
	// arr[1] = 2


	fmt.Println(arr)


	slice := arr[0:1]

	//var slice1[]int

	slice1 := []int{1,2,3,4}

	slice1 = []int{1,2}

	fmt.Println(slice1)

	fmt.Println(slice)


	//maps


	mp := make(map[string]int)
	mp["Divyanshu"] = 23

	fmt.Println(mp["Divyanshu"])

	mp1:=make(map[string]Person)

	mp1["Divyanshu"] = Person{Name: "Vartika",Age: 24}

	fmt.Println(mp1)


	mp2 := make(map[int]float64)

	mp2[2] = 4.5

	mp2[3] = 9.0


	fmt.Println(mp2[2])

	delete(mp2,2)


	key,ok := mp2[2]

	fmt.Println(key , " " , ok)



	


}