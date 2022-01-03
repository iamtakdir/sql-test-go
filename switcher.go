package main

import "fmt"

func switcher() {
	var take int
	fmt.Printf(`
Enter 1 for all data
Enter 2 for inserting data
Enter 3 for update data
Enter 4 for delete data

`)
	fmt.Scanln(&take)

	switch {

	case take == 1:
		//get_all()

	case take == 2:
		var key int
		var value string
		fmt.Println("enter a id")
		fmt.Scanln(&key)
		fmt.Println("enter a name")
		fmt.Scanln(&value)

		insert_into(key, value)

	case take == 3:
		//get_all()
		var key int
		var value string
		fmt.Println("enter a id")
		fmt.Scanln(&key)
		fmt.Println("enter a name")
		fmt.Scanln(&value)
		update_data(key, value)

	case take == 4:
		//get_all()

		var key int
		fmt.Println("enter id that you want to delete ?")
		fmt.Scanln(&key)

		delete_data(key)

	}

}
