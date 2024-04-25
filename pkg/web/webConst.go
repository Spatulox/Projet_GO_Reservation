package web

import (
	"html/template"
)

const (
	// Global things
	PORT       = "8085"
	RouteIndex = "/"

	// Route for reservation
	RouteIndexReservation  = "/reservation"
	RouteListReservation   = "/reservation/list"
	RouteCreateReservation = "/reservation/create"
	RouteUpdateReservation = "/reservation/update"
	RouteCancelReservation = "/reservation/cancel"

	// Route for Rooms
	RouteGetAllRoolAvailable = "/salle/getAllAvail"
	RouteGetAllRooms         = "/salle/getRoomAll"
	RouteCreateRoom          = "/salle/create"
	RouteCancelSalle         = "/salle/cancel"

	// Route for JSON
	RouteDownloadJson = "/download"
	RouteExportJson   = "/reservation/export"
	//RouteImportJson = "/reservation/import"
)

//old templates
// var templates = template.Must(template.ParseGlob("pkg/web/html/*.html"))

// New templates
var templates = template.Must(template.ParseGlob("templates/*.html"))
