package main

func DBFetchContactList(contacts *[]Contact) error {
	err := db.Model(&Contact{}).Find(&contacts).Error
	return err
}

func DBCreateContact(contact *Contact) error {
	err := db.Create(contact).Error
	return err
}

func DBCreateTask(task *Task) error {
	err := db.Create(task).Error
	return err
}

func DBListTask(tasks *[]Task) error {
	err := db.Model(&Task{}).Preload("Contact").Find(tasks).Error
	return err
}

func DBGetTask(task *Task, taskID uint64) error {
	err := db.Preload("Contact").Preload("Reminders").First(task, taskID).Error
	return err
}

func DBUpdateTask(task *Task) error {
	err := db.Save(task).Error
	return err
}

func DBDeleteTask(taskID uint64) error {
	err := db.Delete(&Task{}, taskID).Error
	return err
}

// Reminders DB Queries

func DBGetReminder(reminder *Reminder, reminderID uint64) error {
	err := db.Preload("Task.Contact").First(reminder, reminderID).Error
	return err
}

func DBUpdateReminder(reminder *Reminder) error {
	err := db.Save(reminder).Error
	return err
}

func DBDeleteReminder(reminderID uint64) error {
	err := db.Delete(&Reminder{}, reminderID).Error
	return err
}

func DBCreateReminder(reminder *Reminder) error {
	err := db.Create(reminder).Error
	return err
}

func DBListReminder(reminders *[]Reminder) error {
	err := db.Model(&Reminder{}).Preload("Task.Contact").Find(reminders).Error
	return err
}
