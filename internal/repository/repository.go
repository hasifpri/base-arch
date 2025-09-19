package repository

import (
	"gorm.io/gorm"
	helperutilities "panel-ektensi/helper/utilities"
)

type Repository[T any] struct {
	queryGenerator *helperutilities.SqlGenerator
}

func NewRepositoryImpl[T any](queryGenerator *helperutilities.SqlGenerator) *Repository[T] {
	return &Repository[T]{
		queryGenerator: queryGenerator,
	}
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Select("deleted_at").Updates(entity).Error
}
