package config

import (
	"os"

	"github.com/pusher/pusher-http-go/v5"
)

type Pusher interface {
	MessageTrigger(eventName string, data any) error
	OrderTrigger(eventName string, data any) error
	NotificationTrigger(eventName string, data any) error
}

type PusherImpl struct {
	pusher.Client
}

func (pusher *PusherImpl) MessageTrigger(eventName string, data any) error {
	return pusher.Client.Trigger("message_channel", eventName, data)
}

func (pusher *PusherImpl) OrderTrigger(eventName string, data any) error {
	return pusher.Client.Trigger("order_channel", eventName, data)
}

func (pusher *PusherImpl) NotificationTrigger(eventName string, data any) error {
	return pusher.Client.Trigger("notification_channel", eventName, data)
}

func NewPusher() Pusher {
	return &PusherImpl{Client: pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}}
}
