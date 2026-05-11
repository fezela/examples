package main

import "fmt"


func main() {
	
	i := 1
	//whimsy := true	
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for { //This is how you do while in go
		whimsy := true
	 	if whimsy {
			fmt.Println("whmisy == true")
			}
		for j := 0; j < 5; j++{
			fmt.Println("in my whimsy state")
		}
		whimsy = false
		break
	}

	for { //Without the break this is an infinite loop
	
		fmt.Println("If the code that follows this one wasn't BREAK this would run forever")
		break
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for n := range 6 {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
