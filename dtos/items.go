package dtos

type VegetableItem struct {
	ProductID   string `json:"productId"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	Price       string `json:"price"`
}

type GrainItem struct {
	ItemID   string `json:"itemId"`
	ItemName string `json:"itemName"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}
