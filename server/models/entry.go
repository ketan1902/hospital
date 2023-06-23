package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID          primitive.ObjectID `bson: "id"`
	PatientName *string            `json: "patientName"`
	Disease     *string            `json: "disease"`
	Age         *float64           `json: "age"`
	ContactNo   *float64           `json: "contactNo"`
}
