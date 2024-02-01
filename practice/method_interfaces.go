package practice

import (
	"fmt"
)

type Vehicle struct {
	Name string
	Year int
	Color string
}

func (vehicle Vehicle) printVehicleDetails() {
	fmt.Printf("Vehicle Name is %s\n",vehicle.Name)
	fmt.Printf("Vehicle Color is %s\n",vehicle.Color)
	fmt.Printf("Vehicle year of make is %d\n",vehicle.Year)
}

type myInt int

func (a myInt) printSum() {
	fmt.Print(a+2)


}

func Interfaces(){
	 fmt.Println("Hello from Interfaces!!")
	 vehicle1 := Vehicle{
		Name:"Creta",
		Year: 2023,
		Color: "Black",
	 }

	 vehicle1.printVehicleDetails()

	 a:=myInt(2)
	 a.printSum()


}