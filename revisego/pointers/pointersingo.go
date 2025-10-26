package pointers

func FetchAddressOfInteger(a *int) (*int, int) {
	return a, *a
}