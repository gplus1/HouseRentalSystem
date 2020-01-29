package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gplus1/HouseRentalSystem/Tourist"
	"github.com/gplus1/julienschmidt/httprouter"
)

 /landLordRoleHandler is used to implement role related http requests
type landLordRoleHandler struct {
	roleService Tourist.RoleService
}

/ NewlandLordRoleHandler initializes and returns new landLordRoleHandler object
func NewlandLordRoleHandler(roleSrv Tourist.RoleService) *landLordRoleHandler {
	return &landLordRoleHandler{roleService: roleSrv}
}

/ GetRoles handles GET /v1/landLord/roles requests
func (arh *landLordRoleHandler) GetRoles(w http.ResponseWriter,
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

/ GetSingleRole handles GET /v1/landLord/roles/:id requests
func (arh *landLordRoleHandler) GetSingleRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

/ PutRole handles PUT /v1/landLord/roles/:id requests
func (arh *landLordRoleHandler) PutRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

/ PostRole handles POST /v1/landLord/roles requests
func (arh *landLordRoleHandler) PostRole(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

}

/ DeleteRole handles DELETE /v1/landLord/roles/:id requests
func (arh *landLordRoleHandler) DeleteRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}
