package server

import (
	"github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/controller"
	"github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/repository"
	"github.com/MarkTBSS/082_Item_Creating/pkg/itemManaging/service"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemMangingRepository := repository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := service.NewItemManagingServiceImpl(itemMangingRepository)
	itemManaging := controller.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManaging.Creating)
}
