package main

// Params represents params for GoThree.js library.
type Params struct {
	Angle int `json:"angle"`
}

func GuessParams(c Commands) *Params {
	return &Params{
		Angle: 5,
	}
}
