package main

func fibonacci() func() int {
	prev := 0
	current := 1
	count := 0

	return func() int {
		if count == 0 {
			count += 1
			return prev
		}
		if count == 1 {
			count += 1
			return current
		}

		result := prev + current
		prev = current
		current = result

		return result
	}
}
