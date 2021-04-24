package mongo

import (
	"context"
	"github.com/ra9dev/safe-and-sound/internal/models"
	"github.com/ra9dev/safe-and-sound/internal/police-server/database/drivers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Incidents struct {
	col *mongo.Collection
}

func NewIncidents(col *mongo.Collection) drivers.IncidentsRepository {
	return &Incidents{col: col}
}

func (i Incidents) Create(ctx context.Context, incident *models.Incident) error {
	if incident == nil {
		return drivers.ErrEmptyStruct
	}

	if _, err := i.col.InsertOne(ctx, incident); err != nil {
		return err
	}

	return nil
}

func (i Incidents) All(ctx context.Context) ([]*models.Incident, error) {
	incidents := make([]*models.Incident, 0)

	cur, err := i.col.Find(ctx, bson.D{})
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return incidents, nil
		default:
			return nil, err
		}
	}

	if err := cur.All(ctx, &incidents); err != nil {
		return nil, err
	}

	return incidents, nil
}
