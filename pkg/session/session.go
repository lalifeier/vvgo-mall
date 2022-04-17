package session

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
	// "gopkg.in/boj/redistore.v1"
)

type (
	Options struct {
		Name     string
		SecreKey string
		MaxAge   int
	}

	Session struct {
		Store *sessions.CookieStore
		Name  string
	}
)

var (
	store *sessions.CookieStore
	once  sync.Once
)

func GetSession() *sessions.CookieStore {
	once.Do(func() {
		store = NewSession(&Options{
			Name:     "session_id",
			SecreKey: "secret-key",
			MaxAge:   60 * 60 * 30,
		}).Store
	})
	return store
}

func NewSession(o *Options) *Session {
	store := sessions.NewCookieStore([]byte(o.SecreKey))

	// store, err := NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	// store.KeyPrefix("session_")

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   o.MaxAge,
		HttpOnly: true,
	}
	return &Session{
		Store: store,
		Name:  o.Name,
	}
}

func (s *Session) Session(r *http.Request) (session *sessions.Session) {
	session, err := s.Store.Get(r, s.Name)
	if err != nil {
		log.Fatalf("session err : %v", err)
	}
	return
}

func (s *Session) Get(r *http.Request, key string) (val interface{}, err error) {
	session, err := s.Store.Get(r, s.Name)
	if err != nil {
		return
	}

	val = session.Values[key]
	return
}

func (s *Session) Set(w http.ResponseWriter, r *http.Request, key interface{}, val interface{}) (err error) {
	session, err := s.Store.Get(r, s.Name)
	if err != nil {
		return
	}

	session.Values[key] = val
	err = session.Save(r, w)
	return
}

func (s *Session) Delete(w http.ResponseWriter, r *http.Request, key interface{}) (err error) {
	session, err := s.Store.Get(r, s.Name)
	if err != nil {
		return
	}

	delete(session.Values, key)
	err = session.Save(r, w)
	return
}

func (s *Session) Clear(w http.ResponseWriter, r *http.Request) (err error) {
	session, err := s.Store.Get(r, s.Name)
	if err != nil {
		return
	}

	for key := range session.Values {
		s.Delete(w, r, key)
	}
	return
}
