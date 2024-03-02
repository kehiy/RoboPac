package database

import "gorm.io/gorm"

type User struct {
	gorm.Model

	DiscordID      string `gorm:"unique"          json:"discord_id"`
	DepositAddress string `json:"deposit_address"`
}
