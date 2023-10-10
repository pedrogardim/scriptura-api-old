package theographic

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	db "github.com/pedrogardim/scriptura-api/database"
	"github.com/pedrogardim/scriptura-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateIds() {
	ids := make(map[string]string)
	path := "data_importers/theographic/json/"
	files, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	//read files
	for i := range files {
		dat, err := os.ReadFile(path + files[i].Name())

		if err != nil {
			panic(err)
		}

		splitted := strings.Split(string(dat), `"`)
		for j := range splitted {
			if strings.HasPrefix(splitted[j], "rec") {
				ids[splitted[j]] = mongoObjectId()
			}
		}
	}

	fmt.Println(ids)

	json, _ := json.Marshal(ids)
	os.WriteFile(path+"id_map.json", json, 0644)
}

func mongoObjectId() string {
	ts := time.Now().UnixMilli() / 1000
	id := strconv.FormatInt(ts, 16)
	for i := 0; i < 16; i++ {
		id += fmt.Sprintf("%x", rand.Intn(16))
	}
	return id
}

///////

type OriginalVerse struct {
	ID     string         `json:"id"`
	Fields OriginalFields `json:"fields"`
}

type OriginalFields struct {
	OsisRef         string   `json:"osisRef"`
	VerseNum        string   `json:"verseNum"`
	VerseText       string   `json:"verseText"`
	Book            []string `json:"book"`
	People          []string `json:"people"`
	YearNum         int      `json:"yearNum"`
	Chapter         []string `json:"chapter"`
	Status          string   `json:"status"`
	VerseID         string   `json:"verseID"`
	Timeline        []string `json:"timeline"`
	Places          []string `json:"places"`
	PeopleGroups    []string `json:"peopleGroups"`
	EventsDescribed []string `json:"eventsDescribed"`
}

func ImportVerses() {
	idsPath := "utils/importers/theographic/json/id_map.json"
	idsJson, err := os.ReadFile(idsPath)
	if err != nil {
		panic(err)
	}
	var parsedIdMap map[string]string

	json.Unmarshal(idsJson, &parsedIdMap)

	// var importedJson []models.OriginalVerse
	var importedJson []OriginalVerse
	var parsed []models.Verse

	versesPath := "utils/importers/theographic/json/verses.json"
	versesJson, err := os.ReadFile(versesPath)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(versesJson, &importedJson)

	for i := range importedJson {
		idObjID, _ := primitive.ObjectIDFromHex(parsedIdMap[importedJson[i].ID])
		bookObjID, _ := primitive.ObjectIDFromHex(parsedIdMap[importedJson[i].Fields.Book[0]])
		chapterObjID, _ := primitive.ObjectIDFromHex(parsedIdMap[importedJson[i].Fields.Chapter[0]])

		verseId, _ := strconv.Atoi(importedJson[i].Fields.VerseID)
		verseNum, _ := strconv.Atoi(importedJson[i].Fields.VerseNum)

		people := make([]primitive.ObjectID, 0)
		places := make([]primitive.ObjectID, 0)
		peopleGroups := make([]primitive.ObjectID, 0)
		eventsDescribed := make([]primitive.ObjectID, 0)
		timeline := make([]primitive.ObjectID, 0)

		for _, person := range importedJson[i].Fields.People {
			person, _ := primitive.ObjectIDFromHex(parsedIdMap[person])
			people = append(people, person)
		}
		for _, place := range importedJson[i].Fields.Places {
			place, _ := primitive.ObjectIDFromHex(parsedIdMap[place])
			places = append(places, place)
		}
		for _, peopleGroup := range importedJson[i].Fields.PeopleGroups {
			peopleGroup, _ := primitive.ObjectIDFromHex(parsedIdMap[peopleGroup])
			peopleGroups = append(peopleGroups, peopleGroup)
		}
		for _, eventDescribed := range importedJson[i].Fields.EventsDescribed {
			eventDescribed, _ := primitive.ObjectIDFromHex(parsedIdMap[eventDescribed])
			eventsDescribed = append(eventsDescribed, eventDescribed)
		}
		for _, timelineItem := range importedJson[i].Fields.EventsDescribed {
			timelineItem, _ := primitive.ObjectIDFromHex(parsedIdMap[timelineItem])
			timeline = append(timeline, timelineItem)
		}

		verse := models.Verse{
			ID:        idObjID,
			OsisRef:   importedJson[i].Fields.OsisRef,
			Status:    importedJson[i].Fields.Status,
			VerseID:   verseId,
			Book:      bookObjID,
			Chapter:   chapterObjID,
			VerseNum:  verseNum,
			VerseText: importedJson[i].Fields.VerseText,
			YearNum:   importedJson[i].Fields.YearNum,
		}

		if len(people) > 0 {
			verse.People = people
		}

		if len(places) > 0 {
			verse.Places = places
		}

		if len(peopleGroups) > 0 {
			verse.PeopleGroups = peopleGroups
		}

		if len(eventsDescribed) > 0 {
			verse.EventsDescribed = eventsDescribed
		}

		if len(timeline) > 0 {
			verse.Timeline = timeline
		}

		parsed = append(parsed, verse)
	}

	versesInterface := make([]interface{}, len(parsed))
	for i, v := range parsed {
		versesInterface[i] = v
	}

	// json, _ := json.Marshal(string(dat))
	result, err := db.Db.Collection("verses").InsertMany(db.Ctx, versesInterface)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted", result.InsertedIDs)
}
