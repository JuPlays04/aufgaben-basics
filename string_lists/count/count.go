package count

// Erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Liefer die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string) int {
	count := 0
	for i := 0; i < len(strings); i++ {
		if strings[i] == search {
			count++
		}
	}
	return count
}
func ContainsDuplicate(strings []string) bool {
	seen := make(map[string]bool)
	for _, str := range strings {
		if seen[str] {
			return true
		}
		seen[str] = true

	}
	return false

}
