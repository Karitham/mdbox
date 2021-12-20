package main

func elems(s string, els []string) bool {
	for _, e := range els {
		if s == e {
			return true
		}
	}
	return false
}
