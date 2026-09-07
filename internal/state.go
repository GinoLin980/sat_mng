package internal

type Vocabulary struct {
	Word        string
	WordType    string
	Description string
}

var PendingWords = map[string]Vocabulary{}

const Filename = "words.dat"

// set initial value to true, as unmodified means the disk and ram are synced, which equal to saved
var Saved bool = true
