package domain

type Warehouse struct {
	ID      int
	Name    string
	Address string
}

type WarehouseReport struct {
	Warehouse
	ProductCount int
}