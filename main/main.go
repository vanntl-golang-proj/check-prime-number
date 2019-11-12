package main

import (
	"../util"
	"fmt"
)

func hello(){
	print("welcome to my program \n")
}

func main(){
	var num int

	hello()

	fmt.Printf("Input number to check prime:")

	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Print(err)
		return
	}

	var isPrime = utils.IsPrime(num)
	if isPrime{
		fmt.Printf("%d is a prime", num)
	} else{
		fmt.Printf("%d is not a prime", num)
	}

}
