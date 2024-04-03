package data

import "fmt"

type ProcessRequest struct {
	Timestamp string              `json:"timestamp"`
	Data      []PotentiometerData `json:"data"`
}

type SensorData interface {
	GetSensorData() string
}

type PotentiometerData struct {
	ID         int    `json:"id"`
	SensorType string `json:"sensor_type"`
	Value      int    `json:"value"`
}

func (pd *PotentiometerData) GetSensorData() string {
	return fmt.Sprintf(
		"| id: %d | type: %s | value: %d |",
		pd.ID, pd.SensorType, pd.Value,
	)
}
