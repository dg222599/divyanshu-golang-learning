package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterEven(t *testing.T){
	 //given
	 givenNumbers:=[]int{1,2,3,4,5,6,7,8,9,10,11,13,14,20}
     wantNumbers:=[]int{2,4,6,8,10,14,20}

	 gotNumbers := FilterEven(givenNumbers)

	 assert.Equal(t,gotNumbers,wantNumbers)
}


func TestFilterOdd(t *testing.T) {
	 //given
	 givenNumbers:=[]int{1,2,3,4,5,6,7,8,9,10,11,13,14,20}
     wantNumbers:=[]int{1,3,5,7,9,11,13}

	 gotNumbers := FilterOdd(givenNumbers)

	 assert.Equal(t,gotNumbers,wantNumbers)
}

func TestFilterPrime(t *testing.T) {
	 //given
	 givenNumbers:=[]int{1,2,3,4,5,6,7,8,9,10,11,13,14,20}
     wantNumbers:=[]int{2,3,5,7,11,13}

	 gotNumbers := FilterPrime(givenNumbers)

	 assert.Equal(t,gotNumbers,wantNumbers)
}


func TestCheckPrime(t *testing.T) {
	want:=false
	got:=CheckPrime(8)

	assert.Equal(t,want,got)
}

func TestCheckMultipleN(t *testing.T) {
	number :=45

	wantResult:=false
	gotResult:=CheckMultipleN(number,4)

	assert.Equal(t,wantResult,gotResult)
}

func TestSatisfyAnyCondition(t *testing.T){
     odd:=func (n int) bool {return n%2==1}
	 even:=func (n int) bool {return n%2==0}
	 greaterthanN:=func(n int) Condition {return func(m int)bool{return m>n}}
	 //lessthanN:=func(n int) Condition {return func(m int)bool{return (m<n)}}
	 greaterthan5:=greaterthanN(5)
	 //lessthan15:=lessthanN(15)
	 multiplesof:=func(n int) Condition{return func(m int)bool{return (m%n==0)}}
	 multiplesof3 := multiplesof(3)

	 testcases:=[]struct{
		nums []int
		conds []Condition
		want []int
	 }{
		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds: []Condition{odd,greaterthan5,multiplesof3},
			want: []int{1,3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds: []Condition{even,multiplesof3},
			want: []int{2,3,4,6,8,9,10,12,14,15,16,18,20},
		},
	 }

	 for _,testcase := range testcases {
		got:= SatisfyAnyCondition(testcase.nums,testcase.conds)
		assert.Equal(t,got,testcase.want)
	 }
}

func TestSatsifyAllCondition(t *testing.T) {
	odd:=func (n int) bool {return n%2==1}
	even:=func (n int) bool {return n%2==0}
	greaterthanN:=func(n int) Condition {return func(m int)bool{return m>n}}
	lessthanN:=func(n int) Condition {return func(m int)bool{return (m<n)}}
	greaterthan5:=greaterthanN(5)
	lessthan15:=lessthanN(15)
	multiplesof:=func(n int) Condition{return func(m int)bool{return (m%n==0)}}
	multiplesof3 := multiplesof(3)

	testcases:=[]struct{
	   nums []int
	   conds []Condition
	   want []int
	}{
	   {
		   nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		   conds: []Condition{odd,greaterthan5,multiplesof3},
		   want: []int{9,15},
	   },
	   {
		   nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		   conds: []Condition{even,multiplesof3,lessthan15},
		   want: []int{6,12},
	   },
	}

	for _,testcase := range testcases {
	   got:= SatisfyAllCondition(testcase.nums,testcase.conds)
	   assert.Equal(t,got,testcase.want)
	}
}