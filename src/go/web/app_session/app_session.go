package app_session

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"net/http"
	"time"
    "../../configuration"
    "github.com/michaeljs1990/sqlitestore"
)

var _store sessions.Store

const (
    defaultSessionName = "default-app-session"
	sessionKey_loginFamily = "sessionKey_loginFamily"
)

func init() {
    config := configuration.Get().Session
    switch (config.Kind) {
    case configuration.SessionKind_Cookie:
        _store = sessions.NewCookieStore([]byte(config.Cookie.SecretPhrase))
        return
    case configuration.SessionKind_SQLite:
        store, err := sqlitestore.NewSqliteStore(
                config.SQLite.DatabaseFilePath,
                config.SQLite.TableName,
                "/", // TODO 意味のつかめないパラメータ
                config.MaxAge,
                []byte(config.SQLite.SecretPhrase))
        if err != nil {
            panic(err)
        }
        _store = store
        return
    }
    panic(config)
}

type LoginFamily struct {
    Name          string
	FamilyId      string
	LoginDatetime time.Time
}

func SetLoginFamily(loginFamily LoginFamily, w http.ResponseWriter, r *http.Request) {
	d, err := json.Marshal(loginFamily)
	if err != nil {
		panic(err)
	}
	save(sessionKey_loginFamily, string(d), w, r)
}

func GetLoginFamily(r *http.Request) LoginFamily {
	session := mustGetSession(r)
    str := session.Values[sessionKey_loginFamily].(string)
    var family LoginFamily
    if err := json.Unmarshal([]byte(str), &family); err != nil {
        panic(err)
    }
    return family
}

func IsLogin(r *http.Request) bool {
	session := mustGetSession(r)
	_, found := session.Values[sessionKey_loginFamily]
	return found
}

func UnsetLoginFamily(w http.ResponseWriter, r *http.Request) {
	remove(sessionKey_loginFamily, w, r)
}

func get(key string, r *http.Request) interface{} {
	session := mustGetSession(r)
	return session.Values[key]
}

func save(key string, value interface{}, w http.ResponseWriter, r *http.Request) {
	session := mustGetSession(r)
	session.Values[key] = value
	if e := session.Save(r, w); e != nil {
		panic(e)
	}
}

func remove(key string, w http.ResponseWriter, r *http.Request) {
	session := mustGetSession(r)
	delete(session.Values, key)
	if e := session.Save(r, w); e != nil {
		panic(e)
	}
}

func mustGetSession(r *http.Request) *sessions.Session {
	session, err := _store.Get(r, defaultSessionName)
	if err != nil {
		panic(err)
	}
	return session
}
