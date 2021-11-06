package customer

type Customer struct {
	Name string
	Ssn  string
}

func (customer Customer) New(name string, ssn string) Customer {
	customer.Name = name
	customer.Ssn = ssn
	return customer
}
