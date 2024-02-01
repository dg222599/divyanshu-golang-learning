package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T){

	
	 t.Run("saying Hello to People",func(t *testing.T){
		got:=Hello("Chris")
		want:="Hello Chris !!"
		assert.Equal(t,got,want)
	 })

	 t.Run("Checking Empty String",func(t *testing.T){
		got:=Hello("") //sending empty string
		want:="Hi There!!"
		assert.Equal(t,got,want)
	 })

	 
}