package almostpefect

import (
	"fmt"
	"math"
)



func AlmostPerfect(n int) {
	sq := int(math.Sqrt(float64(n))+1)
	sum := 1

	for i:=2; i<sq;i++ {
		if n%1 == 0 {
			other := n/i
			if i != other {
				sum += other
			}
			sum +=  i
		}
	}

	switch {
	case sum == n:
		fmt.Printf("%d perfect", n)
	case math.Abs(float64(sum - n)) <= 2:
		fmt.Printf("%d almost perfect", n)
	default:
		fmt.Printf("%d not perfect", n)
	}
}