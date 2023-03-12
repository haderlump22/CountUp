package main

import (
	"fmt"
	"time"
)

const ClearLine = "\033[2K\r"

func main() {
	// Fastenzeitstart in 2023
	start := time.Date(2023, 02, 22, 0, 0, 0, 0, time.Local)

	// berechne Zeit seit Begin der Fastenzeit ohne Süsigkeiten
	for {
		aktuell := time.Now()
		Abstand := aktuell.Sub(start)
		TageSeit := uint(Abstand.Hours() / 24)
		StundenSeit := uint(uint(Abstand.Hours()) - (TageSeit * 24))
		MinutenSeit := uint(Abstand.Minutes()) - (TageSeit * 1440) - (StundenSeit * 60)
		SekundenSeit := uint(Abstand.Seconds()) - (TageSeit * 86400) - (StundenSeit * 60 * 60) - (MinutenSeit * 60)
		fmt.Printf("Tage=%v, Stunden=%v, Minuten=%v, Sekunden=%v", TageSeit, StundenSeit, MinutenSeit, SekundenSeit)
		time.Sleep(1 * time.Second)
		fmt.Print(ClearLine)
	}

}
