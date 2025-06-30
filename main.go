package main

import (
	"fmt"
	"reflect"
)

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

func NewContact_Details(Id int, Type, NumberorEmail string) (*Contact_Details, error) {
	if Type == "" || (Type != "Number" && Type != "Email") {
		return nil, fmt.Errorf("Invalid Type!")
	}

	cd := &Contact_Details{
		Contact_Details_ID: Id,
		Type:               Type,
	}

	if Type == "Email" {
		cd.Email = NumberorEmail
	} else {
		cd.Number = NumberorEmail
	}

	return cd, nil

}

func NewContact(Id int, F_name string, L_name string) (*Contact, error) {
	if F_name == "" {
		return nil, fmt.Errorf("First Name Cannot be Empty!")
	}
	if L_name == "" {
		return nil, fmt.Errorf("Last Name Cannot be Empty!")
	}

	c := &Contact{
		Contact_ID:      Id,
		F_name:          F_name,
		L_name:          L_name,
		isActive:        true,
		Contact_Details: []*Contact_Details{},
	}
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

func (U *User) CreateNewStaff(F_name, L_name string) (*User, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("You Don't have admin previlages to create a User!")
	}
	anewUser, err := NewUser(F_name, L_name, false)
	if err != nil {
		return nil, err
	}
	return anewUser, nil
}

func (U *User) CreateNewAdmin(F_name, L_name string) (*User, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("You Don't have admin previlages to create a User!")
	}
	anewUser, err := NewUser(F_name, L_name, true)
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
	if !U.isActive {
		return fmt.Errorf("Inactive Users Cannot Delete Users!")
	}
	UserWithMatchingId, err := U.GetUserById(id)
	if err != nil {
		return err
	}
	UserWithMatchingId.isActive = false
	return nil
}

func (U *User) GetContactById(id int) (*Contact, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("Admin not allowed to get Contacts from Id!")
	}
	if !U.isActive {
		return nil, fmt.Errorf("Inactive Users are not allowed to get Contacts from Id!")
	}
	if !U.ValidateContactId(id) {
		return nil, fmt.Errorf("Please give a valid ID")
	}
	for _, contactsval := range U.Contacts {
		if contactsval.CheckIfContactActivebyId() == id {
			return contactsval, nil
		}
	}
	return nil, fmt.Errorf("Didn't Found Contact with given id:%d", id)
}

func (U *User) ValidateContactId(id int) bool {
	if id < 0 || id > len(U.Contacts) {
		return false
	}
	return true
}

func (C *Contact) ValidateContact_DetailsId(id int) bool {
	if id < 0 || id > len(C.Contact_Details) {
		return false
	}
	return true
}

func (C *Contact) CheckIfContactActivebyId() int {
	if C.isActive {
		return C.Contact_ID
	}
	return -1
}

func (U *User) GetContact_DetailsById(ContactId, Contact_DetailsId int) (*Contact_Details, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("Admin not allowed to get Contacts from Id!")
	}
	if !U.isActive {
		return nil, fmt.Errorf("Inactive Users are not allowed to get Contacts from Id!")
	}

	TargetContact, err := U.GetContactById(ContactId)
	if err != nil {
		return nil, err
	}

	if !TargetContact.ValidateContact_DetailsId(Contact_DetailsId) {
		return nil, fmt.Errorf("Please give a valid ID")
	}

	for _, ConDetails := range TargetContact.Contact_Details {
		if ConDetails.Contact_Details_ID == Contact_DetailsId {
			return ConDetails, nil
		}
	}
	return nil, fmt.Errorf("Didn't Found Contact with given id:%d", Contact_DetailsId)
}

