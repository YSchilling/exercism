package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("")
	}

	current_number := 2
	for {
		if is_prime(current_number) {
			n--
			if n == 0 {
				return current_number, nil
			}
		}
		current_number++
	}
}

func is_prime(n int) bool {
	for i := 2; i < n ; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}