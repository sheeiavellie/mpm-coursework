package data

import "fmt"

type ProcessRequest struct {
	Timestamp string       `json:"timestamp"`
	Data      []SensorData `json:"data"`
}

type SensorData interface {
	GetSensorData() string
}

type Potentiometer struct {
	ID         int    `json:"id"`
	SensorType string `json:"sensor_type"`
	Value      int    `json:"value"`
}

func (pd *Potentiometer) GetSensorType() string {
	return fmt.Sprintf(
		"| id: %d | type: %s | value: %d |",
		pd.ID, pd.SensorType, pd.Value,
	)
}
