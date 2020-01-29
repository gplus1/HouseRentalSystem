package handler

import (
	"fmt"
	"github.com/gplus1/HouseRentalSystem/entity"
	"github.com/gplus1/HouseRentalSystem/rate"
	"github.com/gplus1/HouseRentalSystem/sessions"
	"html/template"
	"net/http"
	"strconv"
)

/ UserRateHandler handles category handler admin requests
type UserRateHandler struct {
	tmpl    *template.Template
	rateSrv rate.RateService
}


/ NewRatingHandler initializes and returns new UserRateHandler
func NewRatingHandler(template *template.Template, rs rate.RateService) *UserRateHandler {
	return &UserRateHandler{tmpl: template, rateSrv: rs}
}

/ AddRate hanlde requests on route
func (entitypointer *UserRateHandler) AddRate(w http.ResponseWriter, r *http.Request)  {
	ratedb :=entity.RatingApi{}
   id,success:=sessions.IsLogged(r)
 if !success{
	 http.Redirect(w, r, "/signup", http.StatusSeeOther)
	 return
 }
	rateValueFromHidden:=r.FormValue("hidden_rate_value_container")
	if r.Method == "POST" && rateValueFromHidden!="0"{
		fmt.Println("inside if")
		rateinput,err:=strconv.Atoi(rateValueFromHidden)
		if err!=nil{
			panic(err)
		}
		/structslice,_:=entitypointer.rateSrv.GetUserId(&ratedb)
		singleuserval, _:=entitypointer.rateSrv.GetUserRateValue(uint(id))
		  if singleuserval==nil{
		  	/if user is new
			  fmt.Println("new")
			 ratedb.RateValue= rateinput
			 ratedb.UserId= uint(id)
			 _, errr:=entitypointer.rateSrv.AddRate(&ratedb)
		 if len(errr)>0{panic(errr)}

		 }else{
		 	/if user is not new
		 	fmt.Println("nor new")
			  errr:=entitypointer.rateSrv.UpdateUserRateValue(uint(id),rateinput)

                if len(errr)>0{
                	panic(errr)
				}
		 }
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	}




