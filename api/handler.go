package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"postgres-test/test/internal/postgres-test/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler struct holds dependencies for HTTP handlers
type Handler struct {
	Client *db.SQLDBClient
}

// NewHandler initializes a new Handler with dependencies
func NewHandler(client *db.SQLDBClient) *Handler {
	return &Handler{
		Client: client,
	}
}

// SetupRoutes sets up the HTTP routes for your API
func (h *Handler) SetupRoutes(r *gin.Engine) {
	r.GET("/data/all", h.getAllData)
	r.GET("/data", h.getDataById)
	r.Run(":8080")
}

func (h *Handler) getAllData(c *gin.Context) {
	data, err := h.Client.Select()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": []interface{}{data}})
}

func (h *Handler) getDataById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id parameter not defined"})
		return
	}

	if id < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id must be bigger than 0"})
		return
	}

	data, err := h.Client.GetDataById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound,
				gin.H{"error": fmt.Sprintf("there's no element with id: %s", fmt.Sprint(id))})
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
