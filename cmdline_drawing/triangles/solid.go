package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
func DrawSolidTriangle(length int) {
	for row := 0; row < length; row++ {
		for col := 0; col <= row; col++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

// HINWEIS
// - Diese Aufgabe ähnelt der zum Zeichnen eines gefüllten Rechtecks.
// - Nutzen Sie - wie dort - zwei geschatelte Schleifen, um die Zeilen und Spalten des Dreiecks zu durchlaufen.
// - Die innere Schleife darf aber dieses Mal nicht in jedem Schritt bis zum Ende der Zeile laufen.
