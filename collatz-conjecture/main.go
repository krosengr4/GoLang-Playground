package main

import "fmt"

/*
- Any given number(n) will eventually reach 1 when following the rules below
- If n is even: n / 2
- If n is odd: 3n + 1

- Given n, return how many transformations it takes to get to 1.
- Ex: 5. Answer 5. 5, 16, 8, 4, 2, 1
*/

func calculateTransformations(n int) int {
	fmt.Println(n)

	i := 0
	ifContinue := true

	for ifContinue {
		fmt.Println(n)
		if n == 1 {
			ifContinue = false
			break
		} else if n%2 == 0 {
			n /= 2
		} else {
			n *= 3
			n += 1
		}

		i++
	}

	return i
}

func main() {
	n := 5
	result := calculateTransformations(n)

	fmt.Println("Result:", result)
}
