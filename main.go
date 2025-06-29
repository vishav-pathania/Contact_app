// Thought: isAdmin bool(can only incorporate 2 roles) it will be better if string or enum so in future more roles can be added easily
package main

import (
	"fmt"
	"reflect"
)

var Contact_Detailsmap = make(map[int]*Contact_Details)
var Contact_Detailsmapid = -1
var Contactmap = make(map[int]*Contact)
var Contactmapid = -1
var Usermap = make(map[int]*User)
var Usermapid = -1

type Contact_Details struct {
	Contact_Details_ID int
	Type               string
	Number             string
	Email              string
}

type Contact struct {
	Contact_ID      int
	F_name          string
	L_name          string
	isActive        bool
	Contact_Details []*Contact_Details
}

type User struct {
	User_ID  int
	F_name   string
	L_name   string
	isAdmin  bool
	isActive bool
	Contacts []*Contact
}

func NewContact_Details(Type, NumberorEmail string) (*Contact_Details, error) {
	//i think check for contact details id is not required as will never be passed by user!
	// isActive bhi hamesha true hona chahiye creating ke vakat toh params mai lene ki jarurat nhi
	if Type == "" || (Type != "Number" && Type != "Email") {
		return nil, fmt.Errorf("Invalid Type!")
	}
	Contact_Detailsmapid++
	cd := &Contact_Details{
		Contact_Details_ID: Contact_Detailsmapid,
		Type:               Type,
	}
	Contact_Detailsmap[Contact_Detailsmapid] = cd

	if Type == "Email" {
		cd.Email = NumberorEmail
	} else {
		cd.Number = NumberorEmail
	}

	return cd, nil

}

func NewContact(F_name string, L_name string) (*Contact, error) {
	if F_name == "" {
		return nil, fmt.Errorf("First Name Cannot be Empty!")
	}
	if L_name == "" {
		return nil, fmt.Errorf("Last Name Cannot be Empty!")
	}

	Contactmapid++
	c := &Contact{
		Contact_ID:      Contactmapid,
		F_name:          F_name,
		L_name:          L_name,
		isActive:        true,
		Contact_Details: []*Contact_Details{},
	}
	Contactmap[Contactmapid] = c
	return c, nil

}

func NewUser(F_name, L_name string, isAdmin bool) (*User, error) {
	if F_name == "" {
		return nil, fmt.Errorf("First Name Cannot be Empty!")
	}
	if L_name == "" {
		return nil, fmt.Errorf("Last Name Cannot be Empty!")
	}
	Usermapid++
	u := &User{
		User_ID:  Usermapid,
		F_name:   F_name,
		L_name:   L_name,
		isAdmin:  isAdmin,
		isActive: true,
		Contacts: []*Contact{},
	}
	Usermap[Usermapid] = u
	return u, nil
}

func (U *User) CreateNewUser(F_name, L_name string, isAdmin bool) (*User, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("You Don't have admin previlages to create a User!")
	}
	anewUser, err := NewUser(F_name, L_name, isAdmin)
	if err != nil {
		return nil, err
	}
	return anewUser, nil
}

func (U *User) GetAllSystemUser() ([]*User, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("You Don't have admin previlages to Read all Users!")
	}
	AllUsersSlice := []*User{}
	for _, val := range Usermap {
		AllUsersSlice = append(AllUsersSlice, val)
	}
	return AllUsersSlice, nil
}

func (U *User) GetUserById(id int) (*User, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("You Don't have admin previlages to Read User!")
	}
	UserbyId, ok := Usermap[id]
	if !ok {
		return nil, fmt.Errorf("User with id: %d not found!", id)
	}
	return UserbyId, nil
}

func (U *User) DeleteUserById(id int) error {
	if !U.isAdmin {
		return fmt.Errorf("You Don't have admin previlages to Delete User!")
	}
	UserWithMatchingId, err := U.GetUserById(id)
	if err != nil {
		return err
	}
	// delete(Usermap, id)
	UserWithMatchingId.isActive = false
	return nil
}

