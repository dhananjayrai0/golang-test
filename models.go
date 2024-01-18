package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	// We can store more informations like phone number, email etc
	gorm.Model
	Name  string
	Tasks []Task
}

type Task struct {
	gorm.Model
	ContactID   uint
	Contact     Contact
	Title       string
	Description string
	Priority    string
	DueDateTime time.Time
	Reminders   []Reminder
}

type Reminder struct {
	gorm.Model
	TaskID       uint
	Task         Task
	ReminderTime time.Time
}
