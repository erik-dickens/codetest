package main

import "fmt"

func getDeepestLetter(input string) rune {
	depthArr := make([]int, 0)
	var depth, maxDepth int = 0, 0
	for i := 0; i < len(input); i++ {
		// If end of parenthesis, step down one level
		if input[i] == ')' {
			depth--
		}

		// Error handling
		if depth < 0 {
			break
		}

		// Assign depth
		depthArr = append(depthArr, depth)
		if depth > maxDepth {
			maxDepth = depth
		}

		// If beginning of new parenthesis, next letter will have deeper level
		if input[i] == '(' {
			depth++
		}
	}

	var deepestLetter rune
	for i, v := range depthArr {
		if v == maxDepth {
			if deepestLetter != 0 {
				fmt.Println("More than one deepest letter")
				deepestLetter = '?'
				break
			}
			if input[i] < 'a' || input[i] > 'z' {
				fmt.Println("Character must be a-z (lower case)")
				deepestLetter = '?'
				break
			}
			deepestLetter = rune(input[i])
		}
	}
	
	fmt.Printf("input: %v\n", input)
	fmt.Printf("depthArr: %v\n", depthArr)
	fmt.Printf("deepestLetter: %v\ndepth: %v\n", string(deepestLetter), maxDepth)
	var a, z rune = 'a', 'z'
	fmt.Printf("%v, %v", a, z)
	if depth != 0 {
		fmt.Println("Malformed string")
	}

	return rune(deepestLetter)
}

type letterTestCase struct {
	input    string
	expected rune
}

// func main() {
// 	//TODO:implement here
// 	getDeepestLetter("d(((a)))(j(((d))))kjf")
// }
