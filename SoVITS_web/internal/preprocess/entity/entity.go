package entity

type SliceRequest struct {
	InputPath   string
	OutputRoot  string
	Threshold   float32
	MinLength   int32
	MinInterval int32
	HopSize     int32
	MaxSilKept  int32
	XMax        int32
	Alpha       float32
	NParts      int32
}

type SliceResponse struct {
	Status string
	Data   map[string]string
}
