// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {

	hobbies := [3]string{"doing code", "listening music", "playing games"}
	fmt.Println("My Hobbies list: ", hobbies)
	fmt.Println("My First hobbie: ", hobbies[0])
	lastTwoHobbies := hobbies[1:]
	fmt.Println("My Second and thired hobbies: ", lastTwoHobbies)
	firstHobbie1 := hobbies[:2]
	firstHobbie2 := firstHobbie1[:2]

	fmt.Println(firstHobbie1, firstHobbie2, len(firstHobbie1), cap(firstHobbie1), len(firstHobbie2), cap(firstHobbie2))

	LastTwoHobbie := firstHobbie1[1:3]
	fmt.Println(LastTwoHobbie)

	goals := []string{"learn backend", "build api"}

	fmt.Println(goals)

	goals[1] = "some thing new"
	goals = append(goals, "thired goal")
	fmt.Println(goals)

	product1 := Product{
		title: "product1",
		id:    01,
		price: 1.0,
	}

	product2 := Product{
		title: "product2",
		id:    02,
		price: 2.0,
	}

	productList := []Product{product1, product2}

	fmt.Println(productList)
	product3 := Product{
		title: "product3",
		id:    03,
		price: 3.0,
	}
	productList = append(productList, product3)
	fmt.Println(productList)

}
