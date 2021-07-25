package product

type Product struct {
	ProductId      *int   `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
	PricePerUnit   string `json:"pricePerUnit"`
}
