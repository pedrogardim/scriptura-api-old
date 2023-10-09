package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/pedrogardim/scriptura-api/database"
	"github.com/pedrogardim/scriptura-api/models"
	"github.com/pedrogardim/scriptura-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetVerse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prcdRef, err := utils.ProcessVerseQuery(r.URL.Path)
	if err != nil {
		http.Error(w, "Error processing verse query", http.StatusBadRequest)
		return
	}
	verseID, err := strconv.Atoi(prcdRef)
	if err != nil {
		http.Error(w, "Error converting reference to integer", http.StatusBadRequest)
		return
	}

	fmt.Println(verseID)

	filter := bson.M{
		"verseID": verseID,
	}

	var verse models.Verse
	db.Db.Collection("verses").FindOne(db.Ctx, filter).Decode(&verse)

	json.NewEncoder(w).Encode(verse)
}
