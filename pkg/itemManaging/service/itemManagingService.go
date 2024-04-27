package service

import (
	_itemManagingModel "github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/models"
	_itemShopModel "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/models"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
}
