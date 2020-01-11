package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gplus1/HouseRentalSystem/Back_End/user"
	"github.com/gplus1/julienschmidt/httprouter"
)

type AdminRoleHandler struct {
	roleService user.RoleService
}

func NewAdminRoleHandler(roleSrv user.RoleService) *AdminRoleHandler {
	return &AdminRoleHandler{roleService: roleSrv}
}

func (arh *AdminRoleHandler) GetRoles(w http.ResponseWriter,
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

func (arh *AdminRoleHandler) GetSingleRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

func (arh *AdminRoleHandler) PutRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

func (arh *AdminRoleHandler) PostRole(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

}

func (arh *AdminRoleHandler) DeleteRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}
