package ginpgstore

import (
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/tester"
)

const pgTestServer = "user=test password=test dbname=test sslmode=disable"

var newPGStore = func(_ *testing.T) sessions.Store {
	store, err := NewPGStore(pgTestServer, []byte("secret"))
	if err != nil {
		panic(err)
	}
	return store
}

func TestPGSQL_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newPGStore)
}

func TestPGSQL_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newPGStore)
}

func TestPGSQL_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newPGStore)
}

func TestPGSQL_SessionClear(t *testing.T) {
	tester.Clear(t, newPGStore)
}

func TestPGSQL_SessionOptions(t *testing.T) {
	tester.Options(t, newPGStore)
}
