package itertools

// Product is the cartesian product of input iterables.
//
// Roughly equivalent to nested for-loops in a generator expression.
// For example, product(A, B) returns the same as ((x,y) for x in A for y in B).
//
// The nested loops cycle like an odometer with the rightmost element advancing on every iteration.
// This pattern creates a lexicographic ordering so that if the input’s iterables are sorted, the product tuples are emitted in sorted order.
//
// To compute the product of an iterable with itself, specify the number of repetitions with the optional repeat keyword argument.
// For example, product(4, A) means the same as product(A, A, A, A).
func Product(repeat uint64, args ...[]interface{}) [][]interface{} {
	// product('ABCD', 'xy') --> Ax Ay Bx By Cx Cy Dx Dy
	// product(range(2), repeat=3) --> 000 001 010 011 100 101 110 111
	var pools [][]interface{}
	for _, pool := range args {
		for i := uint64(0); i < repeat; i++ {
			pools = append(pools, pool)
		}
	}

	result := [][]interface{}{{}}
	for _, pool := range pools {
		var newResult [][]interface{}
		for _, x := range result {
			for _, y := range pool {
				newResult = append(newResult, append(x, y))
			}
		}
		result = newResult
	}

	return result
}
