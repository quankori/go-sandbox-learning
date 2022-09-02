package basic

import "fmt"

func maps() {
	var personSalary map[string]int
	if personSalary == nil {
		personSalary = make(map[string]int)
		fmt.Println(personSalary) // Return map[]

		// Add value 1
		personSalary["steve"] = 12000
		personSalary["jamie"] = 15000
		personSalary["mike"] = 9000
		fmt.Println("personSalary map contents:", personSalary)

		// Add value 2
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		fmt.Println("personSalary map contents:", personSalary)

		// Loop
		for key, value := range personSalary {
			fmt.Printf("personSalary[%s] = %d\n", key, value)
		}

		// Length
		fmt.Println("length is", len(personSalary))

		// Remove
		delete(personSalary, "steve")
		fmt.Println("personSalary map contents:", personSalary)

		// Reference
		newPersonSalary := personSalary
		newPersonSalary["jamie"] = 18000
		fmt.Println("Person salary changed", personSalary)
	}
}
