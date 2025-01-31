package make_unique

import "fmt"

// Erwartet eine Liste von Strings.
// Hängt Zahlen an alle mehrfach vorkommenden Strings an, um sie eindeutig zu machen.
func MakeUnique(strings []string) {
	for i := 0; i < len(strings); i++ {
		count := 1
		for j := i + 1; j < len(strings); j++ {
			if strings[i] == strings[j] {
				strings[j] = fmt.Sprintf("%v_%v", strings[i], count+1)
				count++

			}
		}
	}
}

// REMARKS
// - Dies ist eine ähnliche Aufgabe wie das Zählen von Vorkommen in einer Liste.
//   Die innere Schleife macht i.W. das gleiche wie die Funktion `Count` aus der Aufgabe `count`.
//   Zusätzlich wird der Wert von `count` an den String angehängt, um ihn eindeutig zu machen.
