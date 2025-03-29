package database

import (
	"lab6/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect inicializa la conexión a la base de datos y realiza la migración automática del modelo Match
func Connect() {
	var err error

	// Abrir la conexión con una base de datos SQLite llamada matches.db
	DB, err = gorm.Open(sqlite.Open("matches.db"), &gorm.Config{})
	if err != nil {
		// Si hay un error al conectar, detener la aplicación
		log.Fatal("Unable to connect the DB:", err)
	}

	// Ejecutar la migración automática del modelo Match (crear tabla si no existe)
	err = DB.AutoMigrate(&models.Match{})
	if err != nil {
		log.Fatal("Unable to migrate the model:", err)
	}

	// Mensaje informativo si todo funcionó correctamente
	log.Println("Connection successfully")
}
