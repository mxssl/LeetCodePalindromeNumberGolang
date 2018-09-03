package main

import "fmt"

func main() {
	num := 1221

	fmt.Println(isPalindrome(num))
}

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	revertedNumber := 0

	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	if x == revertedNumber || x == revertedNumber/10 {
		return true
	}

	return false
}
