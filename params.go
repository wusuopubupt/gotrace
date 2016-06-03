package main

// Params represents params for GoThree.js library.
type Params struct {
	Angle          int  `json:"angle"`
	AngleSecond    int  `json:"angle2"`
	Caps           bool `json:"allCaps"`
	Distance       int  `json:"distance"`
	DistanceSecond int  `json:"distance2"`
}

func GuessParams(c Commands) *Params {
	return &Params{
		Angle:          -15,
		AngleSecond:    360 / 3,
		Caps:           true,
		Distance:       80,
		DistanceSecond: 20,
	}
}
