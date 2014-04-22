package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func main() {

	// define our usage string here
	// TODO: fix long names
	usage := `note
Usage:
    note [-h] [-m TEXT] [-t TAGS] [-n NAME] 

    -h,--help                   : Show this help message
    -m TEXT, --message TEXT     : Include note text, if not included will open and editor
    -t TAGS, --tags TAGS        : include comma separated tags
    -n NAME, --name NAME        : Name of note file`

	arguments, _ := docopt.Parse(usage, nil, true, "notes_v0.1", false)
	fmt.Println(arguments)
}

//******************************************************************************
// Global variables and constants
//******************************************************************************
const EDITOR string = "vim"
const NOTEDIR string = "/Users/fgimenez/Dropbox/notes/"
const TAGID string = "#TAGS"

//******************************************************************************
// Struct definition
//******************************************************************************
type Note struct {
	text string
	path string
	tags []string
}

//******************************************************************************
// Struct Constructors
//******************************************************************************
func NewNote(noteText string, noteName string, tagString string) *Note {

	// TEXT
	if !noteText {
		noteText = ""
	}

	// PATH
	if !noteName {
		noteName = time.Now().Format("2006-01-02_15:04:05")
	}

	if !filepath.IsAbs(noteName) {
		notePath = filepath.Join(NOTEDIR, noteName)
	} else {
		notePath = noteName
	}

	// TAGS
	if !tagString {
		tagString = TAGID
	}

	noteTags = parseTagString(tagString)

	return &Note{text: noteText, path: notePath, tags: noteTags}
}

//******************************************************************************
// Utility functions
//******************************************************************************

// Parse a tag string into a set of unique tags
func parseTagString(tagString string) []string {
	startIdx := strings.Index(tagString, TAGID)
	if startIdx == -1 {
		startIdx = 0
	} else {
		startIdx += len(TAGID)
	}

	// remove tagid
	tagString = tagString[startIdx:]

	// split into tags
	tags := strings.Split(tagString, ",")

	// trim all tags of leading/training whitespace
	for i, t := range tags {
		tags[i] = strings.Trim(t, " \t")
	}

	return tags
}
