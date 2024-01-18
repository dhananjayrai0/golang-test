package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var contactID uint
var taskID uint

func TestCreateContact(t *testing.T) {
	const name = "D. Rai"
	payload := map[string]interface{}{
		"name": name,
	}

	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/contacts", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	createContact(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdContact Contact
	err = json.Unmarshal(rr.Body.Bytes(), &createdContact)
	if err != nil {
		t.Fatal(err)
	}
	contactID = createdContact.ID
	assert.Equal(t, name, createdContact.Name)
}

func TestListContact(t *testing.T) {
	req, err := http.NewRequest("GET", "/contacts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	listContact(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var rawContacts interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &rawContacts)
	if err != nil {
		t.Fatal(err)
	}

	contacts := rawContacts.([]interface{})

	assert.GreaterOrEqual(t, len(contacts), 0)
}

func TestCreateTask(t *testing.T) {
	const (
		title       = "Sample Task"
		description = "This is a sample task"
		priority    = "high"
		dueDateTime = "2024-01-19T01:00:00Z"
	)

	payload := map[string]interface{}{
		"Title":       title,
		"Description": description,
		"Priority":    priority,
		"ContactID":   contactID,
	}

	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	createTask(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdTask Task
	err = json.Unmarshal(rr.Body.Bytes(), &createdTask)
	if err != nil {
		t.Fatal(err)
	}
	taskID = createdTask.ID
	assert.Equal(t, title, createdTask.Title)
	assert.Equal(t, description, createdTask.Description)
	assert.Equal(t, priority, createdTask.Priority)
}

func TestListTask(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	listTask(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var rawTasks interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &rawTasks)
	if err != nil {
		t.Fatal(err)
	}

	tasks := rawTasks.([]interface{})

	assert.GreaterOrEqual(t, len(tasks), 0)
}

func TestCreateReminder(t *testing.T) {
	const reminderTime = "2024-01-10T01:00:00Z"

	payload := map[string]interface{}{
		"TaskID":       1,
		"ReminderTime": reminderTime,
	}

	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/reminders", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	createReminder(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdReminder Reminder
	err = json.Unmarshal(rr.Body.Bytes(), &createdReminder)
	if err != nil {
		t.Fatal(err)
	}

	reminderTimeParsed, err := time.Parse(time.RFC3339, reminderTime)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, reminderTimeParsed, createdReminder.ReminderTime)
}

func TestListReminder(t *testing.T) {
	req, err := http.NewRequest("GET", "/reminders", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	listReminder(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var rawReminders interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &rawReminders)
	if err != nil {
		t.Fatal(err)
	}

	reminders := rawReminders.([]interface{})

	assert.GreaterOrEqual(t, len(reminders), 0)
}
