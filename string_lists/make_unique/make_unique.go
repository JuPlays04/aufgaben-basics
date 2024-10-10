package make_unique

import "fmt"

// Erwartet eine Liste von Strings.
// Hängt Zahlen an alle mehrfach vorkommenden Strings an, um sie eindeutig zu machen.
func MakeUnique(strings []string) {
	for i, s := range strings {
		count := 2
		for j := i + 1; j < len(strings); j++ {
			if strings[j] == s {
				strings[j] = fmt.Sprintf("%s_%d", s, count)
				count++
			}
		}
	}
}

// HINWEIS
// - Nutzen Sie eine Schleife, um durch die Liste der Strings zu iterieren.
// - Suchen Sie für jedes Element ab dessen Position nach weiteren Vorkommen.
//   Zählen Sie dabei die Anzahl der Vorkommen in einer Variablen (z.B. `count`).
//   Jedes Mal, wenn Sie ein weiteres Vorkommen finden, hängen Sie den aktuellen
//   Wert von `count` an den String an und erhöhen ihn danach um 1.
// REMARKS
// - Dies ist eine ähnliche Aufgabe wie das Zählen von Vorkommen in einer Liste.
//   Die innere Schleife macht i.W. das gleiche wie die Funktion `Count` aus der Aufgabe `count`.
//   Zusätzlich wird der Wert von `count` an den String angehängt, um ihn eindeutig zu machen.
// TAGS: [Arrays/Slices, Loops, Sprintf]
