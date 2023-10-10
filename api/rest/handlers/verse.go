package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/pedrogardim/scriptura-api/database"
	"github.com/pedrogardim/scriptura-api/models"
	"github.com/pedrogardim/scriptura-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Get a verse by standard reference
// @Description Get a verse by standard reference
// @ID get-verse
// @Accept  json
// @Produce  json
// @Param   stdRef     path    string     true        "Standard reference"
// @Success 200 {string} string  "ok"
// @Router /api/verse/{stdRef} [get]
func GetVerse(c *gin.Context) {
	queryString := c.Param("std_ref")
	prcdRef, err := utils.ProcessVerseQuery(queryString)
	if err != nil {
		c.Error(errors.New("error processing verse query"))
	}
	verseID, err := strconv.Atoi(prcdRef)
	if err != nil {
		c.Error(errors.New("error converting reference to integer"))
	}

	fmt.Println(verseID)

	filter := bson.M{
		"verseID": verseID,
	}

	var verse models.Verse
	db.Db.Collection("verses").FindOne(db.Ctx, filter).Decode(&verse)

	c.JSON(200, verse)
}
