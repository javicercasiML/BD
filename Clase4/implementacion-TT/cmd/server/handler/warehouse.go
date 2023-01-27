package handler

import (
	"errors"
	"strconv"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/internal/warehouse"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type warehouseHandler struct {
	s warehouse.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewWarehouseHandler(s warehouse.Service) *warehouseHandler {
	return &warehouseHandler{
		s: s,
	}
}

// Get obtiene un producto por id
func (h *warehouseHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		warehouse, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("warehouse not found"))
			return
		}
		web.Success(c, 200, warehouse)
	}
}

func (h *warehouseHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		warehouse, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 500, errors.New("internal server error"))
			return
		}
		web.Success(c, 200, warehouse)
	}
}

func (h *warehouseHandler) GetReport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request
		var id int
		var err error
		ids, ok := ctx.GetQuery("id")
		if ok {
			id, err = strconv.Atoi(ids)
			if err != nil {
				web.Failure(ctx, 400, errors.New("invalid id"))
				return
			}
		}

		// process
		report, err := h.s.GetReport(id)
		if err != nil {
			web.Failure(ctx, 500, errors.New("internal server error"))
			return
		}

		// response
		web.Success(ctx, 200, report)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptysW(warehouse *domain.Warehouse) (bool, error) {
	switch {
	case warehouse.Name == "" || warehouse.Adress == "" || warehouse.Telephone == "":
		return false, errors.New("fields can't be empty")
	case warehouse.Capacity <= 0:
		if warehouse.Capacity <= 0 {
			return false, errors.New("capacity must be greater than 0")
		}
	}
	return true, nil
}

// Post crea un nuevo producto
func (h *warehouseHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var warehouse domain.Warehouse
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != "123" {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&warehouse)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysW(&warehouse)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		w, err := h.s.Create(warehouse)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, w)
	}
}

// Delete elimina un producto
func (h *warehouseHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != "123" {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un producto
func (h *warehouseHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != "123" {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("product not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var warehouse domain.Warehouse
		err = c.ShouldBindJSON(&warehouse)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysW(&warehouse)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		p, err := h.s.Update(id, warehouse)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un producto o alguno de sus campos
func (h *warehouseHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name      string `json:"name,omitempty"`
		Adress    string `json:"adress,omitempty"`
		Telephone string `json:"telephone,omitempty"`
		Capacity  int    `json:"capacity,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != "123" {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("warehouse not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Warehouse{
			Name:      r.Name,
			Adress:    r.Adress,
			Telephone: r.Telephone,
			Capacity:  r.Capacity,
		}

		w, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, w)
	}
}
