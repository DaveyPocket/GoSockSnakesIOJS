// Game code goes here

package logic

type points struct {
	x, y		int		// Coordinates
}

type s struct {
	body		[]points	// Slice of points
	dir			byte		// Head direction
	//owner					// Connected user
}

type game struct {
	snake		[]s		// Slice of snakes
	food		[]f		// Slice of foods
	state		byte	// State byte
	nowTick		byte	// Current tick byte
}
