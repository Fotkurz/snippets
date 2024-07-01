package kangaroo

import "fmt"

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	var maxDistance int32 = 10000
	isPossible := false

	for {
		if x1 == x2 {
			isPossible = true
			break
		}
		if x1 > maxDistance || x2 > maxDistance {
			break
		}

		// do jump
		x1 = x1 + v1
		x2 = x2 + v2
		fmt.Printf("1->%d,\n2->%d\n", x1, x2)
	}

	if isPossible {
		return "YES"
	}
	return "NO"
}

func KangarooV2(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// x1 + (y * v1) = x2 + (y * v2)
	// (y * v1) = (x2 - x1) + (y * v2)
	// (y * v1) - (y * v2) = (x2 - x1)
	// (x2 - x1)/(v1 - v2) = y
	if v1 > v2 && (x2-x1)%(v1-v2) == 0 {
		return "YES"
	}

	return "NO"
}
