package data

import (
	"log"

	"github.com/go-oauth2/oauth2/generates"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

var clientMap = map[string]*models.Client{
	"1": {
		ID:     "test_client_1",
		Secret: "test_secret_1",
		Domain: "http://localhost",
	},
}

func NewOauthServer() *server.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := store.NewClientStore()
	for _, v := range clientMap {
		clientStore.Set(v.ID, &models.Client{
			ID:     v.ID,
			Secret: v.Secret,
			Domain: v.Domain,
		})
	}
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
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
