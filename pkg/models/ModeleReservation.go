package models

type Reservation struct {
	HoraireStart  string
	HoraireEnd    string
	IdEtat        int64
	NomEtat       string
	IdReservation int64
	IdSalle       int64
	NomSalle      string
	PlaceSalle    int64
}
