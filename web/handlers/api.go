package handlers

import (
	"fmt"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type apiHandler struct {
	container *service.Container
}

type ApiHandlerInterface interface {
	Hello(c *gin.Context)
	CreateBusiness(c *gin.Context)
	EditBusiness(c *gin.Context)
	DeleteBusiness(c *gin.Context)
	GetBusinessesByParams(c *gin.Context)
}

func NewApiHandler(container *service.Container) ApiHandlerInterface {
	return &apiHandler{
		container: container,
	}
}

// Hello godoc
// @Summary Hello
// @Description Hello
// @Tags Hello
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /hello [get]
func (h *apiHandler) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})
	return
}

type BusinessRequest struct {
	Name         string `json:"name"`
	Alias        string `json:"alias"`
	ImageURL     string `json:"image_url"`
	IsClosed     bool   `json:"is_closed"`
	URL          string `json:"url"`
	CategoriesID []uint `json:"categories_id"`
	Coordinates  struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Transactions []string `json:"transactions"`
	Price        string   `json:"price"`
	Location     struct {
		Address1       string `json:"address1"`
		Address2       string `json:"address2"`
		Address3       string `json:"address3"`
		City           string `json:"city"`
		ZipCode        string `json:"zip_code"`
		Country        string `json:"country"`
		State          string `json:"state"`
		DisplayAddress []string
	} `json:"location"`
	Phone        string  `json:"phone"`
	DisplayPhone string  `json:"display_phone"`
	Distance     float32 `json:"distance"`
}

// CreateBusiness godoc
// @Summary Create a business
// @Description Create a business
// @Tags business
// @Accept  json
// @Produce  json
// @Param business body BusinessRequest true "Business"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Router /business [post]
func (h *apiHandler) CreateBusiness(c *gin.Context) {
	var input BusinessRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	business := &core.Business{
		Name:     input.Name,
		Alias:    input.Alias,
		ImageURL: input.ImageURL,
		IsClosed: input.IsClosed,
		URL:      input.URL,
		Coordinates: core.Coordinates{
			Latitude:  input.Coordinates.Latitude,
			Longitude: input.Coordinates.Longitude,
		},
		Transactions: input.Transactions,
		Price:        input.Price,
		Location: core.Location{
			Address1:       input.Location.Address1,
			Address2:       input.Location.Address2,
			Address3:       input.Location.Address3,
			City:           input.Location.City,
			ZipCode:        input.Location.ZipCode,
			Country:        input.Location.Country,
			State:          input.Location.State,
			DisplayAddress: input.Location.DisplayAddress,
		},
		Phone:        input.Phone,
		DisplayPhone: input.DisplayPhone,
		Distance:     input.Distance,
	}

	err := h.container.Business.CreateBusiness(business)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Business created"})
	return
}

// EditBusiness godoc
// @Summary Edit a business
// @Description Edit a business
// @Tags business
// @Accept  json
// @Produce  json
// @Param id path int true "Business ID"
// @Param business body BusinessRequest true "Business"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Router /business/{id} [put]
func (h *apiHandler) EditBusiness(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input BusinessRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	business := &core.Business{
		Name:     input.Name,
		Alias:    input.Alias,
		ImageURL: input.ImageURL,
		IsClosed: input.IsClosed,
		URL:      input.URL,
		Coordinates: core.Coordinates{
			Latitude:  input.Coordinates.Latitude,
			Longitude: input.Coordinates.Longitude,
		},
		Transactions: input.Transactions,
		Price:        input.Price,
		Location: core.Location{
			Address1:       input.Location.Address1,
			Address2:       input.Location.Address2,
			Address3:       input.Location.Address3,
			City:           input.Location.City,
			ZipCode:        input.Location.ZipCode,
			Country:        input.Location.Country,
			State:          input.Location.State,
			DisplayAddress: input.Location.DisplayAddress,
		},
		Phone:        input.Phone,
		DisplayPhone: input.DisplayPhone,
		Distance:     input.Distance,
	}
	business.ID = uint(id)

	err = h.container.Business.UpdateBusiness(business)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Business updated"})
	return
}

// DeleteBusiness godoc
// @Summary Delete a business
// @Description Delete a business
// @Tags business
// @Accept  json
// @Produce  json
// @Param id path int true "Business ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Router /business/{id} [delete]
func (h *apiHandler) DeleteBusiness(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	business := &core.Business{}
	business.ID = uint(id)

	h.container.Business.DeleteBusiness(business)

	c.JSON(200, gin.H{"message": "Business deleted"})
	return
}

// GetBusinessesByParams godoc
// @Summary Get businesses
// @Description Get businesses
// @Tags business
// @Accept  json
// @Produce  json
// @Param location query string false "Location"
// @Param category query string false "Category"
// @Param sort_by query string false "Sort by"
// @Param price query string false "Price"
// @Param open_now query string false "Open now"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Param longitude query string false "Longitude"
// @Param latitude query string false "Latitude"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Router /business/search [get]
func (h *apiHandler) GetBusinessesByParams(c *gin.Context) {
	location := c.Query("location")
	category := c.Query("category")
	sortBy := c.Query("sort_by")
	price := c.Query("price")
	openNow := c.Query("open_now")
	limit := c.Query("limit")
	offset := c.Query("offset")
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")

	openNowBool := false
	if openNow == "true" {
		openNowBool = true
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 1
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	longitudeFloat, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		longitudeFloat = 0
	}
	latitudeFloat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		latitudeFloat = 0
	}

	fmt.Println("location", location, "category", category, "sortBy", sortBy, "price", price, "openNowBool", openNowBool, "limitInt", limitInt, "offsetInt", offsetInt, "longitudeFloat", longitudeFloat, "latitudeFloat", latitudeFloat)
	businesses, err := h.container.Business.GetBusinessByParam(location, category, sortBy, price, openNowBool, limitInt, offsetInt, longitudeFloat, latitudeFloat)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"businesses": businesses})
	return
}
