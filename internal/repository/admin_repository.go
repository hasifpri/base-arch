package internalrepository

import (
	"context"
	"github.com/bytesaddict/dancok"
	"gorm.io/gorm"
	"log/slog"
	"math"
	"panel-ektensi/core"
	helperutilities "panel-ektensi/helper/utilities"
	internalentity "panel-ektensi/internal/entity"
)

type AdminRepository struct {
	*Repository[internalentity.Admin]
	Log *slog.Logger
}

func NewAdminRepository(
	log *slog.Logger,
) *AdminRepository {
	sqlGenerator := helperutilities.NewSqlGenerator((&internalentity.Admin{}).TableName(), "created_at")

	repo := NewRepositoryImpl[internalentity.Admin](sqlGenerator)
	return &AdminRepository{
		Repository: repo,
		Log:        log,
	}
}

func (r *AdminRepository) Select(tx *gorm.DB, ctx context.Context, param core.QueryInfo) (result []internalentity.Admin, totalItems int32, page int32, pageSize int32, errres error) {

	// Sort
	var sortDefault dancok.SortDescriptor
	sortDefault.FieldName = r.queryGenerator.DefaultFieldForSort
	sortDefault.SortDirection = dancok.Descending

	if len(param.SelectParameter.SortDescriptors) == 0 {
		param.SelectParameter.SortDescriptors = append(param.SelectParameter.SortDescriptors, sortDefault)
	}

	// default deleted_at
	var filterDeleted dancok.FilterDescriptor
	filterDeleted.Value = "0"
	filterDeleted.Operator = dancok.IsEqual
	filterDeleted.FieldName = "deleted_at"
	filterDeleted.Condition = dancok.And

	param.SelectParameter.FilterDescriptors = append(param.SelectParameter.FilterDescriptors, filterDeleted)

	sqlQuery, sqlQueryCount := r.queryGenerator.Generate(param, (&internalentity.Admin{}).TableName())

	r.Log.Info("AdminRepository.Select()", "queryGenerator.Generate()", "query", sqlQuery)

	queryResult := []internalentity.Admin{}
	queryTotal := []internalentity.Admin{}

	err := tx.Raw(sqlQuery).Scan(&queryResult).Error
	if err != nil {
		r.Log.Error("AdminRepository.Select()", "tx.Raw()", "Error", err)
	}

	err = tx.Raw(sqlQueryCount).Scan(&queryTotal).Error
	if err != nil {
		r.Log.Error("AdminRepository.Select()", "tx.Raw()", "Error", err)
	}

	// Output
	result = queryResult
	totalItems = int32(len(queryTotal))
	page = int32(math.Ceil(float64(totalItems) / float64(param.SelectParameter.PageDescriptor.PageSize)))

	return

}
