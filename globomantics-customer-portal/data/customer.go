package data

import "errors"

type Customer struct {
	CustomerID   string `json:"customerid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	SubType      string `json:"subtype"`
	Active       string `json:"active"`
	CreationDate string `json:"creation"`
}

type Customers []*Customer

var (
	customerList = Customers{}
	/* may not need this
	EmptyCustomer = Customer{
		CustomerID:   "",
		FirstName:    "",
		LastName:     "",
		Address:      "",
		Phone:        "",
		Email:        "",
		SubType:      "",
		Active:       "",
		CreationDate: "",
	}
	*/
)

func GetCustomers() []*Customer {
	return customerList
}

func GetCustomerByIdOrNil(id string) *Customer {
	for _, c := range customerList {
		if c.CustomerID == id {
			return c
		}
	}
	return nil
}

func Contains(id string) bool {
	for _, c := range customerList {
		if c.CustomerID == id {
			return true
		}
	}
	return false
}

func AddCustomer(c *Customer) (bool, error) {
	if !Contains(c.CustomerID) {
		customerList = append(customerList, c)
		return true, nil
	} else {
		return false, errors.New("customer already exists")
	}
}

func DeleteCustomer(c *Customer) {
	matchingIndex := -1
	for index, item := range customerList {
		if item.CustomerID == c.CustomerID {
			matchingIndex = index
			break
		}
	}
	if matchingIndex == -1 {
		return
	}

	// maintaining order
	customerList = append(customerList[:matchingIndex], customerList[matchingIndex+1:]...)
	/* or not maintaining order
	customerList[matchingIndex] = customerList[len(customerList) - 1]
	customerList = customerList[:len(customerList) - 1]
	*/
}
