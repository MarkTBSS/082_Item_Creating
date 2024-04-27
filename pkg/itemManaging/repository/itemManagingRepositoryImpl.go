package repository

import (
	"github.com/MarkTBSS/082_Item_Creating/databases"
	"github.com/MarkTBSS/082_Item_Creating/entities"
	"github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/exception"
	"github.com/labstack/echo/v4"
)

type itemMangingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemMangingRepositoryImpl{db, logger}
}

func (r *itemMangingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)
	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Error("Item creating failed:", err.Error())
		return nil, &exception.ItemCreating{}
	}
	return item, nil
}
