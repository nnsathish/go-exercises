package models

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
  // TODO: need to read the values from env / secrets
  dsn := "host=localhost user=sathish.n password= dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Kolkata"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  // TODO: check if connection pooling is mandatory!

  if err != nil {
    panic("Failed to connect to database!")
  }

  db.AutoMigrate(&Product{})

  DB = db
}
