package forms


type OrderItem struct {
	Name          string         			`json:"name" bson:"name"  binding:"required"`
	Amount        int64								`json:"amount" bson:"amount" binding:"required"`
	Size          string         			`json:"size" bson:"size"`
	Ice           string         			`json:"ice" bson:"ice"`
	Sugar         string         		  `json:"sugar" bson:"sugar"`
	Padding       string           		`json:"padding" bson:"padding"`
}

type CreateOrderCommand struct {
	UserId      string  `json:"userId"  bson:"userId" binding:"required"`
	OrderList   []OrderItem  `json:"orderList"  bson:"OrderList"  binding:"required"`
}

type ListOrderCommand struct {
	Time  int64    `json:"time"  bson:"time" binding:"required"`
}


type FindLastOrderCommand struct {
	UserId      string  `json:"userId"  bson:"userId" binding:"required"`
}

type UpdateProductCommand struct {
	ProductId      string  `json:"productId"  bson:"productId" binding:"required"`
	Available      *bool   `json:"available" validate:"exists"`
}
