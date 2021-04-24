package database

import (
	"github.com/ra9dev/safe-and-sound/internal/police-server/database/drivers"
	"github.com/ra9dev/safe-and-sound/internal/police-server/database/drivers/mongo"
)

func New(conf drivers.DataStoreConfig) drivers.DataStore {
	switch conf.Name {
	case "mongo":
		return mongo.New(conf)
	default:
		panic("no default db")
	}
}
