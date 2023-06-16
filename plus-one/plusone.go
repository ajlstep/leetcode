package plusone

func plusOne(digits []int) []int {
	oneinmid := 0

	for i := len(digits) - 1; i >= 0; i-- {
		var rez int
		if i == len(digits)-1 {
			rez = digits[i] + 1 + oneinmid
		} else {
			rez = digits[i] + oneinmid
		}
		if oneinmid != 0 {
			oneinmid -= 1
		}

		if rez > 9 {
			digits[i] = 0
			oneinmid += 1
		} else {
			digits[i] = rez
		}
	}
	if oneinmid != 0 {
		digits = append([]int{oneinmid}, digits...)
	}
	return digits
}
