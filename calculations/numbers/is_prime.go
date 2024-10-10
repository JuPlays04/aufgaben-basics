package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	return CountDivisors(n) == 2
}

// HINWEIS
// - Eine Primzahl ist eine Zahl, die genau zwei Teiler hat.
// - Sie können den Code von `CountDivisors` wiederverwenden,
//   oder Sie verwenden direkt die Funktion `CountDivisors` in dieser Funktion.
