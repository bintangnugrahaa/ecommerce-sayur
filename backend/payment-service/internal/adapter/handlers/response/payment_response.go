package response

type PaymentListResponse struct {
	ID            uint64  `json:"id"`
	OrderCode     string  `json:"order_code"`
	PaymentStatus string  `json:"payment_status"`
	PaymentMethod string  `json:"payment_method"`
	GrossAmount   float64 `json:"gross_amount"`
	ShippingType  string  `json:"shipping_type"`
}
