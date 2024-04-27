package controller

import (
	"net/http"

	"github.com/MarkTBSS/082_Item_Creating/pkg/custom"
	_itemManagingService "github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/service"
	"github.com/labstack/echo/v4"

	_itemManagingModel "github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/models"
)

type itemManagingImpl struct {
	itemMangingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemMangingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingImpl{itemMangingService: itemMangingService}
}

func (c *itemManagingImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)
	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.Bind(itemCreatingReq); err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err)
	}
	item, err := c.itemMangingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusCreated, item)
}
