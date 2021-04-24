package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Incident struct {
		ID         primitive.ObjectID `bson:"_id" json:"id"`
		SensorID   string             `bson:"sensorID" json:"sensorID"`
		CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
		SoundLevel int                `bson:"soundLevel" json:"soundLevel"`
	}
)

func NewIncident(sensorID string, soundLvl int) *Incident {
	return &Incident{
		ID:         primitive.NewObjectID(),
		SensorID:   sensorID,
		CreatedAt:  time.Now().UTC(),
		SoundLevel: soundLvl,
	}
}