func (U *User) DeleteContactById(id int) error {
	if U.isAdmin {
		return fmt.Errorf("Admins are not allowed to delete Contacts!")
	}
	if !U.isActive {
		return fmt.Errorf("Inactive Users are not allowed to delete Contacts")
	}
	ContactWithMatchingId, err := U.GetContactById(id)
	if err != nil {
		return err
	}
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

func (U *User) DeleteContact_DetailsById(Contactid, Contact_Details_ID int) error {
	Contact_DetailsObjectToDelete, err := U.GetContact_DetailsById(Contactid, Contact_Details_ID)
	if err != nil {
		return err
	}
	ContactSliceToPerformDelete, err := U.GetContactById(Contactid)
	if err != nil {
		return err
	}
	NewContact_DetailsSliceToset := []*Contact_Details{}
	for _, Contact_Detailsval := range ContactSliceToPerformDelete.Contact_Details {
		if Contact_Detailsval != Contact_DetailsObjectToDelete {
			NewContact_DetailsSliceToset = append(NewContact_DetailsSliceToset, Contact_Detailsval)
		}
	}
	ContactSliceToPerformDelete.Contact_Details = NewContact_DetailsSliceToset
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

func (U *User) UpdateUser(UserId int, param string, value interface{}) error {
	if !U.isAdmin {
		return fmt.Errorf("You Don't have admin previlages to update users!")
	}
	if !U.isActive {
		return fmt.Errorf("Inactive Users Cannot Update users!")
	}
	UserToBeUpdated, err := U.GetUserById(UserId)
	if err != nil {
		return err
	}

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
	// case "isActive":
	// 	err := UserToBeUpdated.UpdateisActive(value)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	default:
		return fmt.Errorf("No matching params founnd to Update!")
	}
}

func (U *User) UpdateContactById(id int, param string, value interface{}) error {
	if U.isAdmin {
		return fmt.Errorf("Admin's are not allowed to update contact!")
	}
	if !U.isActive {
		return fmt.Errorf("Inactive Users are not allowed to update contact!")
	}
	ContacWithMatchingId, err := U.GetContactById(id)
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
	// case "isActive":
	// 	err := C.UpdateisActive(value)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
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

// func (U *User) UpdateisActive(value interface{}) error {
// 	if !U.isAdmin {
// 		return fmt.Errorf("You Don't have admin previlages to Update User Active Status!")
// 	}
// 	if getVariableType(value) != "bool" {
// 		return fmt.Errorf("Please Enter a boolean value!")
// 	}
// 	conval, ok := value.(bool)
// 	if !ok {
// 		return fmt.Errorf("Error in setting isActive boolean!")
// 	}
// 	U.isActive = conval
// 	return nil
// }

func (U *User) AddNewContact(F_name, L_name string) (*Contact, error) {
	//active or not admin
	if !U.isActive {
		return nil, fmt.Errorf("Inactive User Cannot Add New Contact!")
	}
	if U.isAdmin {
		return nil, fmt.Errorf("Admin Cannot Add New Contacts!")
	}
	newId := len(U.Contacts)
	anewContact, err := NewContact(newId, F_name, L_name)
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

// func (C *Contact) UpdateisActive(value interface{}) error {
// 	if getVariableType(value) != "bool" {
// 		return fmt.Errorf("Please Enter a boolean value!")
// 	}
// 	conval, ok := value.(bool)
// 	if !ok {
// 		return fmt.Errorf("Error in setting isActive boolean!")
// 	}
// 	C.isActive = conval
// 	return nil
// }

func (U *User) AddNewContact_DetailsByContacId(ContactId int, Type, NumberorEmail string) (*Contact_Details, error) {
	if U.isAdmin {
		return nil, fmt.Errorf("Admin is not allowed to create contact_details!")
	}
	if !U.isActive {
		return nil, fmt.Errorf("Inactive Users are not allowed to create contact_details!")
	}
	ContactToAddContactDetails, err := U.GetContactById(ContactId)
	if err != nil {
		return nil, err
	}
	NewId := len(ContactToAddContactDetails.Contact_Details)
	anewContact_Details, err := NewContact_Details(NewId, Type, NumberorEmail)
	if err != nil {
		return nil, err
	}
	ContactToAddContactDetails.Contact_Details = append(ContactToAddContactDetails.Contact_Details, anewContact_Details)
	return anewContact_Details, nil
}

func (U *User) UpdateContact_DetailsById(ContactId, ContactDetailid int, param string, value interface{}) error {
	if U.isAdmin {
		return fmt.Errorf("Admin is not allowed to create contact_details!")
	}
	if !U.isActive {
		return fmt.Errorf("Inactive Users are not allowed to create contact_details!")
	}
	ContacWithMatchingId, err := U.GetContact_DetailsById(ContactId, ContactDetailid)
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
	user2, err := user1.CreateNewStaff("Yash", "Shah")
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
	err = user1.DeleteContactById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user1.Contacts[1])
	err = user1.UpdateContactById(1, "F_name", "Something")
	if err != nil {
		fmt.Println(err)
	}
	// user1.Contacts[1].AddNewContact_Details("Email", "vishavpathania40@gmail.com")
	// fmt.Println("Checking", *user1.Contacts[1].Contact_Details[0])
	// err = user1.Contacts[1].DeleteContact_DetailsById(0, user1.Contacts[1].Contact_Details)
	// if err != nil {
	// 	fmt.Println(err)
	// }
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
