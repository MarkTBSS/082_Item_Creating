package controller

import (
	"net/http"

	_custom "github.com/MarkTBSS/082_Item_Creating/pkg/custom"
	_exception "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/exeption"
	_models "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/models"
	"github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService service.ItemShopService
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_models.ItemFilter)
	validatingContext := _custom.NewCustomEchoRequest(pctx)

	err := validatingContext.Bind(itemFilter)
	if err != nil {
		return _custom.CustomError(pctx, http.StatusBadRequest, err)
	}

	itemModelLists, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return _custom.CustomError(pctx, http.StatusInternalServerError, &_exception.ItemListing{})
	}
	return pctx.JSON(http.StatusOK, itemModelLists)
	// Custom Error Testing
	//return _custom.CustomError(pctx, http.StatusInternalServerError, &_exception.ItemListing{})
}

func NewItemShopControllerImpl(itemShopService service.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}
