package storeApi

type Product struct {
	ID          int     `json:"id,omitempty" db:"id"`
	OrderNum    int     `json:"order_num" binding:"required" db:"order_num"`
	Name        string  `json:"name" binding:"required" db:"name"`
	Price       float64 `json:"price" binding:"required" db:"price"`
	Description string  `json:"description" binding:"required" db:"description"`
	Image       []byte  `json:"image" binding:"required" db:"image"`
}
