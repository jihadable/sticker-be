package config

import "github.com/pusher/pusher-http-go/v5"

var pusherClient = pusher.Client{
	AppID:   "2024160",
	Key:     "8df6d4d768d779f1f8b3",
	Secret:  "33631569a1cdd682a4b4",
	Cluster: "ap1",
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
