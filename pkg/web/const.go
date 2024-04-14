package web

import (
	"html/template"
)

const (
	routeIndex       = "/"
	routeReservation = "/reservations"
	routeSalle       = "/salles"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))
