package connection

import (
	"echo/database/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func DropUnusedColumns(dst interface{}) {
	DB := Connect()

	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(dst)
	fields := stmt.Schema.Fields
	columns, _ := DB.Debug().Migrator().ColumnTypes(dst)

	for i := range columns {
		found := false
		for j := range fields {
			if columns[i].Name() == fields[j].DBName {
				found = true
				break
			}
		}
		if !found {
			DB.Migrator().DropColumn(dst, columns[i].Name())
		}
	}
}

func HandleMigration() {
	fmt.Println("\033[34mRunning migrations...\033[0m")
	db := Connect()
	if db == nil {
		log.Fatal("Database connection failed, migration aborted.")
		return
	}
	err := db.AutoMigrate(&models.User{})
	DropUnusedColumns(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("\033[32mYeah! Migration finished ✅...\033[0m")
}
