package models

type PushNotification struct {
	UserId  string
	Type    string
	Title   string
	Message string
	IsRead  bool
}
