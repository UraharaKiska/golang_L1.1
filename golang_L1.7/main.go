package main

import (
	"fmt"
)

func setBit(number int64, pos uint) int64{
	number |= 1 << pos
	return number
}

func clearBit(number int64, pos uint) int64 {
	number &^= 1 << pos
	return number
}


func checkError(err error) {
	if err != nil {
		fmt.Println("Try again")
		return
	}
}
func main() {
	var number int64
	var pos uint
	fmt.Println("Enter a number:")
	_, err := fmt.Scan(&number)
	checkError(err)
	fmt.Printf("Bits of %d: %b\n", number, number)
	fmt.Println("Enter a position(from 0) in number to set 1:")
	_, err = fmt.Scan(&pos)
	checkError(err)
	fmt.Printf("%b\n", setBit(number, pos))
	fmt.Println("Enter a position(from 0) in number to set 0:")
	_, err = fmt.Scan(&pos)
	checkError(err)
	fmt.Printf("%b", clearBit(number, pos))
}