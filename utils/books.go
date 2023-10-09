package utils

import (
	"strings"
)

var BooksAbbreviations = [][]string{
	{"Genesis", "Gen", "Ge", "Gn"},
	{"Exodus", "Ex", "Exo"},
	{"Leviticus", "Lev", "Lv"},
	{"Numbers", "Num", "Nu", "Nm", "Nb"},
	{"Deuteronomy", "Deut", "Dt"},
	{"Joshua", "Josh", "Jos", "Jsh"},
	{"Judges", "Judg", "Jdg", "Jdgs", "Jg", "Jgs"},
	{"Ruth", "Rth", "Rt"},
	{"1 Samuel", "1 Sam", "1 Sm", "1 Sa", "1S"},
	{"2 Samuel", "2 Sam", "2 Sm", "2 Sa", "2S"},
	{"1 Kings", "1 Kgs", "1 Ki", "1K"},
	{"2 Kings", "2 Kgs", "2 Ki", "2K"},
	{"1 Chronicles", "1 Chr", "1 Ch", "1Ch"},
	{"2 Chronicles", "2 Chr", "2 Ch", "2Ch"},
	{"Ezra", "Ezr"},
	{"Nehemiah", "Neh", "Ne"},
	{"Esther", "Est", "Esth"},
	{"Job", "Jb"},
	{"Psalms", "Ps", "Pslm"},
	{"Proverbs", "Prov", "Pr", "Prv"},
	{"Ecclesiastes", "Ecc", "Eccl", "Eccles"},
	{"Song of Solomon", "Song", "SOS", "So"},
	{"Isaiah", "Isa", "Is"},
	{"Jeremiah", "Jer", "Je", "Jr"},
	{"Lamentations", "Lam", "La"},
	{"Ezekiel", "Ezek", "Eze", "Ezk"},
	{"Daniel", "Dan", "Da", "Dn"},
	{"Hosea", "Hos", "Ho"},
	{"Joel", "Joel", "Jl"},
	{"Amos", "Am"},
	{"Obadiah", "Obad", "Ob"},
	{"Jonah", "Jon", "Jnh"},
	{"Micah", "Mic", "Mc"},
	{"Nahum", "Nah", "Na"},
	{"Habakkuk", "Hab", "Hb"},
	{"Zephaniah", "Zeph", "Zep", "Zp"},
	{"Haggai", "Hag", "Hg"},
	{"Zechariah", "Zech", "Zec", "Zc"},
	{"Malachi", "Mal", "Ml"},
	{"Matthew", "Matt", "Mt"},
	{"Mark", "Mk", "Mrk"},
	{"Luke", "Lk", "Luk"},
	{"John", "Jn", "Jhn"},
	{"Acts", "Acts", "Ac"},
	{"Romans", "Rom", "Ro", "Rm"},
	{"1 Corinthians", "1 Cor", "1 Co", "1C"},
	{"2 Corinthians", "2 Cor", "2 Co", "2C"},
	{"Galatians", "Gal", "Ga"},
	{"Ephesians", "Eph", "Ephes"},
	{"Philippians", "Phil", "Php", "Pp"},
	{"Colossians", "Col", "Colo"},
	{"1 Thessalonians", "1 Thess", "1 Th", "1Th"},
	{"2 Thessalonians", "2 Thess", "2 Th", "2Th"},
	{"1 Timothy", "1 Tim", "1 Ti", "1Tm"},
	{"2 Timothy", "2 Tim", "2 Ti", "2Tm"},
	{"Titus", "Tit", "Tt"},
	{"Philemon", "Philem", "Phm", "Pm"},
	{"Hebrews", "Heb"},
	{"James", "Jas", "Jm"},
	{"1 Peter", "1 Pet", "1 Pe", "1 Pt", "1P"},
	{"2 Peter", "2 Pet", "2 Pe", "2 Pt", "2P"},
	{"1 John", "1 Jn", "1 Jhn", "1J"},
	{"2 John", "2 Jn", "2 Jhn", "2J"},
	{"3 John", "3 Jn", "3 Jhn", "3J"},
	{"Jude", "Jud"},
	{"Revelation", "Rev", "Re", "Rv"},
}

func GetBookIndex(input string) int {
	for i, book := range BooksAbbreviations {
		for _, abbrev := range book {
			abbrev = strings.ReplaceAll(strings.ToLower(abbrev), " ", "")
			if strings.Contains(input, abbrev) {
				return i + 1
			}
		}
	}
	return -1
}
