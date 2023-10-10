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

func GetVerse(c *gin.Context) {
	queryString := c.Param("query_string")
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
