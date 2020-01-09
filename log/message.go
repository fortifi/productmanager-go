package log

import (
	"github.com/fortifi/productmanager-go/pmtime"
	"time"
)

type Type string

const (
	TYPE_DEBUG   Type = "debug"
	TYPE_INFO    Type = "info"
	TYPE_WARNING Type = "warning"
	TYPE_ERROR   Type = "error"
	TYPE_SUCCESS Type = "success"
)

type Message struct {
	MessageType Type   `json:"type"`
	Message     string `json:"Message"`
	Timestamp   int64  `json:"Timestamp"`
}

func NewMessage(messageType Type, message string, time time.Time) Message {
	return Message{
		MessageType: messageType,
		Message:     message,
		Timestamp:   pmtime.FromTime(time).ForTransport(),
	}
}

func Debug(message string) Message   { return NewMessage(TYPE_DEBUG, message, time.Now()) }
func Info(message string) Message    { return NewMessage(TYPE_INFO, message, time.Now()) }
func Warning(message string) Message { return NewMessage(TYPE_WARNING, message, time.Now()) }
func Error(message string) Message   { return NewMessage(TYPE_ERROR, message, time.Now()) }
func Success(message string) Message { return NewMessage(TYPE_SUCCESS, message, time.Now()) }
