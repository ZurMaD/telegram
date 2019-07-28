package telegram_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nasermirzaei89/telegram"
)

const (
	testToken    = "someToken"
	invalidToken = "invalidToken"
)

func except(t *testing.T, actual, excepted interface{}) {
	if fmt.Sprintf("%+v", excepted) != fmt.Sprintf("%+v", actual) {
		t.Errorf("\nexcepted: %+v\nactual:   %+v", excepted, actual)
	}
}

func notExcept(t *testing.T, actual, notExcepted interface{}) {
	if fmt.Sprintf("%+v", notExcepted) == fmt.Sprintf("%+v", actual) {
		t.Errorf("\nnot excepted: %+v\nactual:       %+v", notExcepted, actual)
	}
}

func TestGetMe(t *testing.T) {
	response := []byte(`{"ok":true,"result":{"id":1,"is_bot":true,"first_name":"Test Bot","username":"TestBot"}}`)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := fmt.Sprintf("/bot%s/getMe", testToken)
		if r.URL.String() != u {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"ok":false,"error_code":401,"description":"Unauthorized"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}))

	defer server.Close()

	telegram.BaseURL = server.URL

	// success
	bot := telegram.New(testToken)

	res, err := bot.GetMe()
	except(t, err, nil)

	except(t, res.IsOK(), true)
	except(t, res.GetErrorCode(), 0)
	notExcept(t, res.GetUser(), nil)
	except(t, res.GetUser().ID, 1)
	except(t, res.GetUser().IsBot, true)
	except(t, res.GetUser().FirstName, "Test Bot")
	except(t, res.GetUser().LastName, nil)
	notExcept(t, res.GetUser().Username, nil)
	except(t, *res.GetUser().Username, "TestBot")
	except(t, res.GetUser().LanguageCode, nil)

	// fail
	bot = telegram.New(invalidToken)

	res, err = bot.GetMe()
	except(t, err, nil)

	except(t, res.IsOK(), false)
	except(t, res.GetUser(), nil)
	except(t, res.GetErrorCode(), http.StatusUnauthorized)
	except(t, res.GetDescription(), "Unauthorized")
	except(t, res.GetParameters(), nil)
}
