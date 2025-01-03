package grokkingalgorithmsingo

import "fmt"

func main() {
	for ok := true; ok; {
		var option string
		fmt.Printf("Which algorithm do you wanna test:\n")
		fmt.Printf("1 - Binary Search\n")
		fmt.Printf("Selected Option: ")
		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Printf("Selected Binary Search")
		case "2":
			fmt.Printf("Selected Another Algo")
		default:
			fmt.Printf("Select a valid option\n")
		}
	}
}
