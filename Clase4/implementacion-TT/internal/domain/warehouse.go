package domain

type Warehouse struct {
	Id        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	Adress    string `json:"adress" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required"`
}

//type WarehouseReport struct {
//	Warehouse
//	ProductCount int
//}

type WarehouseReport struct {
	Name         string `json:"warehouse_name"`
	ProductCount int    `json:"product_count"`
}
