package forms

// Erwartet eine Seitenlänge a.
// Liefert den Umfang des entsprechenden Quadrats.
func PerimeterSquare(a float64) float64 {
	return PerimeterRectangle(a, a)
}

// HINWEIS
// Benutzen Sie die Funktion PerimeterRectangle().
