package config

import (
	"github.com/pusher/pusher-http-go/v5"
)

var pusherClient = pusher.Client{
	AppID:   "2025350",
	Key:     "e31a8b7ecfbaf1164828",
	Secret:  "42b4da299d97486eacb4",
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
