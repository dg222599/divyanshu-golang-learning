package exercises

import (
	"math"
)

type Condition func(n int) bool

//function to check whether the number is even or not
func CheckEven(number int) bool {
	if number % 2 ==0 {
		return true
	} else {
		return false
	}
}

//func to check number is prime or not
func CheckPrime(number int) bool {
	if number == 1{
		return false
	} else if number ==2 {
		return true
	}

	currentDivisor :=2

	for currentDivisor <= int(math.Floor(math.Sqrt(float64(number)))) {
		
		if number % currentDivisor == 0 {
			return false
		} else {
			 currentDivisor+=1
		}

	}

	return true
}

func CheckMultipleN(number int,N int) bool {
	if number % N ==0 {
		return true
	} else {
		return false
	}
}

//filter odd numbers from the list
func FilterOdd(numbers []int) []int {
	output:=[]int{}
	for _,number := range numbers {
		if ! CheckEven(number) {
           output = append(output, number)
		}
	}

	return output
}

//filter even numbers from the list
func FilterEven(numbers []int) []int {

	output:=[]int{}
	for _,number := range numbers {
		if CheckEven(number){
           output = append(output, number)
		}
	}

	return output
}

//filter Prime numbers
func FilterPrime(numbers []int) []int{
	output := []int{}
 
	for _,number := range numbers {
		
		if CheckPrime(number)  {
			output = append(output, number)
		}
		
	}

	return output
}

func FilterOddPrime(numbers []int) []int{
	output := []int{}

	for _,number := range numbers {

		if CheckPrime(number) && !CheckEven(number){
			output = append(output, number)
		}
	}

	return output
}

func EvenMultipleFive(numbers []int) []int{
	
	output := []int{}

	for _,number := range numbers {

		if CheckEven(number) && CheckMultipleN(number,5) {
			output = append(output, number)
		}
	}

	return output
}

func OddMultipleThreeGreaterThanTen(numbers []int) []int{
	output:=[]int{}

	for _,number := range numbers {

		if !CheckEven(number) && CheckMultipleN(number,3) && number>10{
			output = append(output, number)
		}
	}

	return output
}

func SatisfyAllCondition(numbers []int,conditions []Condition) []int{

   output := []int{}	 
	
   for _,number := range numbers {
	conditionsMatched := true
	    for _,conds := range conditions {
			if ! conds(number){
				conditionsMatched = false
				break
			}
		}
	 
		 if conditionsMatched {
			output = append(output, number)
		 }
   }

   return output

}

func SatisfyAnyCondition(numbers []int,conditions []Condition) []int{

	 output := []int{}

	 for _,number := range numbers {
		conditionsMatched := false
			for _,conds := range conditions {
				if conds(number) {
					conditionsMatched = true
					break
				}
			}
		 
			 if conditionsMatched {
				output = append(output, number)
			 }
	}
	

	 return output
}

