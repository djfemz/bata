package test

import (
	"github.com/djfemz/simple_bank/app/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	configureTestEnvironment(t)
	db, err := repositories.Connect()
	assert.NotNil(t, db)
	assert.Nil(t, err)
}

func configureTestEnvironment(t *testing.T) {
	t.Setenv("DATABASE_PORT", "8000")
	t.Setenv("DATABASE_HOST", "localhost")
	t.Setenv("DATABASE_USERNAME", "postgres")
	t.Setenv("DATABASE_PASSWORD", "postgres")
	t.Setenv("DATABASE_NAME", "simple_bank_db")
	t.Setenv("DATABASE_PORT", "5432")
}
