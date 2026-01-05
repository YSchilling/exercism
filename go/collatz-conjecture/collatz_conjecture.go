package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("")
	}

	var steps = 0
	for n != 1 {
		steps += 1
		if n % 2 == 0 {
			n /= 2
		} else {
			n = n * 3 + 1
		}
	}
	return steps, nil
}
