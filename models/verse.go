package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Verse struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	OsisRef         string               `bson:"osisRef,omitempty" json:"osisRef,omitempty"`
	Status          string               `bson:"status,omitempty" json:"status,omitempty"`
	VerseID         int                  `bson:"verseID,omitempty" json:"verseID,omitempty"`
	Book            primitive.ObjectID   `bson:"book,omitempty" json:"book,omitempty"`
	Chapter         primitive.ObjectID   `bson:"chapter,omitempty" json:"chapter,omitempty"`
	VerseNum        int                  `bson:"verseNum,omitempty" json:"verseNum,omitempty"`
	VerseText       string               `bson:"verseText,omitempty" json:"verseText,omitempty"`
	YearNum         int                  `bson:"yearNum,omitempty" json:"yearNum,omitempty"`
	People          []primitive.ObjectID `bson:"people,omitempty" json:"people,omitempty"`
	Places          []primitive.ObjectID `bson:"places,omitempty" json:"places,omitempty"`
	PeopleGroups    []primitive.ObjectID `bson:"peopleGroups,omitempty" json:"peopleGroups,omitempty"`
	EventsDescribed []primitive.ObjectID `bson:"eventsDescribed,omitempty" json:"eventsDescribed,omitempty"`
}
