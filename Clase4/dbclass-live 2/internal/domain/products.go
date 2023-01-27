package domain

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
	WarehouseID int     `json:"warehouse_id"`
}

type ProductWithWarehouse struct {
	Product
	WarehouseName    string `json:"warehouse_name"`
	WarehouseAddress string `json:"warehouse_address"`
}