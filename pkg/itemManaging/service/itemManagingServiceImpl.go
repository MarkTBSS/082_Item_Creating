package service

import (
	"github.com/MarkTBSS/082_Item_Creating/entities"
	_itemManagingModel "github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/models"
	_itemManagingRepository "github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/repository"
	_itemShopModel "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/models"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{
		itemManagingRepository,
	}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	item := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	itemEntity, err := s.itemManagingRepository.Creating(item)
	if err != nil {
		return nil, err
	}

	return itemEntity.ChangeToItemModel(), nil
}
