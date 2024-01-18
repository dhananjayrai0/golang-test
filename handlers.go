package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

// Contact Handlers

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newContact Contact

	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON format"})
		return
	}

	newContact.CreatedAt = time.Now()

	err = db.Create(&newContact).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newContact)
}

func listContact(w http.ResponseWriter, r *http.Request) {
	var contacts []Contact
	err := db.Model(&Contact{}).Find(&contacts).Error

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(contacts)
}

// Task Handlers

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON format"})
		return
	}

	newTask.CreatedAt = time.Now()

	err = db.Create(&newTask).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func listTask(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	err := db.Model(&Task{}).Preload("Contact").Find(&tasks).Error

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskIDStr := chi.URLParam(r, "taskID")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid task ID"})
		return
	}

	var task Task
	err = db.Preload("Contact").Preload("Reminders").First(&task, taskID).Error
	if gorm.IsRecordNotFoundError(err) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskIDStr := chi.URLParam(r, "taskID")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid task ID"})
		return
	}

	var updatedTask Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON format"})
		return
	}

	updatedTask.ID = uint(taskID)

	err = db.Save(&updatedTask).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(updatedTask)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskIDStr := chi.URLParam(r, "taskID")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid task ID"})
		return
	}

	err = db.Delete(&Task{}, taskID).Error
	if gorm.IsRecordNotFoundError(err) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Reminder Handlers

func getReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reminderIDStr := chi.URLParam(r, "reminderID")
	reminderID, err := strconv.ParseUint(reminderIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid reminder ID"})
		return
	}

	var reminder Reminder
	err = db.Preload("Task.Contact").First(&reminder, reminderID).Error
	if gorm.IsRecordNotFoundError(err) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Reminder not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(reminder)
}

func updateReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reminderIDStr := chi.URLParam(r, "reminderID")
	reminderID, err := strconv.ParseUint(reminderIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid reminder ID"})
		return
	}

	var updatedReminder Reminder
	err = json.NewDecoder(r.Body).Decode(&updatedReminder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON format"})
		return
	}

	updatedReminder.ID = uint(reminderID)

	err = db.Save(&updatedReminder).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(updatedReminder)
}

func deleteReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reminderIDStr := chi.URLParam(r, "reminderID")
	reminderID, err := strconv.ParseUint(reminderIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid reminder ID"})
		return
	}

	err = db.Delete(&Reminder{}, reminderID).Error
	if gorm.IsRecordNotFoundError(err) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Reminder not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func createReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newReminder Reminder

	err := json.NewDecoder(r.Body).Decode(&newReminder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON format"})
		return
	}

	newReminder.CreatedAt = time.Now()

	err = db.Create(&newReminder).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newReminder)
}

func listReminder(w http.ResponseWriter, r *http.Request) {
	var reminders []Reminder
	err := db.Model(&Reminder{}).Preload("Task.Contact").Find(&reminders).Error

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(reminders)
}

func main() {
	r := chi.NewRouter()

	// Contact Routers
	r.Post("/contacts", createContact)
	r.Get("/contacts", listContact)

	// Task Routers
	r.Post("/tasks", createTask)
	r.Get("/tasks", listTask)
	r.Get("/tasks/{taskID}", getTask)
	r.Put("/tasks/{taskID}", updateTask)
	r.Delete("/tasks/{taskID}", deleteTask)

	// Reminder Routers
	r.Post("/reminders", createReminder)
	r.Get("/reminders", listReminder)
	r.Get("/reminders/{reminderID}", getReminder)
	r.Put("/reminders/{reminderID}", updateReminder)
	r.Delete("/reminders/{reminderID}", deleteReminder)

	http.ListenAndServe(":8080", r)
}
