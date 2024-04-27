package repository

import (
	"github.com/MarkTBSS/082_Item_Creating/entities"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}
