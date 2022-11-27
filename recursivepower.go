package piscine

func RecursivePower(nb int, power int) (pow int) {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	pow = nb * RecursivePower(nb, power-1)

	return pow
}
