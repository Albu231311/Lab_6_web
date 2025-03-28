package database

import (
	"laliga/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Conexión a la base de datos y migraciones
func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("laliga_matches.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ No se pudo conectar a la base de datos:", err)
	}

	err = DB.AutoMigrate(&models.Match{})
	if err != nil {
		log.Fatal("❌ Error al migrar la base de datos:", err)
	}

	log.Println("✅ Base de datos conectada y migrada correctamente")
}
