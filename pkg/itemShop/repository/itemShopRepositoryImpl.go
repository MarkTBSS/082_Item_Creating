package repository

import (
	"github.com/MarkTBSS/082_Item_Creating/databases"
	_entities "github.com/MarkTBSS/082_Item_Creating/entities"
	_exception "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/exeption"
	_models "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/models"
	"github.com/labstack/echo/v4"
)

type itemRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func (r *itemRepositoryImpl) Listing(itemFilter *_models.ItemFilter) ([]*_entities.Item, error) {
	query := r.db.Connect().Model(&_entities.Item{}).Where("is_archive = ?", false)
	itemLists := make([]*_entities.Item, 0)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	size := int(itemFilter.Size)

	err := query.Offset(offset).Limit(size).Find(&itemLists).Order("id asc").Error
	if err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &_exception.ItemListing{}
	}
	return itemLists, nil
	// Error Testing
	//return nil, &_exception.ItemListing{}
}

func (r *itemRepositoryImpl) Counting(itemFilter *_models.ItemFilter) (int64, error) {
	query := r.db.Connect().Model(&_entities.Item{}).Where("is_archive = ?", false)
	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}
	var count int64
	//count := new(int64)

	err := query.Count(&count).Error
	if err != nil {
		r.logger.Error("Counting items failed : ", err.Error())
		return -1, &_exception.ItemCounting{}
	}
	return count, nil
	// Error Testing
	//return nil, &_exception.ItemListing{}
}
func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}
