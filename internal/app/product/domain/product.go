package domain

import (
	"time"
)

type Product struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	CategoryID           string `json:"category,omitempty"`
	Status               string `json:"status"`
	BasePriceNumerator   int64
	BasePriceDenominator int64     `json:"BasePriceDenominator"`
	DiscountPercent      string    `json:"discount_percent"`
	DiscountStartDate    time.Time `json:"discount_start_date"`
	DiscountEndDate      time.Time `json:"discount_end_date"`
	Version              int64     `json:"version"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
