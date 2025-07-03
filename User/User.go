package user

import (
	"fmt"
	contact "package_contactapp/Contact"
	contact_detail "package_contactapp/Contact_Details"
	utils "package_contactapp/Utils"
)

var userMap = make(map[int]*User)
var userMapid = -1

type User struct {
	User_ID  int
	F_name   string
	L_name   string
	isAdmin  bool
	isActive bool
	Contacts []*contact.Contact
}

func newUser(F_name, L_name string, isAdmin bool) (*User, error) {
	if F_name == "" {
		return nil, fmt.Errorf("first Name Cannot be Empty")
	}
	if L_name == "" {
		return nil, fmt.Errorf("last Name Cannot be Empty")
	}
	userMapid++
	u := &User{
		User_ID:  userMapid,
		F_name:   F_name,
		L_name:   L_name,
		isAdmin:  isAdmin,
		isActive: true,
		Contacts: []*contact.Contact{},
	}
	userMap[userMapid] = u
	return u, nil
}

func (U *User) CreateNewStaff(F_name, L_name string) (*User, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("you Don't have admin previlages to create a User")
	}
	anewUser, err := newUser(F_name, L_name, false)
	if err != nil {
		return nil, err
	}
	return anewUser, nil
}

func CreateNewAdmin(F_name, L_name string) (*User, error) {
	// if !U.isAdmin {
	// 	return nil, fmt.Errorf("you Don't have admin previlages to create a User")
	// }
	anewUser, err := newUser(F_name, L_name, true)
	if err != nil {
		return nil, err
	}
	return anewUser, nil
}

// returning shallow copy of Users
func (U *User) GetAllSystemUser() ([]User, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("you don't have admin previlages to Read all Users")
	}
	AllUsersSlice := []User{}
	for _, val := range userMap {
		AllUsersSlice = append(AllUsersSlice, *val)
	}
	return AllUsersSlice, nil
}

func (U *User) GetAllUserContacts(Userid int) ([]contact.Contact, error) {
	if U.isAdmin {
		return nil, fmt.Errorf("admin cannot read contacts")
	}
	if !U.isActive {
		return nil, fmt.Errorf("inactive users cannot read contacts")
	}
	copyOfUserContacts := []contact.Contact{}
	for _, userContact := range U.Contacts {
		copyOfUserContacts = append(copyOfUserContacts, *userContact)
	}
	return copyOfUserContacts, nil
}

func (U *User) GetUserById(id int) (*User, error) {
	if !U.isAdmin {
		return nil, fmt.Errorf("you don't have admin previlages to Read User")
	}
	if !U.isActive {
		return nil, fmt.Errorf("inactive admin cannot read users")
	}
	UserbyId, ok := userMap[id]
	if !ok {
		return nil, fmt.Errorf("user with id: %d not found", id)
	}
	return UserbyId, nil
}

func (U *User) DeleteUserById(id int) error {
	if !U.isAdmin {
		return fmt.Errorf("you Don't have admin previlages to Delete User")
	}
	if !U.isActive {
		return fmt.Errorf("inactive Users Cannot Delete Users")
	}
	UserWithMatchingId, err := U.GetUserById(id)
	if err != nil {
		return err
	}
	UserWithMatchingId.isActive = false
	return nil
}

func (U *User) ValidateContactId(id int) bool {
	if id < 0 || id > len(U.Contacts) {
		return false
	}
	return true
}

func (U *User) GetContactById(id int) (*contact.Contact, error) {
	if U.isAdmin {
		return nil, fmt.Errorf("admin cannot get contact by id")
	}
	if !U.isActive {
		return nil, fmt.Errorf("inactive users cannot get contact by id")
	}
	checkid := U.ValidateContactId(id)
	if !checkid {
		return nil, fmt.Errorf("please provide a valid contact id")
	}
	for _, contactsval := range U.Contacts {
		if contactsval.CheckIfContactActivebyId() == id {
			return contactsval, nil
		}
	}
	return nil, fmt.Errorf("didn't found active contact with given id:%d", id)
}

func (U *User) GetContact_DetailsById(ContactId, Contact_DetailsId int) (*contact_detail.Contact_Details, error) {
	if U.isAdmin {
		return nil, fmt.Errorf("admin not allowed to get Contacts from Id")
	}
	if !U.isActive {
		return nil, fmt.Errorf("inactive users are not allowed to get contact_details from Id")
	}

	TargetContact, err := U.GetContactById(ContactId)
	if err != nil {
		return nil, err
	}

	if !TargetContact.ValidateContact_DetailsId(Contact_DetailsId) {
		return nil, fmt.Errorf("please give a valid contact_details id")
	}
	resultContact_Details, err := TargetContact.GetContact_DetailsById(Contact_DetailsId)
	if err != nil {
		return nil, err
	}
	return resultContact_Details, nil
}

