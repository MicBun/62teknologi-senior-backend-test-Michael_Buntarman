package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/service"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/web"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloEndpoint(t *testing.T) {
	web.RunTest(func(c *service.Container) {
		w, err := web.MakeRequest(c.Web, http.MethodGet, "/hello", nil)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp struct {
			Message string
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "Hello", resp.Message)
	})
}

func TestCreateBusiness(t *testing.T) {
	web.RunTest(func(c *service.Container) {
		testBusiness := core.Business{
			Name: "Test Business",
		}
		body, err := json.Marshal(testBusiness)
		assert.NoError(t, err)
		ioReader := bytes.NewReader(body)
		w, err := web.MakeRequest(c.Web, http.MethodPost, "/business", ioReader)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp struct {
			Message string
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "Business created", resp.Message)

		// wrong json
		body, err = json.Marshal("test")
		assert.NoError(t, err)
		ioReader = bytes.NewReader(body)
		w, err = web.MakeRequest(c.Web, http.MethodPost, "/business", ioReader)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Name not provided
		testBusiness = core.Business{
			Alias: "test",
		}
		body, err = json.Marshal(testBusiness)
		assert.NoError(t, err)
		ioReader = bytes.NewReader(body)
		w, err = web.MakeRequest(c.Web, http.MethodPost, "/business", ioReader)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestApiHandler_EditBusiness(t *testing.T) {
	web.RunTest(func(c *service.Container) {
		testBusiness := core.Business{
			Name: "Test Business",
		}
		err := c.Business.CreateBusiness(&testBusiness)
		assert.NoError(t, err)
		testBusiness.Name = "Test Business 2"
		body, err := json.Marshal(testBusiness)
		assert.NoError(t, err)
		ioReader := bytes.NewReader(body)
		id := strconv.Itoa(int(testBusiness.ID))
		w, err := web.MakeRequest(c.Web, http.MethodPut, "/business/"+id, ioReader)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp struct {
			Message string
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "Business updated", resp.Message)

		// wrong json
		body, err = json.Marshal("test")
		assert.NoError(t, err)
		ioReader = bytes.NewReader(body)
		w, err = web.MakeRequest(c.Web, http.MethodPut, "/business/"+id, ioReader)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// id is not a number
		body, err = json.Marshal(testBusiness)
		assert.NoError(t, err)
		ioReader = bytes.NewReader(body)
		w, err = web.MakeRequest(c.Web, http.MethodPut, "/business/abc", ioReader)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// updated name is duplicate with another business
		testBusiness2 := core.Business{
			Name: "Test Business 3",
		}
		err = c.Business.CreateBusiness(&testBusiness2)
		assert.NoError(t, err)
		testBusiness2.Name = "Test Business 2"
		body, err = json.Marshal(testBusiness2)
		assert.NoError(t, err)
		ioReader = bytes.NewReader(body)
		id = strconv.Itoa(int(testBusiness2.ID))
		w, err = web.MakeRequest(c.Web, http.MethodPut, "/business/"+id, ioReader)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestApiHandler_DeleteBusiness(t *testing.T) {
	web.RunTest(func(c *service.Container) {
		testBusiness := core.Business{
			Name: "Test Business",
		}
		err := c.Business.CreateBusiness(&testBusiness)
		assert.NoError(t, err)
		id := strconv.Itoa(int(testBusiness.ID))
		w, err := web.MakeRequest(c.Web, http.MethodDelete, "/business/"+id, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp struct {
			Message string
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "Business deleted", resp.Message)

		// id is not a number
		w, err = web.MakeRequest(c.Web, http.MethodDelete, "/business/abc", nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestApiHandler_GetBusinessesByParams(t *testing.T) {
	web.RunTest(func(c *service.Container) {
		testBusiness := core.Business{
			Name: "Test Business",
			Categories: core.Categories{
				{
					Alias: "test",
				},
			},
		}
		err := c.Business.CreateBusiness(&testBusiness)
		assert.NoError(t, err)
		w, err := web.MakeRequest(c.Web, http.MethodGet, "/business/search?open_now=true", nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var resp struct {
			Businesses []core.Business
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(resp.Businesses))
		assert.Equal(t, "Test Business", resp.Businesses[0].Name)

		// result is empty
		w, err = web.MakeRequest(c.Web, http.MethodGet, "/business/search?category=abc", nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
