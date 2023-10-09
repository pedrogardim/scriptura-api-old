package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ProcessVerseQuery(query string) (string, error) {
	ref := strings.TrimPrefix(query, "/api/verse/")
	//get book index
	bookIndex := GetBookIndex(ref)

	if bookIndex == -1 {
		return "", fmt.Errorf("invalid book prefix")
	}

	splittedQuery := strings.Split(query, ":")
	re := regexp.MustCompile("[^0-9]+")

	chapterNum := re.ReplaceAllString(splittedQuery[0], "")
	verseNum := re.ReplaceAllString(splittedQuery[1], "")

	bookNum, _ := formatTwoDigits(strconv.Itoa(bookIndex))
	chapterNum, _ = formatThreeDigits(chapterNum)
	verseNum, _ = formatThreeDigits(verseNum)

	ref = bookNum + chapterNum + verseNum
	return ref, nil
}

func formatThreeDigits(s string) (string, error) {
	switch len(s) {
	case 1:
		return "00" + s, nil
	case 2:
		return "0" + s, nil
	case 3:
		return s, nil
	case 0:
		return "", errors.New("input string is empty")
	default:
		return "", errors.New("input string is too long")
	}
}

func formatTwoDigits(s string) (string, error) {
	switch len(s) {
	case 1:
		return "0" + s, nil
	case 2:
		return s, nil
	case 0:
		return "", errors.New("input string is empty")
	default:
		return "", errors.New("input string is too long")
	}
}
