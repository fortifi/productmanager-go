package pmlog

import (
	"github.com/fortifi/productmanager-go/pmtime"
	"time"
)

type Type string

const (
	TypeDebug   Type = "debug"
	TypeInfo    Type = "info"
	TypeWarning Type = "warning"
	TypeError   Type = "error"
	TypeSuccess Type = "success"
)

type Message struct {
	MessageType Type   `json:"type"`
	Message     string `json:"message"`
	Timestamp   int64  `json:"timestamp"`
}

func NewMessage(messageType Type, message string, time time.Time) Message {
	mtime := pmtime.FromTime(time)
	return Message{
		MessageType: messageType,
		Message:     message,
		Timestamp:   mtime.ForTransport(),
	}
}

func Debug(message string) Message   { return NewMessage(TypeDebug, message, time.Now()) }
func Info(message string) Message    { return NewMessage(TypeInfo, message, time.Now()) }
func Warning(message string) Message { return NewMessage(TypeWarning, message, time.Now()) }
func Error(message string) Message   { return NewMessage(TypeError, message, time.Now()) }
func Success(message string) Message { return NewMessage(TypeSuccess, message, time.Now()) }
