package scanner

func ScanSliceCapacityIncrease(startCap uint, limitCap int) []int {
	slice := make([]int, startCap)

	capacities := make([]int, 0)
	capacities = append(capacities, cap(slice))

	count := cap(slice) - len(slice)
	for cap(slice)+count+1 < limitCap {
		slice = append(slice, make([]int, count+1)...)
		count = cap(slice) - len(slice)
		capacities = append(capacities, cap(slice))
	}

	return capacities
}
