package main

import (
    "fmt"
    "strings"
    "time"
)

func main() {
    var geboortedatumStr, huidigeDatumStr string

    // Gebruikersinput voor de geboortedatum van de kat
    fmt.Print("Voer de geboortedatum van de kat in (dd/mm/yyyy): ")
    fmt.Scanln(&geboortedatumStr)

    // Gebruikersinput voor de huidige datum
    fmt.Print("Voer de huidige datum in (dd/mm/yyyy): ")
    fmt.Scanln(&huidigeDatumStr)

    // Parsing de data
    layout := "02/01/2006" // Go's layout voor parsing moet de referentiedatum 02/01/2006 zijn
    geboortedatumKat, err := time.Parse(layout, geboortedatumStr)
    if err != nil {
        fmt.Println("Fout bij het parsen van de geboortedatum:", err)
        return
    }

    huidigeDatum, err := time.Parse(layout, huidigeDatumStr)
    if err != nil {
        fmt.Println("Fout bij het parsen van de huidige datum:", err)
        return
    }

    // Bereken de leeftijd van de kat in dagen
    leeftijdKatDagen := huidigeDatum.Sub(geboortedatumKat).Hours() / 24

    // Bereken de leeftijd van de kat in mensenjaren
    var leeftijdMensenjaren float64
    if leeftijdKatDagen < 365*2 {
        leeftijdMensenjaren = leeftijdKatDagen / 365 * 12
    } else {
        leeftijdMensenjaren = 24 + ((leeftijdKatDagen - 365*2) / 365 * 4)
    }

    fmt.Printf("Leeftijd van de kat in mensenjaren: %.2f\n", leeftijdMensenjaren)
}
