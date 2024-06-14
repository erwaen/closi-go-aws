
package database

import (
    "lambda-func/types"
)
const SESSION_TABLE = "closi_sessions"

type SessionStore interface {
    InsertSession(session types.Session) error
    GetSession(session string) (types.Session, error)
}