func (C *Contact) GetContactById(id int) (*Contact, error) {
	ContactbyId, ok := Contactmap[id]
	if !ok {
		return nil, fmt.Errorf("Contact with id: %d not found!", id)
	}
	return ContactbyId, nil
}

func (C *Contact) GetContact_DetailsById(id int) (*Contact_Details, error) {
	Contact_DetailsbyId, ok := Contact_Detailsmap[id]
	if !ok {
		return nil, fmt.Errorf("Contact_Details with id: %d not found!", id)
	}
	return Contact_DetailsbyId, nil
}

// func (C *Contact) GetContactByIndex(id int, contactSlice []*Contact) (*Contact, error) {
// 	for _, val := range contactSlice {
// 		if val.Contact_ID == id {
// 			return val, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("Contact with id: %d not found in User Contact slice!", id)
// }

func (C *Contact) DeleteContactById(id int) error {
	ContactWithMatchingId, err := C.GetContactById(id)
	if err != nil {
		return err
	}
	// _, err = C.DeleteContactByIndex()
	// delete(Contactmap, id)
	ContactWithMatchingId.isActive = false
	return nil
}

func (C *Contact_Details) GetContact_DetailsByIndex(id int, contactSlice []*Contact_Details) (*Contact_Details, error) {
	for _, val := range contactSlice {
		if val.Contact_Details_ID == id {
			return val, nil
		}
	}
	return nil, fmt.Errorf("Contact_Details with id: %d not found in User Contact slice!", id)
}

func (C *Contact) DeleteContact_DetailsById(id int, SliceOnWhichPerformDelete []*Contact_Details) error {
	_, err := C.GetContact_DetailsById(id)
	if err != nil {
		return err
	}
	delete(Contact_Detailsmap, id)
	err = C.DeleteContact_DetailsByIndex(id, SliceOnWhichPerformDelete)
	if err != nil {
		return err
	}
	return nil
}

func (C *Contact) DeleteContact_DetailsByIndex(id int, Contact_DetailsSlice []*Contact_Details) error {
	NewContact_DetailsSliceToSet := []*Contact_Details{}
	flag := false
	for _, val := range Contact_DetailsSlice {
		if val.Contact_Details_ID != id {
			NewContact_DetailsSliceToSet = append(NewContact_DetailsSliceToSet, val)
		}
		flag = true
	}
	if !flag {
		return fmt.Errorf("Not found in Contacts Slice of Contact_Details with id: %d !", id)
	}
	Contact_DetailsSlice = NewContact_DetailsSliceToSet
	return nil
}

func getVariableType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func (U *User) UpdateUser(UserToBeUpdated *User, param string, value interface{}) error {

	switch param {
	case "F_name":
		err := UserToBeUpdated.UpdateFname(value)
		if err != nil {
			return err
		}
		return nil
	case "L_name":
		err := UserToBeUpdated.UpdateLname(value)
		if err != nil {
			return err
		}
		return nil
	case "isAdmin":
		err := UserToBeUpdated.UpdateisAdmin(value)
		if err != nil {
			return err
		}
		return nil
	case "isActive":
		err := UserToBeUpdated.UpdateisActive(value)
		if err != nil {
			return err
		}
		return nil
	// case "Contact":
	// 	err := UserToBeUpdated.UpdateContactId(contactId, value)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	default:
		return fmt.Errorf("No matching params founnd to Update!")
	}
}

func (C *Contact) UpdateContactById(id int, param string, value interface{}) error {
	ContacWithMatchingId, err := C.GetContactById(id)
	if err != nil {
		return err
	}
	err = ContacWithMatchingId.UpdateContact(param, value)
	if err != nil {
		return err
	}
	return nil
}

