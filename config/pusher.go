package config

import (
	"os"

	"github.com/pusher/pusher-http-go/v5"
)

var pusherClient = pusher.Client{
	AppID:   os.Getenv("PUSHER_APP_ID"),
	Key:     os.Getenv("PUSHER_KEY"),
	Secret:  os.Getenv("PUSHER_SECRET"),
	Cluster: os.Getenv("PUSHER_CLUSTER"),
	Secure:  true,
}

func MessageTrigger(eventName string, data any) error {
	return pusherClient.Trigger("message_channel", eventName, data)
}

func OrderTrigger(eventName string, data any) error {
	return pusherClient.Trigger("order_channel", eventName, data)
}

func NotificationTrigger(eventName string, data any) error {
	return pusherClient.Trigger("notification_channel", eventName, data)
}
