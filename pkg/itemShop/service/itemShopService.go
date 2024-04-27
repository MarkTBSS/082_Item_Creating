package service

import _models "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/models"

type ItemShopService interface {
	Listing(itemFilter *_models.ItemFilter) (*_models.ItemResult, error)
}
