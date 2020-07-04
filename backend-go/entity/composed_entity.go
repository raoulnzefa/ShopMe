package entity

type CartProduct struct {
	Cart
	Product
}

type PromoteSection struct {
	Title    string    `json:"title"`
	Category int       `json:"category"`
	Products []Product `json:"products"`
}

type SimpleOrderWithProducts struct {
	ID       int              `json:"id"`
	Date     string           `json:"date"`
	Sum      float32          `json:"sum"`
	Status   OrderStatus      `json:"status"`
	Products []ProductOfOrder `json:"products"`
}

type ProductOfOrder struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Num   int     `json:"num"`
	Image string  `json:"image"`
	Price float32 `json:"price"`
}
