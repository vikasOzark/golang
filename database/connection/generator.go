package connection

import (
	"echo/database/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func Generate() {
	// Initialize the generator
	g := gen.NewGenerator(gen.Config{
		OutPath: "/home/vikas/temp/go-lang/database", // Specify the output directory
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	dsn := os.Getenv("DATABASE")
	if dsn == "" {
		log.Fatal("DATABASE environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g.UseDB(db)

	// Generate basic type-safe DAO API for your models
	g.ApplyBasic(
		// List your models here
		models.User{},
	)

	g.Execute()
}
