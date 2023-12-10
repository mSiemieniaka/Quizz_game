package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// flaga dla uzytkownika przy wpisaniu -h w cmd dostaje notke
	plikcsv := flag.String("csvfile", "problem.csv", "Plik przechowujacy pytania do quizu :)")
	flag.Parse()
	// obsluga pliku jak nie otwiera wazne aby plik na koncu zamykac
	plik, err := os.Open(*plikcsv)
	if err != nil {
		log.Fatalf("Nie mozna otworzyc pliku: %s", err)
	}
	defer plik.Close()
	// kolejny raz obsluga pliku gdy nie bedzie znany format
	odczytCsv := csv.NewReader(plik)
	data, err := odczytCsv.ReadAll()
	if err != nil {
		log.Fatalf("Error wyswietlajac plik: %s", err)

	}
	// zmienna pomocnicza do liczenia wyniku
	poprawne_odp := 0
	//petla odpowidajaca za wyswietlanie i odpowidanie na pytania
	for _, wartosc := range data {
		ask := wartosc[0]
		odp := wartosc[1]
		fmt.Println("Pytanie:", ask, "?")
		var wejscie string
		fmt.Scanln(&wejscie)
		if wejscie == odp {

			fmt.Println("Congrats dobra odpowiedz :)")
			poprawne_odp++
		} else {
			fmt.Println("Niestety bledna odpowiedz :(")

		}

	}

	fmt.Printf("Odpowiedziales poprawnie na %d/%d pytan\n", poprawne_odp, len(data))

}
