package main

import (
	"fmt"
	"math"
)

const (
	a = 12.3
	b = "hello"
	c = 45.6
)

func printAge(name string,age int) {
	fmt.Printf("The Age of %s is %d \n" , name,age)
}

func printAgeFloat(name string ,age float64) {
	fmt.Printf("The Age is in decimal and it is %f of %s \n " ,age, name)
}

//This function illustrates multi return ,single type for input vars if same type is there for both
func multiReturn(name1 ,name2 string) (string,string) {
	//This function returns the salutation in front of the names
    
	name1 = "Sir " + name1
	name2 = "Madam " + name2

	
	return name1,name2
}

//This function illustrates named return for the function calls
func namedReturn(person1,person2 string) (allPeople string)  {
 
	   allPeople = person1 + " and " + person2 +   " are present"
	   return 
}

func main(){

	 //normal print statements
	 fmt.Println(math.Pi*12.3)
	 fmt.Println(a,b,c+12)
	
	 //trying to change the constant can not change
	//  c = 12.3
	//  fmt.Println(c)

	//assigning variables without type , 2 variables at once
	var person1,person2 = "Divyanshu","vartika"
	var age1 = 24
	var age2 = 27.1

	printAge(person1,age1)
	printAgeFloat(person2,age2)

	//using function to llustrate multi returns and type omittion
	fmt.Println(multiReturn(person1,person2))

   //to illustrate named returns

   fmt.Println(namedReturn(person1,person2))

   //short declaration

   person3 := "Kiran"
   age3 := 52

   //type conversions and auto assigning type to the variable based on the RHS value
   var floatAge3 = float64(age3)

   fmt.Printf("%s is %f years old",person3,floatAge3)


}