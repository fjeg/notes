package main

import (
	"reflect"
	"testing"
)

func TestNote(t *testing.T) {
	/*
	   n1 := &Note{
	       text: "correct horse battery staple",
	       path: "/Users/fgimenez/Dropbox/notes/n1",
	       tags: []string{"t1","t2"},
	   }
	*/

	n2 := NewNote("correct horse battery staple", "test", "#TAGS t1,t2")

	t.Log(reflect.TypeOf(n2))
	t.Log(n2.TagString())
	t.Log(n2)

	n2.GetTextFromEditor()

	//err := n2.WriteNote()
	/*
	   if *n1 != *n2 {
	       t.Error("Constructor failed.\n%+v\n%+v",*n1,*n2)
	   }
	*/

}
