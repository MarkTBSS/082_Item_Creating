package server

import (
	"github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/controller"
	"github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/repository"
	"github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := repository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := service.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := controller.NewItemShopControllerImpl(itemShopService)

	router.GET("/listing", itemShopController.Listing)
}
