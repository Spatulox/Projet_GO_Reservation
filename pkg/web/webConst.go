package web

import (
	"html/template"
)

const (
	PORT       = "8085"
	RouteIndex = "/"

	RouteIndexReservation  = "/reservation"
	RouteListReservation   = "/reservation/list"
	RouteCreateReservation = "/reservation/create"

	RouteUpdateReservation = "/reservation/update"
	RouteCancelReservation = "/reservation/cancel"
)

// var templates = template.Must(template.ParseGlob("pkg/web/html/*.html"))
var templates = template.Must(template.ParseGlob("templates/*.html"))
