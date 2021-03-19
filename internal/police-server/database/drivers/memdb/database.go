package memdb

import (
	"context"
	"github.com/hashicorp/go-memdb"
	"github.com/ra9dev/safe-and-sound/internal/police-server/database/drivers"
)

type Database struct {
	schema *memdb.DBSchema
	db     *memdb.MemDB
}

func (d Database) Name() string {
	return "memdb"
}

func (d Database) Close(ctx context.Context) error {
	return nil
}

func (d *Database) Connect() error {
	db, err := memdb.NewMemDB(d.schema)
	if err != nil {
		return err
	}
	d.db = db

	return nil
}

func New() drivers.DataStore {
	schema := &memdb.DBSchema{}

	return &Database{
		schema: schema,
	}
}
