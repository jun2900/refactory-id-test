package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func isPalindrome(input string) string {
	var result = "palindrome"
	input = strings.ToLower(input)
	input = strings.Replace(input, " ", "", -1)

	var limit = math.Floor(float64(len(input) / 2))

	for i := 0; i < int(limit); i++ {
		if string(input[i]) != string(input[len(input)-i-2]) {
			result = "not palindrome"
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	result := isPalindrome(input)
	fmt.Println(result)
}
