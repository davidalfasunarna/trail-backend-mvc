package trailcategoriesbundle

import (
	"net/http"

	"github.com/davidalfasunarna/trail-backend-mvc/app/core"
	"github.com/jinzhu/gorm"
)

// TrailCategories handle trial categories resources
type TrailCategoriesBundle struct {
	routes []core.Route
}

// NewTrailCategoriesBundle instance
func NewTrailCategoriesBundle(db *gorm.DB) core.Bundle {
	km := NewTrailCategoriesSQLMapper(db)
	kc := NewTrailCategoriesController(km)

	r := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/trail-category",
			Handler: kc.Index,
		},
		core.Route{
			Method:  http.MethodPost,
			Path:    "/trail-category",
			Handler: kc.Create,
		},
		core.Route{
			Method:  http.MethodDelete,
			Path:    "/trail-category/{id}",
			Handler: kc.Delete,
		},
	}

	return &TrailCategoriesBundle{
		routes: r,
	}
}

// GetRoutes implement interface core.Bundle
func (b *TrailCategoriesBundle) GetRoutes() []core.Route {
	return b.routes
}
