package trailcategoriesbundle

import (
	"net/http"
	"strconv"

	"github.com/davidalfasunarna/trail-backend-mvc/app/core"
	"github.com/gorilla/mux"
)

// TrailCategoriesController struct
type TrailCategoriesController struct {
	core.Controller
	km TrailCategoriesMapper
}

// NewTrailCategoriesController instance
func NewTrailCategoriesController(km TrailCategoriesMapper) *TrailCategoriesController {
	return &TrailCategoriesController{
		Controller: core.Controller{},
		km:         km,
	}
}

// Index func return all trail categories in database
func (c *TrailCategoriesController) Index(w http.ResponseWriter, r *http.Request) {
	k, err := c.km.FindAll()

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, &k, http.StatusOK)
}

// Create trail category
func (c *TrailCategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	var k TrailCategory

	if err := c.GetContent(&k, r); err != nil {
		return
	}

	if !k.Validate() {
		c.SendJSON(w, k.Errors, http.StatusBadRequest)
		return
	}

	// Insert trail category and handle error
	if c.HandleError(c.km.Insert(&k), w) {
		return
	}

	c.SendJSON(w, &k, http.StatusOK)
}

// Delete a trail category
func (c *TrailCategoriesController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if c.HandleError(err, w) {
		return
	}

	// Insert trail category and handle error
	if c.HandleError(c.km.Delete(id), w) {
		return
	}

	c.SendJSON(w, nil, http.StatusNoContent)
}
