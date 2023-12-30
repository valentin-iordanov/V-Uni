package main

import (
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Holiday struct {
	id         int
	title      string
	startDate  time.Time
	duration   time.Time
	price      float64
	freeSlots  int
	locationID int
}

type Location struct {
	id      int
	street  string
	number  int
	city    string
	country string
}

type Reservation struct {
	id          int
	contactName string
	phoneNumber string
	holidayID   int
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/travel")

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	route := mux.NewRouter()

	//holidays
	route.Methods(http.MethodGet).Path("/holidays").HandlerFunc(GetHolidays)
	route.Methods(http.MethodGet).Path("/holidays/{id}").HandlerFunc(GetHoliday)
	route.Methods(http.MethodPost).Path("/holidays/{id}").HandlerFunc(CreateHoliday)
	route.Methods(http.MethodPut).Path("/holidays").HandlerFunc(UpdateHoliday)
	route.Methods(http.MethodDelete).Path("/holidays/{id}").HandlerFunc(DeleteHoliday)

	//locations
	route.Methods(http.MethodGet).Path("/locations").HandlerFunc(GetLocations)
	route.Methods(http.MethodGet).Path("/locations/{id}").HandlerFunc(GetLocation)
	route.Methods(http.MethodPost).Path("/locations/{id}").HandlerFunc(CreateLocation)
	route.Methods(http.MethodPut).Path("/locations").HandlerFunc(UpdateLocation)
	route.Methods(http.MethodDelete).Path("/locations/{id}").HandlerFunc(DeleteLocation)

	//reservations
	route.Methods(http.MethodGet).Path("/reservations").HandlerFunc(GetReservations)
	route.Methods(http.MethodGet).Path("/reservations/{id}").HandlerFunc(GetReservation)
	route.Methods(http.MethodPost).Path("/reservations/{id}").HandlerFunc(CreateReservation)
	route.Methods(http.MethodPut).Path("/reservations").HandlerFunc(UpdateReservation)
	route.Methods(http.MethodDelete).Path("/reservations/{id}").HandlerFunc(DeleteReservation)

	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: route,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

// it can return it with filters also when nothing is there return all
func GetHolidays(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// recive the id only
func GetHoliday(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// resive the whole holiday object
func CreateHoliday(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// resive the whole holiday object
func UpdateHoliday(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// recive the id only
func DeleteHoliday(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// recive the id only
func GetLocation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// resive the whole holiday object
func CreateLocation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// resive the whole holiday object
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// recive the id only
func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

func GetReservations(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// recive the id only
func GetReservation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// resive the whole holiday object
func CreateReservation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// resive the whole holiday object
func UpdateReservation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}

// recive the id only
func DeleteReservation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implimented"))
}
