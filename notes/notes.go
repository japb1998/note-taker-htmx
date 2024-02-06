package notes

import (
	"time"
		"fmt"
	)

type Note struct {
	Id int `form:"id" json:"id"`
	Title     string `form:"title" json:"title" binding:"required"`
	Body      string `form:"body" json:"body" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
}
type CreateNote struct {
	Title string `form:"title" json:"title" binding:"required"`
	Body string `form:"body" json:"body" binding:"required"`
}
var notes []Note = []Note{
	{	
		Id: 0,
		Title:     "First Note",
		Body:      "This is the body of the first note.",
		CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{	
		Id: 1,
		Title:     "Second Note",
		Body:      "This is the body of the second note.",
		CreatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
	},
	{	
		Id: 2,
		Title:     "Third Note",
		Body:      "This is the body of the third note.",
		CreatedAt: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
	},
}
func GetAll() []Note {
	return notes
}

func Add(n CreateNote) {
	note  := Note{
		Id: len(notes),
		Title: n.Title,
		Body: n.Body,
		CreatedAt: time.Now(),
	}
	notes = append(notes, note)
}

func Delete(id int) error {
	for idx, note := range notes {
		if note.Id == id {
			notes = append(notes[:idx], notes[idx+1:]...)
			return nil
		}

	}
	return fmt.Errorf("Note with id %d not found", id)
}
