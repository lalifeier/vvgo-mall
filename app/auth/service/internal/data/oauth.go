package data

import (
	"log"

	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
	"net/http"
	"os"
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
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := initClientStore()
	manager.MapClientStorage(clientStore)

	return manager
}



func NewOauthServer() *server.Server {
	manager := initManage()

	srv := server.NewDefaultServer(manager)
	srv.SetAllowedGrantType("authorization_code", "refresh_token")
	srv.SetAllowGetAccessRequest(true)

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	// srv.SetAllowGetAccessRequest(true)
	// srv.SetClientInfoHandler(server.ClientFormHandler)

	// srv.SetPasswordAuthorizationHandler(func(ctx context.Context, username, password string) (userID string, err error) {
	// 	if username == "test" && password == "test" {
	// 		userID = "test"
	// 	}
	// 	return
	// })
	// srv.SetUserAuthorizationHandler()
	// srv.SetAuthorizeScopeHandler()

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return srv
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
