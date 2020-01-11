package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gplus1/HouseRentalSystem/Back_End/user"
	"github.com/gplus1/julienschmidt/httprouter"
)

// ManagerRoleHandler is used to implement role related http requests
type ManagerRoleHandler struct {
	roleService user.RoleService
}

// NewManagerRoleHandler initializes and returns new ManagerRoleHandler object
func NewManagerRoleHandler(roleSrv user.RoleService) *ManagerRoleHandler {
	return &ManagerRoleHandler{roleService: roleSrv}
}

// GetRoles handles GET /v1/Manager/roles requests
func (arh *ManagerRoleHandler) GetRoles(w http.ResponseWriter,
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

// GetSingleRole handles GET /v1/Manager/roles/:id requests
func (arh *ManagerRoleHandler) GetSingleRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

// PutRole handles PUT /v1/Manager/roles/:id requests
func (arh *ManagerRoleHandler) PutRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

// PostRole handles POST /v1/Manager/roles requests
func (arh *ManagerRoleHandler) PostRole(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

}

// DeleteRole handles DELETE /v1/Manager/roles/:id requests
func (arh *ManagerRoleHandler) DeleteRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}
