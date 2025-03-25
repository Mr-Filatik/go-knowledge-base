package printer

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintSliceCapacityIncrease(capacities []int) {
	fmt.Printf("----- Slice capacity increase -----\n")

	if len(capacities) <= 1 {
		fmt.Println("There must be more than 1 element in a slice.")
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "================\t===============\t============")
	fmt.Fprintln(w, "Current Capacity\tIncrease Factor\tNew Capacity")
	fmt.Fprintln(w, "----------------\t---------------\t------------")

	lastCapacity := capacities[0]
	increaseFactor := float64(0)
	for i := 1; i < len(capacities); i++ {
		currentCapacity := capacities[i]
		increaseFactor = float64(currentCapacity) / float64(lastCapacity)
		fmt.Fprintf(w, "%d\t%.2f\t%d\n",
			lastCapacity,
			increaseFactor,
			currentCapacity,
		)
		lastCapacity = currentCapacity
	}
	fmt.Fprintln(w, "================\t===============\t============")
	w.Flush()
}
