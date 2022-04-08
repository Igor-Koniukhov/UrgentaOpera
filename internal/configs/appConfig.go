package configs

import (
	"github.com/subosito/gotenv"
	"log"
	"time"
)

type AppConfig struct {
	MessageChanDB    chan []byte
	MessageChanWS    chan []byte
	MessagePointToDB chan []byte
	Data             map[string]interface{}
	InfoString       map[string]string
	Int              map[string]int
	Float            map[string]float64
	TimeFormat       string
}

func NewAppConfig() *AppConfig {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return &AppConfig{
		MessageChanDB:    make(chan []byte, 1024),
		MessageChanWS:    make(chan []byte, 1024),
		MessagePointToDB: make(chan []byte, 1024),
		Data:             make(map[string]interface{}),
		InfoString:       make(map[string]string),
		Int:              make(map[string]int),
		Float:            make(map[string]float64),
		TimeFormat:       time.Now().UTC().Format("2006-01-02 15:04:05.999999"),
	}
}
