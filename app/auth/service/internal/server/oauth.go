package server

import (
	"context"
	"log"

	"net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
)

var clientMap = map[string]*models.Client{
	"1": {
		ID:     "test_client_1",
		Secret: "test_secret_1",
		Domain: "http://localhost:8002",
	},
}

func initClientStore() *store.ClientStore {
	clientStore := store.NewClientStore()
	for _, v := range clientMap {
		clientStore.Set(v.ID, &models.Client{
			ID:     v.ID,
			Secret: v.Secret,
			Domain: v.Domain,
		})
	}

	return clientStore
}

func initManage() *manage.Manager {
	manager := manage.NewDefaultManager()
	// manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := initClientStore()
	manager.MapClientStorage(clientStore)

	return manager
}

func NewOAuthServer() *server.Server {
	manager := initManage()

	srv := server.NewDefaultServer(manager)
	srv.SetAllowedGrantType(oauth2.AuthorizationCode, oauth2.Refreshing)
	srv.SetAllowGetAccessRequest(true)

	srv.SetAuthorizeScopeHandler(authorizeScopeHandler)
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)
	srv.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)

	srv.SetInternalErrorHandler(internalErrorHandler)
	srv.SetResponseErrorHandler(responseErrorHandler)

	return srv
}

func internalErrorHandler(err error) (re *errors.Response) {
	log.Println("Internal Error:", err.Error())
	return
}

func responseErrorHandler(re *errors.Response) {
	log.Println("Response Error:", re.Error.Error())
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)

	store.Delete("LoggedInUserID")
	store.Save()

	return
}

func passwordAuthorizationHandler(ctx context.Context, username, password string) (userID string, err error) {
	if username == "test" && password == "123456" {
		userID = "test"
	}
	return
}

func authorizeScopeHandler(w http.ResponseWriter, r *http.Request) (scope string, err error) {
	if r.Form == nil {
		r.ParseForm()
	}

	// r.Form.Get("client_id")
	// r.Form.Get("scope")
	return
}
