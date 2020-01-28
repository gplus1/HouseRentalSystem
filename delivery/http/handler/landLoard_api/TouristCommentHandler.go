package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gplus1/HouseRentalSystem/FeedBack"
	"github.com/gplus1/HouseRentalSystem/entity"
	"github.com/gplus1/julienschmidt/httprouter"
)

type TouristFeedBackHandler struct {
	FeedBackService FeedBack.FeedBackService
}

/ NewTouristFeedBackHandler returns new TouristFeedBackHandler object
func NewTouristFeedBackHandler(cmntService FeedBack.FeedBackService) *TouristFeedBackHandler {
	return &TouristFeedBackHandler{FeedBackService: cmntService}
}

/ GetFeedBacks handles GET /v1/Tourist/FeedBacks request
func (ach *TouristFeedBackHandler) GetFeedBacks(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	FeedBacks, errs := ach.FeedBackService.FeedBacks()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(FeedBacks, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

/ GetSingleFeedBack handles GET /v1/Tourist/FeedBacks/:id request
func (ach *TouristFeedBackHandler) GetSingleFeedBack(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	FeedBack, errs := ach.FeedBackService.FeedBack(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(FeedBack, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

/ PostFeedBack handles POST /v1/Tourist/FeedBacks request
func (ach *TouristFeedBackHandler) PostFeedBack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	FeedBack := &entity.FeedBack{}

	err := json.Unmarshal(body, FeedBack)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	FeedBack, errs := ach.FeedBackService.StoreFeedBack(FeedBack)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/Tourist/FeedBacks/%d", FeedBack.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

/ PutFeedBack handles PUT /v1/Tourist/FeedBacks/:id request
func (ach *TouristFeedBackHandler) PutFeedBack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	FeedBack, errs := ach.FeedBackService.FeedBack(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &FeedBack)

	FeedBack, errs = ach.FeedBackService.UpdateFeedBack(FeedBack)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(FeedBack, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

/ DeleteFeedBack handles DELETE /v1/Tourist/FeedBacks/:id request
func (ach *TouristFeedBackHandler) DeleteFeedBack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ach.FeedBackService.DeleteFeedBack(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
