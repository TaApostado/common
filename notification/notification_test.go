package notification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewNotification(t *testing.T) {
	notificator := NewNotificator()
	assert.Equal(
		t, len(notificator.notifications),
		0, "new notificator should has 0 notification`s")

	notificator.AddNotification(*NewNotification("context", "some error..."))
	assert.Equal(
		t, len(notificator.notifications),
		1, "after add one notification notificator should has 1 notification`s")

	notificator.AddNotification(*NewNotification("context", "some error..."))
	assert.Equal(
		t, len(notificator.notifications),
		2, "after add one notification notificator should has 2 notification`s")
}

func TestHasNotifications(t *testing.T) {
	notificator := NewNotificator()

	assert.Equal(
		t, false,
		notificator.HasNotifications(), "should return false because notificator hasn`t notification`s")

	notificator.notifications = append(
		notificator.notifications, *NewNotification("context", "some error..."))
	assert.Equal(
		t, true,
		notificator.HasNotifications(), "should return true because notificator has 1 notification`s")

	notificator.notifications = append(
		notificator.notifications, *NewNotification("other context", "some error..."))
	assert.Equal(
		t, true,
		notificator.HasNotifications(), "should return true because notificator has 2 notification`s")
}

func TestNotifications(t *testing.T) {
	notificator := NewNotificator()

	assert.Equal(
		t, "",
		notificator.Notifications(), "notificator message malformed")

	notificator.notifications = append(
		notificator.notifications, *NewNotification("context", "some error..."))
	assert.Equal(
		t, "context: some error...",
		notificator.Notifications(), "notificator message malformed")

	notificator.notifications = append(
		notificator.notifications, *NewNotification("other context", "some error..."))
	assert.Equal(
		t, "context: some error..., other context: some error...",
		notificator.Notifications(), "notificator message malformed")
}
