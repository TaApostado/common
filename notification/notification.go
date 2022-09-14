package notification

import (
	"fmt"
	"strings"
)

type Notification struct {
	context, message string
}

func NewNotification(context, message string) *Notification {
	return &Notification{context: context, message: message}
}

type INotificator interface {
	HasNotifications() bool
	AddNotification(Notification)
	Notifications() string
}

type Notificator struct {
	notifications []Notification
}

func NewNotificator() *Notificator {
	return &Notificator{notifications: []Notification{}}
}

func (n *Notificator) HasNotifications() bool {
	return len(n.notifications) > 0
}

func (n *Notificator) AddNotification(notification Notification) {
	if n.notifications == nil {
		n.notifications = []Notification{}
	}
	n.notifications = append(n.notifications, notification)
}

func (n *Notificator) Notifications() string {
	var sb strings.Builder
	var notification Notification
	index := 0

	if !n.HasNotifications() {
		return ""
	}

	if len(n.notifications) == 1 {
		notification = n.notifications[index]
		sb.WriteString(
			fmt.Sprintf("%s: %s", notification.context, notification.message))
		return sb.String()
	}

	for index < len(n.notifications)-1 {
		notification = n.notifications[index]
		sb.WriteString(
			fmt.Sprintf("%s: %s, ", notification.context, notification.message))
		index++
	}

	fmt.Println(index)
	notification = n.notifications[index]
	sb.WriteString(
		fmt.Sprintf("%s: %s", notification.context, notification.message))
	return sb.String()
}
