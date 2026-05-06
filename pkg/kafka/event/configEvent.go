package event

import (
	"time"
)

type ConfigVersionActivatedEvent struct {
	ServiceID     string `json:"service_id"`
	EventType     string `json:"event_type"`
	ConfigID      string `json:"config_id"`
	ConfigType    string `json:"config_type"`
	VersionNumber string `json:"version_number"`
	Timestamp     int64  `json:"timestamp"`
}

func NewConfigVersionActivatedEvent(serviceID, configID, configType, versionNumber string) *ConfigVersionActivatedEvent {
	return &ConfigVersionActivatedEvent{
		ServiceID:     serviceID,
		EventType:     "config.active",
		ConfigID:      configID,
		ConfigType:    configType,
		Timestamp:     time.Now().Unix(),
		VersionNumber: versionNumber,
	}
}
