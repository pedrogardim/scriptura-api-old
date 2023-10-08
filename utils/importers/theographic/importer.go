package theographic

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
