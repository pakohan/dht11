package dht

import "testing"

func TestGetSensorData(t *testing.T) {
	hum, temp, err := GetSensorData(SensorDHT11, 1)
	if err != nil {
		t.Fatalf("Failed to read sensor data: %s", err)
	}
	t.Logf("%.2f%Humidity, %.2f°C", hum, temp)
}
