package main

import (
	"database/sql"
	"net/http"

	"github.com/Kiluameta/events-golang-api/internal/events/infra/repository"
	"github.com/Kiluameta/events-golang-api/internal/events/infra/service"
	"github.com/Kiluameta/events-golang-api/internal/events/usecase"

	httpHandler "github.com/Kiluameta/events-golang-api/internal/events/infra/http"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepository, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://localhost:9080/api1",
		2: "http://localhost:9080/api2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepository)
	getEventUseCase   := usecase.NewGetEventUseCase(eventRepository)
	listSpotsUseCase  := usecase.NewListSpotsUseCase(eventRepository)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepository, partnerFactory)

	eventsHandler := httpHandler.NewEventHandler(
		listEventsUseCase,
		getEventUseCase,
		buyTicketsUseCase,
		listSpotsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)	
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)	
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)	

	http.ListenAndServe(":8080", r)
}