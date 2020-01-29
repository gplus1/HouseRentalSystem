package handler

import (
	"context"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gplus1/HouseRentalSystem/permission"
	"github.com/gplus1/HouseRentalSystem/rtoken"

	"github.com/gplus1/HouseRentalSystem/entity"
	"github.com/gplus1/HouseRentalSystem/session"
	"golang.org/x/crypto/bcrypt"

	"github.com/gplus1/HouseRentalSystem/form"
	"github.com/gplus1/HouseRentalSystem/user"
)

/ UserHandler handler handles Tourist related requests
type UserHandler struct {
	tmpl           *template.Template
	TouristService    Tourist.UserService
	sessionService Tourist.SessionService
	TouristSess       *entity.Session
	loggedInUser   *entity.Tourist
	TouristRole       Tourist.RoleService
	csrfSignKey    []byte
}

type contextKey string

var ctxUserSessionKey = contextKey("signed_in_Tourist_session")

/ NewUserHandler returns new UserHandler object
func NewUserHandler(t *template.Template, usrServ Tourist.UserService,
	sessServ Tourist.SessionService, uRole Tourist.RoleService,
	usrSess *entity.Session, csKey []byte) *UserHandler {
	return &UserHandler{tmpl: t, TouristService: usrServ, sessionService: sessServ,
		TouristRole: uRole, TouristSess: usrSess, csrfSignKey: csKey}
}