func (U *User) DeleteContactById(id int) error {
	if U.isAdmin {
		return fmt.Errorf("admins are not allowed to delete Contacts")
	}
	if !U.isActive {
		return fmt.Errorf("inactive Users are not allowed to delete Contacts")
	}
	checkid := U.ValidateContactId(id)
	if !checkid {
		return fmt.Errorf("please provide valid contact id")
	}
	ContactWithMatchingId, err := U.GetContactById(id)
	if err != nil {
		return err
	}
	err = ContactWithMatchingId.DeleteContactById()
	if err != nil {
		return err
	}
	return nil
}

func (U *User) DeleteContact_DetailsById(Contactid, Contact_Details_ID int) error {
	if U.isAdmin {
		return fmt.Errorf("admins are not allowed to deleted contact_details")
	}
	if !U.isActive {
		return fmt.Errorf("inactive users are not allowed to delete contact_details")
	}
	ContacthavingDetailsSlice, err := U.GetContactById(Contactid)
	if err != nil {
		return err
	}
	err = ContacthavingDetailsSlice.DeleteContact_DetailsById(Contact_Details_ID)
	if err != nil {
		return err
	}

	return nil
}

func (U *User) UpdateUser(UserId int, param string, value interface{}) error {
	if !U.isAdmin {
		return fmt.Errorf("you Don't have admin previlages to update users")
	}
	if !U.isActive {
		return fmt.Errorf("inactive Users Cannot Update users")
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
	default:
		return fmt.Errorf("no matching params founnd to Update")
	}
}

func (U *User) UpdateContactById(Contactid int, param string, value interface{}) error {
	if U.isAdmin {
		return fmt.Errorf("admin's are not allowed to update contact")
	}
	if !U.isActive {
		return fmt.Errorf("inactive users are not allowed to update contact")
	}
	ContacWithMatchingId, err := U.GetContactById(Contactid)
	if err != nil {
		return err
	}
	err = ContacWithMatchingId.UpdateContact(param, value)
	if err != nil {
		return err
	}
	return nil
}

func (U *User) UpdateFname(value interface{}) error {
	if !U.isActive {
		return fmt.Errorf("cannot update inactive users")
	}
	if utils.GetVariableType(value) != "string" {
		return fmt.Errorf("please Enter a string value")
	}
	if value == "" {
		return fmt.Errorf("f_name cannot be Empty")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("error in setting F_name string")
	}
	U.F_name = conval
	return nil
}

func (U *User) UpdateLname(value interface{}) error {
	if !U.isActive {
		return fmt.Errorf("cannot update inactive users")
	}
	if utils.GetVariableType(value) != "string" {
		return fmt.Errorf("please Enter a string value")
	}
	if value == "" {
		return fmt.Errorf("l_name cannot be Empty")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("error in setting L_name string")
	}
	U.L_name = conval
	return nil
}

func (U *User) UpdateisAdmin(value interface{}) error {
	if !U.isActive {
		return fmt.Errorf("cannot update inactive users")
	}
	if utils.GetVariableType(value) != "bool" {
		return fmt.Errorf("please Enter a boolean value")
	}
	conval, ok := value.(bool)
	if !ok {
		return fmt.Errorf("error in setting isAdmin boolean")
	}
	U.isAdmin = conval
	return nil
}

func (U *User) AddNewContact(F_name, L_name string) (*contact.Contact, error) {
	if !U.isActive {
		return nil, fmt.Errorf("inactive User Cannot Add New Contact")
	}
	if U.isAdmin {
		return nil, fmt.Errorf("admin Cannot Add New Contacts")
	}
	newId := len(U.Contacts) + 1
	anewContact, err := contact.NewContact(newId, F_name, L_name)
	if err != nil {
		return nil, err
	}
	U.Contacts = append(U.Contacts, anewContact)
	return anewContact, nil
}

func (U *User) AddNewContact_DetailsByContactId(ContactId int, Type, NumberorEmail string) (*contact_detail.Contact_Details, error) {
	if U.isAdmin {
		return nil, fmt.Errorf("admin is not allowed to create contact_details")
	}
	if !U.isActive {
		return nil, fmt.Errorf("inactive Users are not allowed to create contact_details")
	}
	ContactToAddContactDetails, err := U.GetContactById(ContactId)
	if err != nil {
		return nil, err
	}
	anewContact_Details, err := ContactToAddContactDetails.AddNewContact_Details(Type, NumberorEmail)
	if err != nil {
		return nil, err
	}
	return anewContact_Details, nil
}

func (U *User) UpdateContact_DetailsById(ContactId, ContactDetailid int, param string, value interface{}) error {
	if U.isAdmin {
		return fmt.Errorf("admin is not allowed to create contact_details")
	}
	if !U.isActive {
		return fmt.Errorf("inactive Users are not allowed to create contact_details")
	}
	ContactDetailsWithMatchingId, err := U.GetContact_DetailsById(ContactId, ContactDetailid)
	if err != nil {
		return err
	}
	err = ContactDetailsWithMatchingId.UpdateContact_Details(param, value)
	if err != nil {
		return err
	}
	return nil
}
