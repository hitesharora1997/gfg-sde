package main

import "fmt"

func checkPalindrome(num int) bool {
	temp := num
	var rev int
	for temp != 0 {
		j := temp % 10

		rev = rev*10 + j

		temp = temp / 10

	}
	return (rev == num)
}

func main() {
	var num int
	fmt.Println("Enter the number to check the palindrome")
	fmt.Scanf("%d", &num)

	if true == checkPalindrome(num) {
		fmt.Println("is a palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}

}