/ Authenticated checks if a Tourist is authenticated to access a given route
func (uh *UserHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.TouristSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

/ Authorized checks if a Tourist has proper authority to access a give route
func (uh *UserHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uh.loggedInUser == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		roles, errs := uh.TouristService.UserRoles(uh.loggedInUser)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		for _, role := range roles {
			permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
			if !permitted {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		if r.Method == http.MethodPost {
			ok, err := rtoken.ValidCSRF(r.FormValue("_csrf"), uh.csrfSignKey)
			if !ok || (err != nil) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

/ Login hanldes the GET/POST /login requests
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}

	if r.Method == http.MethodPost {
		/ Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		usr, errs := uh.TouristService.UserByEmail(r.FormValue("email"))
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(r.FormValue("password")))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}

		uh.loggedInUser = usr
		claims := rtoken.Claims(usr.Email, uh.TouristSess.Expires)
		session.Create(claims, uh.TouristSess.UUID, uh.TouristSess.SigningKey, w)
		newSess, errs := uh.sessionService.StoreSession(uh.TouristSess)
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Failed to store session")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		uh.TouristSess = newSess
		roles, _ := uh.TouristService.UserRoles(usr)
		if uh.checkLandLord(roles) {
			http.Redirect(w, r, "/LandLord", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

/ Logout hanldes the POST /logout requests
func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	TouristSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	session.Remove(TouristSess.UUID, w)
	uh.sessionService.DeleteSession(TouristSess.UUID)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

/ Signup hanldes the GET/POST /signup requests
func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		signUpForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
		return
	}

	if r.Method == http.MethodPost {
		/ Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		/ Validate the form contents
		singnUpForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		singnUpForm.Required("fullname","UserName", "email","Country","State","City", "password", "confirmpassword")
		singnUpForm.MatchesPattern("email", form.EmailRX)
		singnUpForm.MatchesPattern("phone", form.PhoneRX)
		singnUpForm.MinLength("password", 8)
		singnUpForm.PasswordMatches("password", "confirmpassword")
		singnUpForm.CSRF = token
		/ If there are any errors, redisplay the signup form.
		if !singnUpForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		pExists := uh.TouristService.PhoneExists(r.FormValue("phone"))
		if pExists {
			singnUpForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}
		eExists := uh.TouristService.EmailExists(r.FormValue("email"))
		if eExists {
			singnUpForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			singnUpForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		role, errs := uh.TouristRole.RoleByName("Tourist")

		if len(errs) > 0 {
			singnUpForm.VErrors.Add("role", "could not assign role to the Tourist")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		Tourist := &entity.Tourist{
			FullName: r.FormValue("fullname"),
			UserName: r.FormValue("Touristname"),
			Email:    r.FormValue("email"),
			Country: r.FormValue("country"),
			State: r.FormValue("state"),
			City: r.FormValue("city"),
			Password: string(hashedPassword),
			RoleID:   role.ID,
		}
		_, errs = uh.TouristService.StoreUser(Tourist)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (uh *UserHandler) loggedIn(r *http.Request) bool {
	if uh.TouristSess == nil {
		return false
	}
	TouristSess := uh.TouristSess
	c, err := r.Cookie(TouristSess.UUID)
	if err != nil {
		return false
	}
	ok, err := session.Valid(c.Value, TouristSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}

/ LandLordUsers handles Get /LandLord/Tourists request
func (uh *UserHandler) LandLordUsers(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	Tourists, errs := uh.TouristService.Users()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		Users   []entity.Tourist
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		Users:   Tourists,
		CSRF:    token,
	}
	uh.tmpl.ExecuteTemplate(w, "LandLord.Tourists.layout", tmplData)
}

/ LandLordUsersNew handles GET/POST /LandLord/Tourists/new request
func (uh *UserHandler) LandLordUsersNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		roles, errs := uh.TouristRole.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		accountForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   []entity.Role
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			Roles:   roles,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.new.layout", accountForm)
		return
	}

	if r.Method == http.MethodPost {
		/ Parse the form data
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		/ Validate the form contents
		FullName: r.FormValue("fullname"),
		UserName: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Country: r.FormValue("country"),
		State: r.FormValue("state"),
		City: r.FormValue("city"),
		Licence: r.FormValue("licence"),
		Password: string(hashedPassword),
		RoleID:   role.ID,
		accountForm.CSRF = token
		/ If there are any errors, redisplay the signup form.
		if !accountForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.new.layout", accountForm)
			return
		}

		pExists := uh.TouristService.PhoneExists(r.FormValue("phone"))
		if pExists {
			accountForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.new.layout", accountForm)
			return
		}
		eExists := uh.TouristService.EmailExists(r.FormValue("email"))
		if eExists {
			accountForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.new.layout", accountForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			accountForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.new.layout", accountForm)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			accountForm.VErrors.Add("role", "could not retrieve role id")
			uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.new.layout", accountForm)
			return
		}
		Search : entity.Search{
			ID: Search.ID,
			Country: r.FormValue("country"),
		}
		func (sh *UserHandler)checkSearch(sc []entity.Country){
			for _, r:=range sc {
				if strings.ToUpper(r.Country)==String.ToUpper("country"){
					return true
				}

			}
			http.Redirect(w, r, "/Tourists/SearchPage", http.StatusSeeOther)
		}
		Tourist := &entity.Tourist{
			ID:       Tourist.ID,
			FullName: r.FormValue("fullname"),
			UserName: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Country: r.FormValue("country"),
			State:    r.FormValue("state"),
			City:    r.FormValue("city"),
			Password: Tourist.Password,
			ConfirmPassword: Tourist.ConfirmPassword,
			RoleID:   uint(roleID),
		}
		}
		_, errs := uh.TouristService.StoreUser(Tourist)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/LandLord/Tourists", http.StatusSeeOther)
	}
}


func (uh *UserHandler) checkLandLord(rs []entity.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("LandLord") {
			return true
		}
	}
	return false
}

/ LandLordUsersUpdate handles GET/POST /LandLord/Tourists/update?id={id} request
func (uh *UserHandler) LandLordUsersUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		Tourist, errs := uh.TouristService.Tourist(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		roles, errs := uh.TouristRole.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		role, errs := uh.TouristRole.Role(Tourist.RoleID)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		values := url.Values{}
		values.Add("Touristid", idRaw)
		values.Add("fullname", Tourist.FullName)
		values.Add("username", Tourist.UserName)
		values.Add("Email", (Tourist.Email))
		values.Add("country", (Tourist.Country))
		values.Add("state", Tourist.State)
		values.Add("city", Tourist.City)
		values.Add("rolename", role.Name)

		upAccForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   []entity.Role
			Tourist    *entity.Tourist
			CSRF    string
		}{
			Values:  values,
			VErrors: form.ValidationErrors{},
			Roles:   roles,
			Tourist:    Tourist,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.update.layout", upAccForm)
		return
	}

	if r.Method == http.MethodPost {
		/ Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		/ Validate the form contents
		upAccForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		upAccForm.Required("fullname", "username","email", )
		upAccForm.MatchesPattern("email", form.EmailRX)
		upAccForm.CSRF = token
		/ If there are any errors, redisplay the signup form.
		if !upAccForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.update.layout", upAccForm)
			return
		}
		TouristID := r.FormValue("Touristid")
		uid, err := strconv.Atoi(TouristID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		Tourist, errs := uh.TouristService.Tourist(uint(uid))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		eExists := uh.TouristService.EmailExists(r.FormValue("email"))
		if (Tourist.Email != r.FormValue("email")) && eExists {
			upAccForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.update.layout", upAccForm)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			upAccForm.VErrors.Add("role", "could not retrieve role id")
			uh.tmpl.ExecuteTemplate(w, "LandLord.Tourist.update.layout", upAccForm)
			return
		}
		usr := &entity.Tourist{
			ID:       Tourist.ID,
			FullName: r.FormValue("fullname"),
			UserName: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Country: r.FormValue("country"),
			State:    r.FormValue("state"),
			City:    r.FormValue("city"),
			Password: Tourist.Password,
			ConfirmPassword: Tourist.ConfirmPassword,
			RoleID:   uint(roleID),
		}
		_, errs = uh.TouristService.UpdateUser(usr)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/LandLord/Tourists", http.StatusSeeOther)
	}
}

/ LandLordUsersDelete handles Delete /LandLord/Tourists/delete?id={id} request
func (uh *UserHandler) LandLordUsersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := uh.TouristService.DeleteUser(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/LandLord/Tourists", http.StatusSeeOther)
}
