package models

type Telemetry struct {
	Timestamp     int64         `json:"timestamp,omitempty" bson:"timestamp"`
	UserID        string        `json:"user_id,omitempty" bson:"user_id"`
	GPS           GPS           `json:"gps,omitempty" bson:"gps"`
	Accelerometer Accelerometer `json:"accelerometer,omitempty" bson:"accelerometer"`
}

type GPS struct {
	Latitude            float64 `json:"latitude,omitempty" bson:"latitude"`
	Longitude           float64 `json:"longitude,omitempty" bson:"longitude"`
	Speed               float64 `json:"speed,omitempty" bson:"speed"`
	Direction           string  `json:"direction,omitempty" bson:"direction"`
	ConnectedSatellites int64   `json:"connected_satellites,omitempty" bson:"connected_satellites"`
}

type Accelerometer struct {
	AngleX float64 `json:"angle_x,omitempty" bson:"angle_x"`
	AngleY float64 `json:"angle_y,omitempty" bson:"angle_y"`
	AngleZ float64 `json:"angle_z,omitempty" bson:"angle_z"`
	GForce float64 `json:"g_force,omitempty" bson:"g_force"`
}
