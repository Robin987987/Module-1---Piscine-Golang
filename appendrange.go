package piscine

/*
Write a function that takes an int min and an int max as parameters.
The function should return a slice of ints with all the values between the min and max.
Min is included, and max is excluded.
If min is greater than or equal to max, a nil slice is returned.
*/

func AppendRange(min, max int) []int {
	var array []int
	if max > min {
		for i := min; i < max; i++ { // i<=max loaks kuni 10ni aga ainult < ei loe viimast numbrit sisse
			array = append(array, i)
		}
	}
	return array
}
