package daemons

import (
	"context"
	"encoding/json"
	"github.com/ra9dev/safe-and-sound/internal/models"
	"github.com/ra9dev/safe-and-sound/internal/police-server/database/drivers"
	"log"
	"net/http"
	"sync"
	"time"
)

type SensorsWatcher struct {
	ctx    context.Context
	client *http.Client

	sensors             []string
	incidentsRepo       drivers.IncidentsRepository
	minIncidentSoundLvl int

	pause time.Duration
}

func NewSensorsWatcher(ctx context.Context, sensors []string, incidentsRepo drivers.IncidentsRepository) *SensorsWatcher {
	return &SensorsWatcher{
		ctx: ctx,
		client: &http.Client{
			Timeout: time.Second * 3,
		},

		sensors:             sensors,
		incidentsRepo:       incidentsRepo,
		minIncidentSoundLvl: 80,

		pause: time.Second * 10,
	}
}

func (w SensorsWatcher) Run() {
	log.Println("Watcher started watching")

	for {
		select {
		case <-w.ctx.Done():
			log.Println("Watcher terminated")
			return
		default:
			w.checkSensors()
			time.Sleep(w.pause)
		}
	}
}

func (w SensorsWatcher) checkSensors() {
	log.Println("checking sensors...")

	wg := new(sync.WaitGroup)
	for _, sensor := range w.sensors {
		wg.Add(1)

		go func(sensor string) {
			defer wg.Done()
			log.Printf("checking sensor with address [%s]...", sensor)

			resp, err := w.client.Get(sensor)
			if err != nil {
				log.Printf("[ERROR] could not check sensor with address [%s]: %+v", sensor, err)
				return
			}
			defer resp.Body.Close()

			sensorLog := new(models.SensorLog)
			if err := json.NewDecoder(resp.Body).Decode(sensorLog); err != nil {
				log.Printf("[ERROR] could not parse sensor log from address [%s]: %+v", sensor, err)
				return
			}

			if sensorLog.SoundLevel >= w.minIncidentSoundLvl {
				log.Printf("[WARN] incident at sensor [%s], sound level: %d", sensor, sensorLog.SoundLevel)
				if err := w.incidentsRepo.Create(w.ctx, models.NewIncident(sensorLog.ID, sensorLog.SoundLevel)); err != nil {
					log.Printf("[ERROR] could not create incident: %+v", err)
					return
				}
			}
		}(sensor)
	}

	wg.Wait()
}
