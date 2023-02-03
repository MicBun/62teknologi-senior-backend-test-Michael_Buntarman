package web

import (
	"fmt"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/database"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/service"
	"io"
	"net/http"
	"net/http/httptest"
)

type TestFunc func(*service.Container)

func RunTest(testFunc TestFunc) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	if err := database.Migrate(db); err != nil {
		panic(err)
	}

	tx := db.Begin()
	defer tx.Rollback()
	c := service.New(tx)
	RegisterAPIRoutes(c)
	testFunc(c)
}

func MakeRequest(h http.Handler, method string, endpoint string, body io.Reader, headers ...map[string]string) (*httptest.ResponseRecorder, error) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, endpoint, body)
	if len(headers) > 0 {
		for key, val := range headers[0] {
			req.Header.Add(key, val)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("Error building request %w", err)
	}
	h.ServeHTTP(w, req)
	return w, nil
}
