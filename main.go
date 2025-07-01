package main

import (
	"fmt"
	user "package_contactapp/User"
)

func main() {
	user1, err := user.CreateNewAdmin("Vishav", "Pathania")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user1)
	// Admin cannot create new contacts!
	user1contact1, err := user1.AddNewContact("Yash", "Shah")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user1contact1)
	user2, err := user1.CreateNewStaff("Aniket", "Pardeshi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user2)
	// creating contact from staff
	user2contact1, err := user2.AddNewContact("Brijesh", "Mavani")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user2contact1)
	// updating contacts first name
	err = user2.UpdateContactById(1, "F_name", "brijesh")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user2contact1)
	// adding new contactdetails to user2contact1
	user2contact1detail1, err := user2contact1.AddNewContact_Details("Email", "Brijesh@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user2contact1detail1)
	fmt.Println(*user2contact1)
	// deleting contact_details
	err = user2.DeleteContact_DetailsById(1, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user2contact1)
	// deleting contact
	err = user2.DeleteContactById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user2contact1)
}
