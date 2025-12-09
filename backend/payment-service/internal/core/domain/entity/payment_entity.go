package entity

type PaymentEntity struct {
	ID               uint
	OrderID          uint
	UserID           uint
	PaymentMethod    string
	PaymentStatus    string
	PaymentGatewayID string
	GrossAmount      float64
	PaymentURL       string
	PaymentLogs      []PaymentLogEntity
	Remarks          string
	CustomerName     string
	CustomerEmail    string
}
