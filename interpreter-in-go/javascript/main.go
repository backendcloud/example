package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {

	vm := otto.New()

	// Run something in the VM
	vm.Run(`
    	abc = 2 + 2;
    	console.log("The value of abc is " + abc); // 4
	`)

	// Get a value out of the VM
	if value, err := vm.Get("abcdef"); err == nil {
		if value_int, err := value.ToInteger(); err == nil {
			fmt.Printf("%v, %v\n", value_int, err)
		}
	}

	// Set a number
	vm.Set("def", 11)
	vm.Run(`
    	console.log("The value of def is " + def);
    	// The value of def is 11
	`)

	// Set a string
	vm.Set("xyzzy", "Nothing happens.")
	vm.Run(`
    	console.log(xyzzy.length); // 16
	`)

	// Get the value of an expression
	value, _ := vm.Run("xyzzy.length")
	{
		// value is an int64 with a value of 16
		valueInt, _ := value.ToInteger()
		fmt.Printf("var valueInt is %v\n", valueInt)
	}

	// An error happens
	_, err := vm.Run("abcdefghijlmnopqrstuvwxyz.length")
	if err != nil {
		// err = ReferenceError: abcdefghijlmnopqrstuvwxyz is not defined
		// If there is an error, then value.IsUndefined() is true
		fmt.Println("Err happened!")
	}
}
