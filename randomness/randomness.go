// Randomness - General Purpose Random Number Routines
// KEC - 06/2020

package randomness

import (
	"math/rand"
	"time"
)

// GetRandom - Returns a random number from 1 to rollamt.
func GetRandom(rollamt int) int {
	// General purpose, reusable, integer random number generator.
	// rand.Intn(int) returns a value between 0 and n, so add 1 to return.
	var randnum int
	rand.Seed(time.Now().UnixNano()) // Use Unix-formatted time for random seed.

	randnum = rand.Intn(rollamt) // Get random number (integer).
	randnum = randnum + 1        // Add "1" to random number, since it starts at "0".

	return randnum // Here's your random number. Enjoy that!
}
