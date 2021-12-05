package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UsersID      uint      `json:"users_id" form:"users_id"`
	RoomsID      uint      `json:"rooms_id" form:"rooms_id"`
	Check_In     time.Time `gorm:"type:datetime;not null" json:"check_in" form:"check_in"`
	Jumlah_Malam int       `json:"jumlah_malam" form:"jumlah_malam"`
	Check_Out    time.Time `gorm:"type:datetime" json:"check_out" form:"check_out"`
	Total_Harga  int       `json:"total_harga" form:"total_harga"`
	Review       Review    `gorm:"foreignKey:ReservationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Struct untuk body create reservation
type ReservationBody struct {
	RoomsID   uint   `json:"rooms_id" form:"rooms_id"`
	Check_In  string `gorm:"type:datetime;not null" json:"check_in" form:"check_in"`
	Check_Out string `gorm:"type:datetime" json:"check_out" form:"check_out"`
	Phone     string `json:"phone" form:"phone"`
}

// Struct response data untuk get reservation
type GetReservation struct {
	RoomsID     uint
	Check_In    string
	Check_Out   string
	Total_Harga int
}

// Struct Body
type Metadata struct {
	BranchArea string `json:"branch_area" validate:"required"`
	BranchCity string `json:"branch_city" validate:"required"`
}
type ChannelProperties struct {
	MobileNumber string `json:"mobile_number" validate:"required"`
}
type RequestBodyStruct struct {
	ReferenceID       string            `json:"reference_id" validate:"required"`
	Currency          string            `json:"currency" validate:"required"`
	Amount            float64           `json:"amount" validate:"required"`
	CheckoutMethod    string            `json:"checkout_method" validate:"required"`
	ChannelCode       string            `json:"channel_code" validate:"required"`
	ChannelProperties ChannelProperties `json:"channel_properties" validate:"required"`
	Metadata          Metadata          `json:"metadata" validate:"required"`
}

type BasketItem struct {
	ReferenceID string                 `json:"reference_id" validate:"required"`
	Name        string                 `json:"name" validate:"required"`
	Category    string                 `json:"category" validate:"required"`
	Currency    string                 `json:"currency" validate:"required"`
	Price       float64                `json:"price" validate:"required"`
	Quantity    int                    `json:"quantity" validate:"required"`
	Type        string                 `json:"type" validate:"required"`
	Url         string                 `json:"url,omitempty"`
	Description string                 `json:"description,omitempty"`
	SubCategory string                 `json:"sub_category,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}
type ResponsePayment struct {
	ID                 string            `json:"id"`
	BusinessID         string            `json:"business_id"`
	ReferenceID        string            `json:"reference_id"`
	Status             string            `json:"status"`
	Currency           string            `json:"currency"`
	ChargeAmount       float64           `json:"charge_amount"`
	CaptureAmount      float64           `json:"capture_amount"`
	CheckoutMethod     string            `json:"checkout_method"`
	ChannelCode        string            `json:"channel_code"`
	ChannelProperties  ChannelProperties `json:"channel_properties"`
	Actions            map[string]string `json:"actions"`
	IsRedirectRequired bool              `json:"is_redirect_required"`
	CallbackURL        string            `json:"callback_url"`
	Created            string            `json:"created"`
	Updated            string            `json:"updated"`
	VoidedAt           string            `json:"voided_at,omitempty"`
	CaptureNow         bool              `json:"capture_now"`
	CustomerID         string            `json:"customer_id,omitempty"`
	PaymentMethodID    string            `json:"payment_method_id,omitempty"`
	FailureCode        string            `json:"failure_code,omitempty"`
	Basket             []BasketItem      `json:"basket,omitempty"`
	Metadata           Metadata          `json:"metadata,omitempty"`
}
