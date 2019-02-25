package notify_test

import (
	"kajito/notify"
	"testing"
)

func TestNotification_NewNotification(t *testing.T) {

	var tests = []struct {
		name         string
		in_title     string
		in_message   string
		want_title   string
		want_message string
	}{
		{
			name:         "attributes are set",
			in_title:     "the title",
			in_message:   "the message",
			want_title:   "the title",
			want_message: "the message",
		},
		{
			name:         "no title defaults to 'notification'",
			in_title:     "",
			in_message:   "the message",
			want_title:   "notification",
			want_message: "the message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := notify.NewNotification(tt.in_title, tt.in_message)
			got_title, got_message := n.Title, n.Message
			if got_title != tt.want_title {
				t.Errorf("got %v, want %v", got_title, tt.want_title)
			}

			if got_message != tt.want_message {
				t.Errorf("got %v, want %v", got_message, tt.want_message)
			}
		})
	}
}
