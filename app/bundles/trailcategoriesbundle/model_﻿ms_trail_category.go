package trailcategoriesbundle

// Trail struct
type TrailCategory struct {
	Trail_category_id int               `json:"trail_category_id" gorm:"AUTO_INCREMENT"`
	Trail_name        string            `json:"trail_name"`
	Errors            map[string]string `json:"-" gorm:"-"`
}

// NewTrailCategory create a new Trail Category
func NewTrailCategory(trail_name string) *TrailCategory {
	return &TrailCategory{
		Trail_name: trail_name,
	}
}

// Validate a Trail Category
func (k *TrailCategory) Validate() bool {
	k.Errors = make(map[string]string)

	if k.Trail_name == "" {
		k.Errors["name"] = "name can not be empty"
	}

	if len(k.Errors) > 0 {
		return false
	}

	return true
}
