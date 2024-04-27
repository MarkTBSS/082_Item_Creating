package main

import (
	"github.com/MarkTBSS/082_Item_Creating/config"
	"github.com/MarkTBSS/082_Item_Creating/databases"
	"github.com/MarkTBSS/082_Item_Creating/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
