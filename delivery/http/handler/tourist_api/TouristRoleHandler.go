package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gplus1/HouseRentalSystem/user"
	"github.com/gplus1/julienschmidt/httprouter"
)

type TouristRoleHandler struct {
	roleService Tourist.RoleService
}

/ NewTouristRoleHandler initializes and returns new TouristRoleHandler object
func NewTouristRoleHandler(roleSrv Tourist.RoleService) *TouristRoleHandler {
	return &TouristRoleHandler{roleService: roleSrv}
}

/ GetRoles handles GET /v1/Tourist/roles requests
func (arh *TouristRoleHandler) GetRoles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	roles, errs := arh.roleService.Roles()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

/ GetSingleRole handles GET /v1/Tourist/roles/:id requests
func (arh *TouristRoleHandler) GetSingleRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

/ PutRole handles PUT /v1/Tourist/roles/:id requests
func (arh *TouristRoleHandler) PutRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

/ PostRole handles POST /v1/Tourist/roles requests
func (arh *TouristRoleHandler) PostRole(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

}

/ DeleteRole handles DELETE /v1/Tourist/roles/:id requests
func (arh *TouristRoleHandler) DeleteRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}
