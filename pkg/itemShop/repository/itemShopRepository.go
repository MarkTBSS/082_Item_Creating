package repository

import (
	_entities "github.com/MarkTBSS/082_Item_Creating/entities"
	_models "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/models"
)

type ItemShopRepository interface {
	Listing(itemFilter *_models.ItemFilter) ([]*_entities.Item, error)
	Counting(itemFilterDto *_models.ItemFilter) (int64, error)
}
