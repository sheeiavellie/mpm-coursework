package data

import "time"

type ProcessRequest struct {
	Timestamp time.Time    `json:"timestamp"`
	Data      []SensorData `json:"data"`
}

type SensorData interface {
	GetSensorType() string
}

type Potentiometer struct {
	sensorType string
	ID         int `json:"id"`
	Value      int `json:"value"`
}

func (pd *Potentiometer) GetSensorType() string {
	return pd.sensorType
}
