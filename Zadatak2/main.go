/*2. Declare two variables, `firstName` and `lastName` assign them with wanted values.
Declare constant named `fullname`
Combine constant and both strings into a full name by concatenating strings with a space in between and print them out.*/

package main

import "fmt"

func main() {
	var firstName string
	firstName = "Damir"

	var lastName string
	lastName = "Mihajlović"

	//Nemoguće je koristit konstantu jer first i last name su varijable
	var fullname string
	fullname = firstName + " " + lastName

	fmt.Println(fullname)

}
