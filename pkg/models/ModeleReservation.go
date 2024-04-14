package models

type Reservation struct {
	HoraireStart  string
	HoraireEnd    string
	IdEtat        int64
	IdReservation int64
	NomSalle      string
	PlaceSalle    int64
}
