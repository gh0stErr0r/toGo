package db

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"toGo/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DbNote *gorm.DB // Database for notes
	DbTask *gorm.DB // Database for tasks
	err    error
	format = time.UnixDate // Date format for storage
)

// Dinit initializes the databases for notes and tasks, and performs schema migration.

func Dinit() {
	var notesDBPath, tasksDBPath string

	switch runtime.GOOS {
	case "linux":
		notesDBPath = utils.ExpandHome("~/.local/state/toGo/notes.db")
		tasksDBPath = utils.ExpandHome("~/.local/state/toGo/tasks.db")
	case "darwin":
		notesDBPath = utils.ExpandHome("~/.local/state/toGo/notes.db")
		tasksDBPath = utils.ExpandHome("~/.local/state/toGo/tasks.db")
	case "windows":
		notesDBPath = filepath.Join(os.Getenv("USERPROFILE"), "go", "bin", "notes.db")
		tasksDBPath = filepath.Join(os.Getenv("USERPROFILE"), "go", "bin", "tasks.db")
	default:
		fmt.Println("Unknown operating system")
		return
	}

	if err := os.MkdirAll(filepath.Dir(notesDBPath), os.ModePerm); err != nil {
		utils.Fatal("Failed to create directory for notes database:", err)
	}
	if err := os.MkdirAll(filepath.Dir(tasksDBPath), os.ModePerm); err != nil {
		utils.Fatal("Failed to create directory for tasks database:", err)
	}
	DbNote, err = gorm.Open(sqlite.Open(notesDBPath), &gorm.Config{})
	if err != nil {
		utils.Fatal("Failed to connect to the notes database:", err)
	}
	err = DbNote.AutoMigrate(&Note{}) // Automatic schema migration for notes
	if err != nil {
		utils.Fatal("Failed to perform notes migration:", err)
	}

	// Initialize the database for tasks
	DbTask, err = gorm.Open(sqlite.Open(tasksDBPath), &gorm.Config{})
	if err != nil {
		utils.Fatal("Failed to connect to the tasks database:", err)
	}
	err = DbTask.AutoMigrate(&Task{}) // Automatic schema migration for tasks
	if err != nil {
		utils.Fatal("Failed to perform tasks migration:", err)
	}
}

// GetNotes returns all notes from the database.
func GetNotes() []Note {
	var notes []Note
	result := DbNote.Find(&notes)
	if result.Error != nil {
		utils.Fatal("Error while retrieving notes:", result.Error)
	}
	return notes
}

// GetTasks returns all tasks from the database.
func GetTasks() []Task {
	var tasks []Task
	result := DbTask.Find(&tasks)
	if result.Error != nil {
		utils.Fatal("Error while retrieving tasks:", result.Error)
	}
	return tasks
}

// AddNotes adds a new note to the database.
func AddNotes(message string) {
	some := Note{
		Message: message,
		Date:    time.Now().UTC().Format(format), // Set current date
	}
	result := DbNote.Create(&some)
	if result.Error != nil {
		utils.Fatal("Error while adding note:", result.Error)
	}
}

// DelNotes deletes a note by the specified ID.
func DelNotes(id int) {
	result := DbNote.Delete(&Note{}, id)
	if result.Error != nil {
		utils.Fatal("Error while deleting note:", result.Error)
	}
}

// AddTask adds a new task to the database.
func AddTask(message string) {
	some := Task{
		Message: message,
		Date:    time.Now().UTC().Format(format), // Set current date
		Flag:    false,                           // Task not completed yet
	}
	result := DbTask.Create(&some)
	if result.Error != nil {
		utils.Fatal("Error while adding task:", result.Error)
	}
}

// DelTask deletes a task by the specified ID.
func DelTask(id int) {
	result := DbTask.Delete(&Task{}, id)
	if result.Error != nil {
		utils.Fatal("Error while deleting task:", result.Error)
	}
}

// DoneTask changes the status of a task to completed by the specified ID.
func DoneTask(id int) {
	var task Task
	result := DbTask.First(&task, id)
	if result.Error != nil {
		utils.Fatal("Error while searching for task:", result.Error)
		return
	}
	task.Flag = true // Set the completed task flag
	result = DbTask.Save(&task)
	if result.Error != nil {
		utils.Fatal("Error while updating task:", result.Error)
	}
}

func RemessTask(id int, message string) {
	var task Task
	result := DbTask.First(&task, id)
	if result.Error != nil {
		utils.Fatal("Error while searching for task:", result.Error)
		return
	}
	task.Message = message
	result = DbTask.Save(&task)
	if result.Error != nil {
		utils.Fatal("Error while updating task:", result.Error)
	}
}

func RemessNote(id int, message string) {
	var note Note
	result := DbNote.First(&note, id)
	if result.Error != nil {
		utils.Fatal("Error while searching for task:", result.Error)
		return
	}
	note.Message = message
	result = DbNote.Save(&note)
	if result.Error != nil {
		utils.Fatal("Error while updating task:", result.Error)
	}
}
