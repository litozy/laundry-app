package entity

type Transaction struct {
	Transaction_Id, Customer_Id, Service_Id int
	Transaction_In, Transaction_Done        string
	Received_By                             string
	Quantity                                int
	Unit                                    string
	Price, Total_Price                      float64
}

type Customer struct {
	ID    string
	Name  string
	Phone string
}

type Service struct {
	ID    int
	Name  string
	Price float64
}