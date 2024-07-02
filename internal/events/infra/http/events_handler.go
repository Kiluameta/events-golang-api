package http

import (
	"encoding/json"
	"net/http"

	"github.com/Kiluameta/events-golang-api/internal/events/usecase"
)

type EventsHandler struct {
	listEvetnsUseCase *usecase.ListEventsUseCase
	getEventUseCase   *usecase.GetEventUseCase
	buyTicketsUseCase *usecase.BuyTicketsUseCase
	listSpotsUseCase  *usecase.ListSpotsUseCase
}

func NewEventHandler(
	listEvetnsUseCase *usecase.ListEventsUseCase,
	getEventUseCase   *usecase.GetEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
	listSpotsUseCase  *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEvetnsUseCase: listEvetnsUseCase,
		getEventUseCase:   getEventUseCase,
		buyTicketsUseCase: buyTicketsUseCase,
		listSpotsUseCase:  listSpotsUseCase,
	}
}

func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEvetnsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}

	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}

	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 

	output, err := h.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(output)
}