package trailcategoriesbundle

import (
	"github.com/jinzhu/gorm"
)

// TrailCategoriesMapperSQL define a SQL mapper
type TrailCategoriesMapperSQL struct {
	db *gorm.DB
}

// NewTrailCategoriesSQLMapper instance
func NewTrailCategoriesSQLMapper(db *gorm.DB) *TrailCategoriesMapperSQL {
	return &TrailCategoriesMapperSQL{
		db: db,
	}
}

// Find All Trail Categories in database
func (m *TrailCategoriesMapperSQL) FindAll() ([]TrailCategory, error) {
	var trailCategories []TrailCategory

	m.db.Find(&trailCategories)

	return trailCategories, nil
}

// Insert implement TrailCategoriesMapper interface
func (m *TrailCategoriesMapperSQL) Insert(k *TrailCategory) error {
	return m.db.Create(k).Error
}

// Delete implement TrailCategoriesMapper interface
func (m *TrailCategoriesMapperSQL) Delete(id int) error {
	return m.db.Delete(&TrailCategory{Trail_category_id: id}).Error
}
