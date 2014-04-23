package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"os"
//    "os/exec"
	"path/filepath"
	"strings"
	"time"
    "log"
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

	// PATH
	if noteName == "" {
		noteName = time.Now().Format("2006-01-02_15:04:05")
	}

	var notePath string
	if !filepath.IsAbs(noteName) {
		notePath = filepath.Join(NOTEDIR, noteName)
	} else {
		notePath = noteName
	}

    if _,err := os.Stat(notePath); os.IsExist(err){
        log.Fatal("Note already exists, choose new name.", notePath)
    }

	// TAGS
	if tagString == "" {
		tagString = TAGID
	}
	noteTags := parseTagString(tagString)

	return &Note{text: noteText, path: notePath, tags: noteTags}
}

//******************************************************************************
// Note methods
//******************************************************************************

// write out note text
// possibly return error in the future
func (n *Note) WriteNote() {

    //TODO check if file exists and handle name collisions

    // if not, create the file
    outFile,err := os.Create(n.path)
    if err != nil {
        log.Fatal("WriteNote:",err)
    }
    defer outFile.Close()

    // write the file out
    _,err = outFile.WriteString(n.text + "\n" + n.TagString())
    if err != nil{
        log.Fatal("WriteNote:",err)
    }


}

// convert note tags to string
func (n *Note) TagString() string {
    return TAGID + " " + strings.Join(n.tags,",")

}

// get text from editor
func (n *Note) GetTextFromEditor() {
//http://stackoverflow.com/questions/12088138/trying-to-launch-an-external-editor-from-within-a-go-program

}


//******************************************************************************
// Utility functions
//******************************************************************************

// Parse a tag string into a set of unique tags
func parseTagString(tagString string) []string {

	// remove tagid
	startIdx := strings.Index(tagString, TAGID)

	if startIdx == -1 {
		startIdx = 0
	} else {
		startIdx += len(TAGID)
	}

	tagString = tagString[startIdx:]

	// split into tags
	tags := strings.Split(tagString, ",")

	// trim all tags of leading/training whitespace
	for i, t := range tags {
		tags[i] = strings.Trim(t, " \t")
	}

	return tags
}
