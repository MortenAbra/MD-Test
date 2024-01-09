package repositories

import (
	"context"
	"errors"
	"fmt"
	"media-devoted/db"
	"media-devoted/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RocketRepository interface {
	GetRockets(ctx context.Context) (*[]types.Rocket, error)
	GetRocket(ctx context.Context, id *uuid.UUID) (*types.Rocket, error)
	AddRocket(ctx context.Context, rocket *types.Rocket) error
	UpdateRocket(ctx context.Context, rocket *types.Rocket) error
	DeleteRocket(ctx context.Context, id uuid.UUID) error
}

type RocketRepositoryImpl struct {
	db *gorm.DB
}

func RocketRepositoryInstance() RocketRepository {
	return &RocketRepositoryImpl{
		db: db.GetDB(),
	}
}

func (r *RocketRepositoryImpl) GetRockets(ctx context.Context) (*[]types.Rocket, error) {
	var rockets *[]types.Rocket
	// Retrieves all rockets
	if result := r.db.Table("rockets").Find(&rockets); result.Error != nil {
		return rockets, result.Error
	}

	return rockets, nil
}

func (r *RocketRepositoryImpl) GetRocket(ctx context.Context, id *uuid.UUID) (*types.Rocket, error) {
	var rocket *types.Rocket
	// Retreving rocket based on id
	if result := r.db.Table("rockets").Where("id = ?", id).First(&rocket); result.Error != nil {
		return rocket, result.Error
	}

	return rocket, nil

}

func (r *RocketRepositoryImpl) AddRocket(ctx context.Context, rocket *types.Rocket) error {
	// Creates a new rocket in rockets table
	if result := r.db.Table("rockets").Create(&rocket); result.Error != nil {
		return result.Error
	}

	return nil
}

// Transaction used to ensure only one instance can modify/update the rocket at once
func (r *RocketRepositoryImpl) UpdateRocket(ctx context.Context, rocket *types.Rocket) error {
	// Start a new transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error // handle transaction start error
	}

	// Ensure that the transaction is rolled back in case of panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingRocket types.Rocket
	// Perform the query within the transaction
	if err := tx.Table("rockets").Where("id = ?", rocket.Id).First(&existingRocket).Error; err != nil {
		tx.Rollback() // Roll back in case of error
		return err
	}

	// Update the fields of existingRocket
	existingRocket.Name = rocket.Name
	existingRocket.Mission = rocket.Mission
	existingRocket.Speed = rocket.Speed

	if err := tx.Table("rockets").Session(&gorm.Session{FullSaveAssociations: true}).Updates(existingRocket).Error; err != nil {
		tx.Rollback() // Roll back in case of error
		return err
	}

	// Commit the transaction if no errors
	if err := tx.Commit().Error; err != nil {
		return err // handle commit error
	}

	return nil
}

func (r *RocketRepositoryImpl) DeleteRocket(ctx context.Context, id uuid.UUID) error {
	var rocket *types.Rocket

	result := r.db.Table("rockets").First(&rocket, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Rocket does exists
			return fmt.Errorf("rocket with id: %s not found", id)
		}
		return result.Error
	}

	if err := r.db.Table("rockets").Delete(&rocket, id); err != nil {
		return err.Error
	}

	return nil
}
