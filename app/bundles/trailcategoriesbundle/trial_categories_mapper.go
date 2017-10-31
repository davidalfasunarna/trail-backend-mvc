package trailcategoriesbundle

// TrailCategoriesMapper define the base contract for mapper
type TrailCategoriesMapper interface {
	FindAll() ([]TrailCategory, error)
	Insert(*TrailCategory) error
	Delete(int) error
}
