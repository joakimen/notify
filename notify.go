package notify

import (
	"log"
	"os/exec"
)

// Notification holds the content of the notification
type Notification struct {
	Title   string
	Message string
}

// NewNotification returns a notification with the specified content
func NewNotification(title, message string) Notification {
	if title == "" {
		title = "notification"
	}
	return Notification{
		Title:   title,
		Message: message,
	}
}

// notifier is a displayable object
type notifier interface {
	display()
}

func (n Notification) display() {

	cmd := exec.Command(
		"terminal-notifier",
		"-title", n.Title,
		"-message", n.Message,
	)

	err := cmd.Run()

	if err != nil {
		log.Fatalln("An error occurred displaying notification:", err)
	}

}

// Display a simple notification using "terminal-notifier"
func Display(n notifier) {
	n.display()
}
