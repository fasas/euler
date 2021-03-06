// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

// Problem1 is solution for finding the sum of all the multiples of 3 or 5 below 1000.
func Problem1() int {
	return sn(3, 999) + sn(5, 999) - sn(15, 999)
}

// sn returns the sum of arithmetic sequence from 0 up to the an with difference d.
func sn(d, an int) int {
	n := an / d
	return int(float64(n) / 2 * float64((2*d)+(n-1)*d))
}