func (C *Contact) UpdateContact(param string, value interface{}) error {
	switch param {
	case "F_name":
		err := C.UpdateFname(value)
		if err != nil {
			return err
		}
		return nil
	case "L_name":
		err := C.UpdateLname(value)
		if err != nil {
			return err
		}
		return nil
	case "isActive":
		err := C.UpdateisActive(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("No matching params founnd to Update!")

	}
}

func (U *User) UpdateFname(value interface{}) error {
	if !U.isAdmin {
		return fmt.Errorf("You Don't have admin previlages to Update User First Name!")
	}
	if getVariableType(value) != "string" {
		return fmt.Errorf("Please Enter a string value!")
	}
	if value == "" {
		return fmt.Errorf("F_name cannot be Empty!")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("Error in setting F_name string!")
	}
	U.F_name = conval
	return nil
}

func (U *User) UpdateLname(value interface{}) error {
	if !U.isAdmin {
		return fmt.Errorf("You Don't have admin previlages to Update User Last Name!")
	}
	if getVariableType(value) != "string" {
		return fmt.Errorf("Please Enter a string value!")
	}
	if value == "" {
		return fmt.Errorf("L_name cannot be Empty!")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("Error in setting L_name string!")
	}
	U.L_name = conval
	return nil
}

func (U *User) UpdateisAdmin(value interface{}) error {
	if !U.isAdmin {
		return fmt.Errorf("You Don't have admin previlages to Update User role!")
	}
	if getVariableType(value) != "bool" {
		return fmt.Errorf("Please Enter a boolean value!")
	}
	conval, ok := value.(bool)
	if !ok {
		return fmt.Errorf("Error in setting isAdmin boolean!")
	}
	U.isAdmin = conval
	return nil
}

func (U *User) UpdateisActive(value interface{}) error {
	if !U.isAdmin {
		return fmt.Errorf("You Don't have admin previlages to Update User Active Status!")
	}
	if getVariableType(value) != "bool" {
		return fmt.Errorf("Please Enter a boolean value!")
	}
	conval, ok := value.(bool)
	if !ok {
		return fmt.Errorf("Error in setting isActive boolean!")
	}
	U.isActive = conval
	return nil
}

// func (U *User) UpdateContactId(id int, value interface{}) error {
// 	if getVariableType(value) != "int" {
// 		return fmt.Errorf("Please Enter a int value for Id!")
// 		conval, ok := value.(bool)
// 		if !ok {
// 			return fmt.Errorf("Error in setting isActive boolean!")
// 		}
// 		U.isActive = conval
// 		return nil
// 	}
// }

func (U *User) GetAllSystemContacts() ([]*Contact, error) {
	// if !U.isAdmin {
	// 	return nil, fmt.Errorf("You Don't have admin previlages to Read all Contacts!")
	// }
	AllContactsSlice := []*Contact{}
	for _, val := range Contactmap {
		AllContactsSlice = append(AllContactsSlice, val)
	}
	return AllContactsSlice, nil
}

func (U *User) GetAllSystemContact_Details() ([]*Contact_Details, error) {
	// if !U.isAdmin {
	// 	return nil, fmt.Errorf("You Don't have admin previlages to Read all Contact_Details!")
	// }
	AllContact_DetailssSlice := []*Contact_Details{}
	for _, val := range Contact_Detailsmap {
		AllContact_DetailssSlice = append(AllContact_DetailssSlice, val)
	}
	return AllContact_DetailssSlice, nil
}

func (U *User) AddNewContact(F_name, L_name string) (*Contact, error) {
	anewContact, err := NewContact(F_name, L_name)
	if err != nil {
		return nil, err
	}
	U.Contacts = append(U.Contacts, anewContact)
	return anewContact, nil
}

func (C *Contact) UpdateFname(value interface{}) error {
	if getVariableType(value) != "string" {
		return fmt.Errorf("Please Enter a string value!")
	}
	if value == "" {
		return fmt.Errorf("F_name cannot be Empty!")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("Error in setting F_name string!")
	}
	C.F_name = conval
	return nil
}

func (C *Contact) UpdateLname(value interface{}) error {

	if getVariableType(value) != "string" {
		return fmt.Errorf("Please Enter a string value!")
	}
	if value == "" {
		return fmt.Errorf("L_name cannot be Empty!")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("Error in setting L_name string!")
	}
	C.L_name = conval
	return nil
}

func (C *Contact) UpdateisActive(value interface{}) error {
	if getVariableType(value) != "bool" {
		return fmt.Errorf("Please Enter a boolean value!")
	}
	conval, ok := value.(bool)
	if !ok {
		return fmt.Errorf("Error in setting isActive boolean!")
	}
	C.isActive = conval
	return nil
}

func (C *Contact) AddNewContact_Details(Type, NumberorEmail string) (*Contact_Details, error) {
	anewContact_Details, err := NewContact_Details(Type, NumberorEmail)
	if err != nil {
		return nil, err
	}
	C.Contact_Details = append(C.Contact_Details, anewContact_Details)
	return anewContact_Details, nil
}

func (C *Contact) UpdateContact_DetailsById(id int, param string, value interface{}) error {
	ContacWithMatchingId, err := C.GetContact_DetailsById(id)
	if err != nil {
		return err
	}
	err = ContacWithMatchingId.UpdateContact_Details(param, value)
	if err != nil {
		return err
	}
	return nil
}

func (C *Contact_Details) UpdateContact_Details(Type, value interface{}) error {
	if Type != "Number" && Type != "Email" {
		return fmt.Errorf("Type can either be a Number or Email!")
	}
	if Type == "Number" {
		C.Email = ""
		err := C.UpdateContact_DetailsNumber(value)
		if err != nil {
			return err
		}
		return nil
	} else {
		C.Number = ""
		err := C.UpdateConatct_DetailsEmail(value)
		if err != nil {
			return err
		}
		return nil
	}
}

func (C *Contact_Details) UpdateContact_DetailsNumber(value interface{}) error {
	if getVariableType(value) != "string" {
		return fmt.Errorf("Please Enter a string value!")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("Error in setting Number string!")
	}
	C.Number = conval
	return nil
}

func (C *Contact_Details) UpdateConatct_DetailsEmail(value interface{}) error {
	if getVariableType(value) != "string" {
		return fmt.Errorf("Please Enter a string value!")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("Error in setting Email string!")
	}
	C.Email = conval
	return nil
}

func main() {
	user1, err := NewUser("Vishav", "Pathania", true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user1)
	user2, err := user1.CreateNewUser("Yash", "Shah", false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user2)
	user1contact1, err := user1.AddNewContact("Aniket", "Pardeshi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user1contact1)
	fmt.Println(*user1)
	user1contact2, err := user1.AddNewContact("Someone", "New")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user1contact2)
	fmt.Println(*user1)
	err = user1.Contacts[1].DeleteContactById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user1.Contacts[1])
	err = user1.Contacts[0].UpdateContactById(1, "F_name", "Something")
	if err != nil {
		fmt.Println(err)
	}
	user1.Contacts[1].AddNewContact_Details("Email", "vishavpathania40@gmail.com")
	fmt.Println("Checking", *user1.Contacts[1].Contact_Details[0])
	err = user1.Contacts[1].DeleteContact_DetailsById(0, user1.Contacts[1].Contact_Details)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range user1.Contacts[1].Contact_Details {
		fmt.Println(*val)
		fmt.Println("Checking again----->")
	}

	user1.Contacts[1].Contact_Details[0].UpdateConatct_DetailsEmail("Email@gmail.com")
	fmt.Println(*user1.Contacts[1].Contact_Details[0])

	// fmt.Println(user1.Contacts)
	// for _, val := range user1.Contacts {
	// 	fmt.Println(*val)
	// 	fmt.Println("slice val----->")
	// }
	// sliceofusers, err := user1.GetAllSystemUser()
	// fmt.Println("-------------------------------")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, v := range sliceofusers {
	// 	fmt.Println(*v)
	// }

}
