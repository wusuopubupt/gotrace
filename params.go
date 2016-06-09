package main

// Params represents params for GoThree.js library.
type Params struct {
	Angle          float64 `json:"angle"`
	AngleSecond    float64 `json:"angle2"`
	Caps           bool    `json:"allCaps"`
	Distance       int     `json:"distance"`
	DistanceSecond int     `json:"distance2"`
	AutoAngle      bool    `json:"autoAngle"`
}

func GuessParams(cmds Commands) *Params {
	var (
		goroutines int
	)
	for _, cmd := range cmds {
		if cmd.Command == CmdCreate {
			goroutines++
		}
	}

	return &Params{
		Angle:          360.0 / float64(goroutines-1),
		AngleSecond:    360.0 / float64(goroutines-1),
		Caps:           goroutines < 5, // value from head
		Distance:       80,
		DistanceSecond: 20,
		AutoAngle:      false,
	}
}
