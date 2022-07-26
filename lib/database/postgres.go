package database

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hackathon/lib/datamodel"
	"hackathon/lib/utils"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	err := utils.LoadEnv()
	if err != nil {
		return nil, errors.Wrap(err, "NewPostgresDB failed")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "NewPostgresDB failed")
	}
	db.AutoMigrate(&datamodel.Organization{})
	postgresDB := PostgresDB{
		DB: db,
	}
	return &postgresDB, nil
}

func (p PostgresDB) AddOrganization(organization datamodel.Organization) error {
	tx := p.DB.Create(&organization)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "CreateOrganization failed")
	}
	return nil
}

func (p PostgresDB) GetOrganization(address string) (organization datamodel.Organization, err error) {
	tx := p.DB.First(&organization, "address = ?", address)
	if tx.Error != nil {
		err = errors.Wrap(tx.Error, "GetOrganization failed")
		return
	}
	return
}

func (p PostgresDB) GetOrganizations() ([]datamodel.Organization, error) {
	var organizations []datamodel.Organization

	tx := p.DB.Find(&organizations)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "GetOrganizations failed")
	}
	return organizations, nil
}
