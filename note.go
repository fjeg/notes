package main

// http://blog.gopheracademy.com/vimgo-development-environment

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	// define our usage string here
	// TODO: fix long names

	/*
			usage := `note
		Usage:
		    note [-h] [-m TEXT] [-t TAGS] [-n NAME]

		    -h,--help                   : Show this help message
		    -m TEXT, --message TEXT     : Include note text, if not included will open and editor
		    -t TAGS, --tags TAGS        : include comma separated tags
		    -n NAME, --name NAME        : Name of note file`

			arguments, _ := docopt.Parse(usage, nil, true, "notes_v0.1", false)
	*/

	n2 := NewNote("correct horse battery staple", "test", "#TAGS t1,t2")
	n2.GetTextFromEditor()
	fmt.Print(n2.text)
	//fmt.Println("*******************")
	//fmt.Println(n2.ToString())
}

//******************************************************************************
// Global variables and constants
//******************************************************************************
const EDITOR string = "vim"
const NOTEDIR string = "/Users/fgimenez/Dropbox/notes/"
const TAGID string = "#TAGS"

//******************************************************************************
// Note definition
//******************************************************************************
type Note struct {
	text string
	path string
	tags []string
}

//******************************************************************************
// Note Constructors
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

	if _, err := os.Stat(notePath); os.IsExist(err) {
		log.Fatalf("Note already exists, choose new name.", notePath)
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
	outFile, err := os.Create(n.path)
	if err != nil {
		log.Fatalf("Error in WriteNote:", err)
	}
	defer outFile.Close()

	// write the file out
	_, err = outFile.WriteString(n.text + "\n" + n.TagString())
	if err != nil {
		log.Fatalf("WriteNote:", err)
	}

}

// convert note tags to string
func (n *Note) TagString() string {
	return TAGID + " " + strings.Join(n.tags, ",")

}

// get text from editor
func (n *Note) GetTextFromEditor() {

	// create temporary file and write tag string
	tmp, err := ioutil.TempFile("", "note_")
	if err != nil {
		log.Fatalf("Error opening temp file", err)
	}

	if _, err := tmp.WriteString(n.text + "\n"); err != nil {
		log.Fatalf("Error writing note text before editor %v", err)
	}

	if _, err := tmp.WriteString(n.TagString()); err != nil {
		log.Fatalf("Error writing tagstring before editor %v", err)
	}

	fpath := tmp.Name()
	tmp.Close()
	defer os.Remove(fpath)

	// Run shell command to call editor
	cmd := exec.Command(EDITOR, fpath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error in calling editor %v\n", err)
	}

	// read in temp file after editing
	noteBuffer, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatalf("Error opening temp file %v\n", err)
	}
	noteText, tags := parseNoteString(string(noteBuffer))
	n.text = noteText
	n.tags = tags
}

// convert a note to a convenient string
func (n *Note) ToString() string {

	s := n.text
	s += "\nTags:\t"

	for _, tag := range n.tags {
		s += tag + "\t"
	}

	return s

}

//******************************************************************************
// Utility functions
//******************************************************************************

// parse full string of note contents
func parseNoteString(noteString string) (noteText string, tags []string) {

	scanner := bufio.NewScanner(strings.NewReader(noteString))
	for scanner.Scan() {
		line := scanner.Text()

		startIdx := strings.Index(line, TAGID)

		if startIdx == -1 {
			noteText += line + "\n"
		} else {
			tags = parseTagString(line)
		}
	}

	strings.Trim(noteText, "\r\n")

	return
}

// Parse a tag string into a set of unique tags
func parseTagString(tagString string) []string {

	// remove tagid
	tagString = strings.TrimPrefix(tagString, TAGID)

	// split into tags
	tags := strings.Split(tagString, ",")

	// trim all tags of leading/training whitespace
	for i, t := range tags {
		tags[i] = strings.Trim(t, " \t")
	}

	return tags
}
