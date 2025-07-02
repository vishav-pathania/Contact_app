package contact_detail

import (
	"fmt"
	utils "package_contactapp/Utils"
)

type Contact_Details struct {
	Contact_Details_ID int
	Type               string
	Number             string
	Email              string
}

func NewContact_Details(Id int, Type, NumberorEmail string) (*Contact_Details, error) {
	if Type == "" || (Type != "Number" && Type != "Email") {
		return nil, fmt.Errorf("invalid type")
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

func (C *Contact_Details) UpdateContact_Details(Type, value interface{}) error {
	if Type != "Number" && Type != "Email" {
		return fmt.Errorf("type can either be a number or email")
	}
	if Type == "Number" {
		err := C.UpdateContact_DetailsNumber(value)
		if err != nil {
			return err
		}
		C.Email = ""
		return nil
	} else {
		err := C.UpdateConatct_DetailsEmail(value)
		if err != nil {
			return err
		}
		C.Number = ""
		return nil
	}
}

func (C *Contact_Details) UpdateContact_DetailsNumber(value interface{}) error {
	if utils.GetVariableType(value) != "string" {
		return fmt.Errorf("please enter a string value")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("error in setting number string")
	}
	C.Number = conval
	return nil
}

func (C *Contact_Details) UpdateConatct_DetailsEmail(value interface{}) error {
	if utils.GetVariableType(value) != "string" {
		return fmt.Errorf("please Enter a string value")
	}
	conval, ok := value.(string)
	if !ok {
		return fmt.Errorf("error in setting email string")
	}
	C.Email = conval
	return nil
}
