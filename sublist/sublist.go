package sublist

// Sublist checks relations between two arrays
// Might return: equal, superlist, sublist, unequal
func Sublist(listOne []int, listTwo []int) string {

	switch {
	case len(listOne) == len(listTwo) && isSublist(listOne, listTwo):
		return "equal"

	case len(listOne) > len(listTwo) && isSublist(listTwo, listOne):
		return "superlist"

	case len(listOne) < len(listTwo) && isSublist(listOne, listTwo):
		return "sublist"

	default:
		return "unequal"
	}
}

func isSublist(listOne []int, listTwo []int) bool {
	var j, k = 0, 0

	if len(listOne) == 0 || len(listTwo) == 0 {
		return true
	}

	for j < len(listOne) && k+j < len(listTwo) {

		if listOne[j] == listTwo[k+j] {
			j++

			if j == len(listOne) {
				return true
			}
		} else {
			j = 0
			k++
		}
	}

	return false
}
