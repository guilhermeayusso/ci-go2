package config

import (
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnectDatabase(t *testing.T) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao conectar com o BD: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Erro ao obter conexão SQL: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Erro ao pingar o BD: %v", err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			t.Fatalf("Erro ao fechar conexão SQL: %v", err)
		}
	}()
}
