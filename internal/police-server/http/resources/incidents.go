package resources

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ra9dev/safe-and-sound/internal/police-server/database/drivers"
	"net/http"
)

type IncidentsResource struct {
	repo drivers.IncidentsRepository
}

func NewIncidentsResource(repo drivers.IncidentsRepository) *IncidentsResource {
	return &IncidentsResource{repo: repo}
}

func (ir IncidentsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", ir.All)

	return r
}

func (ir IncidentsResource) All(w http.ResponseWriter, r *http.Request) {
	incidents, err := ir.repo.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, incidents)
}
