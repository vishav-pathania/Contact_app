package contact

import (
	"fmt"
	contact_detail "package_contactapp/Contact_Details"
	utils "package_contactapp/Utils"
)

type Contact struct {
	Contact_ID      int
	F_name          string
	L_name          string
	isActive        bool
	Contact_Details []*contact_detail.Contact_Details
}

func NewContact(Id int, F_name string, L_name string) (*Contact, error) {
	if F_name == "" {
		return nil, fmt.Errorf("first Name Cannot be Empty")
	}
	if L_name == "" {
		return nil, fmt.Errorf("last Name Cannot be Empty")
	}

	c := &Contact{
		Contact_ID:      Id,
		F_name:          F_name,
		L_name:          L_name,
		isActive:        true,
		Contact_Details: []*contact_detail.Contact_Details{},
	}
	return c, nil

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

func (C *Contact) DeleteContact_DetailsById(contact_detailsid int) error {
	if !C.isActive {
		return fmt.Errorf("inactive users cannot perform delete operation")
	}
	checkid := C.ValidateContact_DetailsId(contact_detailsid)
	if !checkid {
		return fmt.Errorf("please provide a valid contact_details id")
	}

	slicetoremovecontactdetails := []*contact_detail.Contact_Details{}
	for _, contactdetailsval := range C.Contact_Details {
		if contactdetailsval.Contact_Details_ID != contact_detailsid {
			slicetoremovecontactdetails = append(slicetoremovecontactdetails, contactdetailsval)
		}
	}
	C.Contact_Details = slicetoremovecontactdetails
	return nil
}

func (C *Contact) UpdateContact(param string, value interface{}) error {
	if !C.isActive {
		return fmt.Errorf("inactive contact cannot update contacts")
	}
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
	default:
		return fmt.Errorf("no matching params founnd to Update")

	}
}

func (C *Contact) UpdateFname(value interface{}) error {
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
	C.F_name = conval
	return nil
}

func (C *Contact) UpdateLname(value interface{}) error {

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
	C.L_name = conval
	return nil
}

func (C *Contact) DeleteContactById() error {
	if !C.isActive {
		return fmt.Errorf("contact already inactive")
	}
	C.isActive = false
	return nil
}

func (C *Contact) GetContact_DetailsById(Contact_DetailsId int) (*contact_detail.Contact_Details, error) {
	if !C.isActive {
		return nil, fmt.Errorf("inactive users cannot get contact_details by id")
	}
	checkid := C.ValidateContact_DetailsId(Contact_DetailsId)
	if !checkid {
		return nil, fmt.Errorf("please provide valid contact_details id")
	}
	for _, contact_detailsvalue := range C.Contact_Details {
		if contact_detailsvalue.Contact_Details_ID == Contact_DetailsId {
			return contact_detailsvalue, nil
		}
	}
	return nil, fmt.Errorf("didn't found contact_details with id: %d", Contact_DetailsId)
}

func (C *Contact) AddNewContact_Details(Type, NumberorEmail string) (*contact_detail.Contact_Details, error) {
	if !C.isActive {
		return nil, fmt.Errorf("inactive contacts cannot add new contact details")
	}
	// newId := len(C.Contact_Details) + 1
	newId := C.Contact_Details[len(C.Contact_Details)-1].Contact_Details_ID + 1
	newContactDetails, err := contact_detail.NewContact_Details(newId, Type, NumberorEmail)
	if err != nil {
		return nil, err
	}
	C.Contact_Details = append(C.Contact_Details, newContactDetails)

	return newContactDetails, nil
}

// returning shallow copy of contact_details
func (C *Contact) GetallContact_Details() ([]contact_detail.Contact_Details, error) {
	if !C.isActive {
		return nil, fmt.Errorf("inactive users cannot read contact_details")
	}
	var contact_detailscopy = []contact_detail.Contact_Details{}
	for _, contact_detailval := range C.Contact_Details {
		contact_detailscopy = append(contact_detailscopy, *contact_detailval)
	}
	return contact_detailscopy, nil
}
