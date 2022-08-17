package database

import "gorm.io/gorm"

func GetTableGames() *gorm.DB { return GetDB().Table("games") }
func GetTableNonce() *gorm.DB { return GetDB().Table("nonce") }
