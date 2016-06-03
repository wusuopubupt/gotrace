package main

// Params represents params for GoThree.js library.
type Params struct {
	Angle int  `json:"angle"`
	Caps  bool `json:"allCaps"`
}

func GuessParams(c Commands) *Params {
	return &Params{
		Angle: 120,
		Caps:  true,
	}
}
