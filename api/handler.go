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
	r.POST("/data/add", h.insertData)
	r.DELETE("/data/del", h.deleteData)
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
	id, err := validateId(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.Client.GetDataById(id)
	if err != nil {
		// not found in db
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound,
				gin.H{"error": fmt.Sprintf("there's no element with id: %s", fmt.Sprint(id))})
		} else { // other error while quering db
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) insertData(c *gin.Context) {
	var data struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	// retrive data from request
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "not all required parameters passed in post body"})
		return
	}

	id, err := h.Client.Insert(data.Title, data.Description)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) deleteData(c *gin.Context) {
	id, err := validateId(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := h.Client.DeleteById(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"deleted": ok})
}

func validateId(text string) (int, error) {
	id, err := strconv.Atoi(text)
	// parameter not defined
	if err != nil {
		return 0, fmt.Errorf("parameter id not defined")
	}
	// id is invalid
	if id < 1 {
		return id, fmt.Errorf("parameter id must be bigger than 0")
	}

	return id, nil
}
