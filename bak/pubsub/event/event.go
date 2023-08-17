package event

import "github.com/google/uuid"

type Event struct {
	UserIds []string
	Payload any
}

type ApplyAddContactEvent struct {
	Event     string
	Id        string
	UserId    string
	UserName  string
	ContactId string
	Message   string
}

func (event *ApplyAddContactEvent) SetDefaultValue() *ApplyAddContactEvent {
	event.Event = "APPLY_ADD_CONTACT_EVENT"
	event.Id = uuid.New().String()
	return event
}

// type MessageEvent struct {
// 	Event     string
// 	Id        string
// 	Type      string
// 	Content   string
// 	CreaterId string
// 	ChatId    string
// }
// func (event *MessageEvent) SetDefaultValue() *MessageEvent {
// 	event.Event = "MESSAGE_EVENT"
// 	event.Id = uuid.New().String()
// 	return event
// }

type MessageEvent struct {
	Event                    string
	Id                       string
	MessageId                string
	CreateTime              string
	Type                     string
	Content                  string
	CreaterId                string
	CreaterContactRemarkName string
	ChatId                   string
}

func (event *MessageEvent) SetDefaultValue() *MessageEvent {
	event.Event = "MESSAGE_EVENT"
	event.Id = uuid.New().String()
	return event
}
