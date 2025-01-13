package Numbers

// Name: Purva Zinjarde

func Solve(width, height, length, mass int) string {

	// used boolean values instead of writing a separate function for checking if the
	// package is bulky, as with this, it can be computed just once and used multiple times.
	isBulky, isHeavy := false, false

	// check if mass is over the constraint.
	// it is okay not to declare a boolean here, and directly compare, but I have introduced
	// a boolean for better readability.
	if mass >= 20 {
		isHeavy = true
	}

	// check if volume is more than allowed.
	volume := width * height * length
	if width >= 150 || height >= 150 || length >= 150 || volume >= 1000000 {
		isBulky = true
	}

	// return according to the constraints.
	if !isBulky && !isHeavy {
		return "STANDARD"
	} else if isBulky && isHeavy {
		return "REJECTED"
	} else {
		return "SPECIAL"
	}
}
