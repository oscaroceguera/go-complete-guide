package second

import "fmt"

func Example2() {
	countries := [6]string{"mexico", "japan", "canada", "usa", "brazil", "chile"}

	fmt.Printf("Countries: %v\n", countries) // output: Countries: [mexico japan canada usa brazil chile]

	// implicity 0:2
	fmt.Printf(":2 %v\n", countries[:2]) // output: :2 [mexico japan]
	fmt.Printf("1:3 %v\n", countries[1:3]) // output: 1:3 [japan canada]
	// implicity 2:n
	fmt.Printf("2: %v\n", countries[2:]) // output: 2: [canada usa brazil chile]
	fmt.Printf("2:5 %v\n", countries[2:5]) // output: 2:5 [canada usa brazil]
	fmt.Printf("0:3 %v\n", countries[0:3]) // output: 0:3 [mexico japan canada]
	fmt.Printf("Last element: %v\n", countries[4]) // output: Last element: brazil
	fmt.Printf("Last element: %v\n", countries[len(countries)-1]) // output: Last element: chile
	fmt.Printf("Last element: %v\n", countries[4:]) // output: Last element: [brazil chile]
	fmt.Printf("All elements: %v\n", countries[0:len(countries)]) // output: All elements: [mexico japan canada usa brazil chile]
	fmt.Printf("Last three elements: %v\n", countries[3:len(countries)]) // output: Last three elements: [usa brazil chile]
	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)]) // output: Last two elements: [brazil chile]
	fmt.Println(countries[:]) // output: [mexico japan canada usa brazil chile]
	fmt.Println(countries[0:]) // output: [mexico japan canada usa brazil chile]
	fmt.Println(countries[0:len(countries)]) // output: [mexico japan canada usa brazil chile]

	a := []int{1,2}
	b := []int{11,22}
	fmt.Println(append(a, b...)) // [1 2 11 22]

	c := []int{1,2}
	fmt.Println(append(c, c...)) // [1 2 1 2]

}