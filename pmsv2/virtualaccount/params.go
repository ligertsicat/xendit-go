package virtualaccount

import (
	"time"

	"github.com/xendit/xendit-go/pmsv2/constant"
)

type ChannelCode string

const (
	BCA               ChannelCode = "BCA"
	BJB               ChannelCode = "BJB"
	BNI               ChannelCode = "BNI"
	BRI               ChannelCode = "BRI"
	MANDIRI           ChannelCode = "MANDIRI"
	PERMATA           ChannelCode = "PERMATA"
	BSI               ChannelCode = "BSI"
	CIMB              ChannelCode = "CIMB"
	SAHABAT_SAMPOERNA ChannelCode = "SAHABAT_SAMPOERNA"
)

type ChannelProperties struct {
	CustomerName         string     `json:"customer_name"`
	VirtualAccountNumber *string    `json:"virtual_account_number"`
	ExpiresAt            *time.Time `json:"expires_at"`
	SuggestedAmount      float64    `json:"suggested_amount,omitempty"`
}

type CreateMethod struct {
	Amount            float64               `json:"amount"`
	MinimumAmount     float64               `json:"min_amount,omitempty"`
	MaximumAmount     float64               `json:"max_amount,omitempty"`
	Currency          constant.CurrencyEnum `json:"currency"`
	ChannelCode       ChannelCode           `json:"channel_code"`
	ChannelProperties ChannelProperties     `json:"channel_properties"`
}

type MutableMethod struct {
	Amount            int64                    `json:"amount"`
	MinimumAmount     int64                    `json:"min_amount"`
	MaximumAmount     int64                    `json:"max_amount"`
	ChannelProperties MutableChannelProperties `json:"channel_properties"`
}

type MutableChannelProperties struct {
	ExpiresAt       *time.Time `json:"expires_at"`
	SuggestedAmount int64      `json:"suggested_amount"`
}
