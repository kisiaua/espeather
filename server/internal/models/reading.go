package models

type Reading struct {
	Temperature float32
	Humidity    float32
}

type ReadingFull struct {
	ID          int
	Temperature float32
	Humidity    float32
	Timestamp   string
}
