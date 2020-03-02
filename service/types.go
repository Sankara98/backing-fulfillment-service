package service

type fullfillmentStatus struct {
	ProductID       string `json:"product_id"`
	ShipsWithin     int    `json:"ships_within"`
	QuantityInStock int    `json:"qty_in_stock"`
}
