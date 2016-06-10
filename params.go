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

func GuessParams(c Commands) *Params {
	goroutines := make(map[int]int) // map[depth]number
	var totalG int

	for _, cmd := range c.cmds {
		depth := c.gd[cmd.Name]
		if cmd.Command == CmdCreate {
			totalG++
			goroutines[depth]++
		}
	}

	params := &Params{
		Angle:    360.0 / float64(goroutines[1]),
		Caps:     totalG < 5, // value from head
		Distance: 80,
	}

	if gs, ok := goroutines[2]; ok {
		params.AngleSecond = 360.0 / float64(gs/goroutines[1])
		params.DistanceSecond = 20
	}

	return params
}
