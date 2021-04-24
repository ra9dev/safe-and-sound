package drivers

import (
	"context"
	"github.com/ra9dev/safe-and-sound/internal/models"
)

type DataStore interface {
	Name() string
	Close(ctx context.Context) error
	Connect() error

	Incidents() IncidentsRepository
}

type IncidentsRepository interface {
	Create(ctx context.Context, incident *models.Incident) error
	All(ctx context.Context) ([]*models.Incident, error)
}
