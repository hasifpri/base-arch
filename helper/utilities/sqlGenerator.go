package helperutilities

import (
	"fmt"
	"github.com/bytesaddict/dancok"
	"panel-ektensi/core"
)

type SqlGenerator struct {
	TableName           string
	DefaultFieldForSort string
}

func NewSqlGenerator(tableName string, defaultFieldForSort string) *SqlGenerator {
	return &SqlGenerator{tableName, defaultFieldForSort}
}

func (g *SqlGenerator) Generate(param core.QueryInfo, tableName string) (string, string) {

	limit, offset := g.GeneratePageOFFSET(param)

	result := "select * from (select ROW_NUMBER() OVER(" + g.ParseSort(param.SelectParameter, tableName) + ") as RowNumber,* from \"" + g.TableName + "\" " + g.ParseFilter(param.SelectParameter, tableName) + ") AS T LIMIT " + limit + " OFFSET " + offset

	resultCount := "select * from (select ROW_NUMBER() OVER(" + g.ParseSort(param.SelectParameter, tableName) + ") as RowNumber,* from \"" + g.TableName + "\" " + g.ParseFilter(param.SelectParameter, tableName) + ") AS T"
	return result, resultCount
}

func (g *SqlGenerator) GenerateCount(param core.QueryInfo, tableName string) string {

	resultCount := "select count(rownumber) as total from (select ROW_NUMBER() OVER(" + g.ParseSort(param.SelectParameter, tableName) + ") as RowNumber,* from \"" + g.TableName + "\" " + g.ParseFilter(param.SelectParameter, tableName) + ") AS T"

	return resultCount
}

func (g *SqlGenerator) GenerateIncomeJoin(param core.QueryInfo, tableName, conditionJoin, selectData, selectIncome string) string {

	result := "select " + selectIncome + " from (select ROW_NUMBER() OVER(" + g.ParseSort(param.SelectParameter, tableName) + ") as RowNumber," + selectData + " from " + g.TableName + " JOIN " + conditionJoin + ") AS T " + g.ParseFilter(param.SelectParameter, "T")

	return result
}

func (g *SqlGenerator) GenerateJoin(param core.QueryInfo, tableName, conditionJoin, selectData string) (string, string) {
	limit, offset := g.GeneratePageOFFSET(param)

	result := "select * from (select ROW_NUMBER() OVER(" + g.ParseSort(param.SelectParameter, tableName) + ") as RowNumber," + selectData + " from " + g.TableName + " JOIN " + conditionJoin + ") AS T " + g.ParseFilter(param.SelectParameter, "T") + " LIMIT " + limit + " OFFSET " + offset

	resultCount := "select * from (select ROW_NUMBER() OVER(" + g.ParseSort(param.SelectParameter, tableName) + ") as RowNumber," + selectData + " from " + g.TableName + " JOIN " + conditionJoin + ") AS T " + g.ParseFilter(param.SelectParameter, "T")
	return result, resultCount
}

func (g *SqlGenerator) GenerateLeftJoin(param core.QueryInfo, tableName, conditionJoin, selectData string) (string, string) {
	limit, offset := g.GeneratePageOFFSET(param)

	result := "select * from (select ROW_NUMBER() OVER(" + g.ParseSort(param.SelectParameter, tableName) + ") as RowNumber," + selectData + " from " + g.TableName + " LEFT JOIN " + conditionJoin + ") AS T " + g.ParseFilter(param.SelectParameter, "T") + " LIMIT " + limit + " OFFSET " + offset

	resultCount := "select * from (select ROW_NUMBER() OVER(" + g.ParseSort(param.SelectParameter, tableName) + ") as RowNumber," + selectData + " from " + g.TableName + " JOIN " + conditionJoin + ") AS T " + g.ParseFilter(param.SelectParameter, "T")
	return result, resultCount
}

func (g *SqlGenerator) GeneratePageOFFSET(param core.QueryInfo) (string, string) {

	var limit string
	if param.SelectParameter.PageDescriptor.PageSize != -1 {
		limit = fmt.Sprintf("%d", param.SelectParameter.PageDescriptor.PageSize)
	} else {
		limit = fmt.Sprintf("%s", "ALL")
	}
	offset := fmt.Sprintf("%d", param.SelectParameter.PageDescriptor.PageSize*(param.SelectParameter.PageDescriptor.PageIndex-1))

	return limit, offset
}

func (g *SqlGenerator) Parse(param dancok.SelectParameter, tableName string) string {
	result := g.ParseFilter(param, tableName) + g.ParseSort(param, tableName)

	return result
}

func (g *SqlGenerator) ParseFilter(param dancok.SelectParameter, tableName string) string {
	filterText := ""
	if len(param.FilterDescriptors) > 0 {
		filterText = " WHERE "
		isFirstFilter := true
		for _, filter := range param.FilterDescriptors {
			if isFirstFilter {
				filterText = filterText + tableName + "." + filter.FieldName
				isFirstFilter = false
			} else {
				filterText = filterText + " AND " + tableName + "." + filter.FieldName
				//  Uncomment if need OR and AND config
				// if filter.Condition == dancok.And {
				// 	filterText = filterText + " AND " + filter.FieldName
				// } else {
				// 	filterText = filterText + " OR " + filter.FieldName
				// }
			}

			switch opt := filter.Operator; opt {
			case dancok.IsEqual:
				filterText = filterText + " = '" + filter.Value.(string) + "'"
			case dancok.IsNotEqual:
				filterText = filterText + " != '" + filter.Value.(string) + "'"
			case dancok.IsLessThan:
				filterText = filterText + " < " + filter.Value.(string)
			case dancok.IsLessThanOrEqual:
				filterText = filterText + " <= " + filter.Value.(string)
			case dancok.IsMoreThan:
				filterText = filterText + " > " + filter.Value.(string)
			case dancok.IsMoreThanOrEqual:
				filterText = filterText + " >= " + filter.Value.(string)
			case dancok.IsContain:
				filterText = filterText + " ILIKE '%" + filter.Value.(string) + "%'"
			case dancok.IsBeginWith:
				filterText = filterText + " LIKE '" + filter.Value.(string) + "%'"
			case dancok.IsEndWith:
				filterText = filterText + " LIKE '%" + filter.Value.(string) + "'"
			case dancok.IsBetween:
				filterText = filterText + " BETWEEN '" + filter.Value.(string) + "' AND '" + filter.Value2.(string) + "'"
			case dancok.IsIn:
				filterText = filterText + " IN (" + ParseRangeValues(filter.RangeValues) + ")"
			case dancok.IsNotIn:
				filterText = filterText + " NOT IN (" + ParseRangeValues(filter.RangeValues) + ")"
			}
		}
	}

	if len(param.CompositeFilterDescriptors) > 0 {
		isFirstCompositeFilter := true
		for _, filter := range param.CompositeFilterDescriptors {
			if isFirstCompositeFilter {
				if filterText == "" {
					filterText = " WHERE ("
				} else {
					filterText = filterText + " " + string(filter.Condition) + " ("
				}
				isFirstCompositeFilter = false
			} else {
				filterText = filterText + " AND ("
				//  Uncomment if need OR and AND config
				// if filter.Condition == dancok.And {
				// 	filterText = filterText + " AND ("
				// } else {
				// 	filterText = filterText + " OR ("
				// }
			}

			isFirstItem := true
			for _, item := range filter.GroupFilterDescriptor.Items {
				if isFirstItem {
					switch opt := item.Operator; opt {
					case dancok.IsEqual:
						filterText = filterText + tableName + "." + item.FieldName + " = '" + item.Value.(string) + "'"
					case dancok.IsNotEqual:
						filterText = filterText + tableName + "." + item.FieldName + " != '" + item.Value.(string) + "'"
					case dancok.IsLessThan:
						filterText = filterText + tableName + "." + item.FieldName + " < " + item.Value.(string)
					case dancok.IsLessThanOrEqual:
						filterText = filterText + tableName + "." + item.FieldName + " <= " + item.Value.(string)
					case dancok.IsMoreThan:
						filterText = filterText + tableName + "." + item.FieldName + " > " + item.Value.(string)
					case dancok.IsMoreThanOrEqual:
						filterText = filterText + tableName + "." + item.FieldName + " >= " + item.Value.(string)
					}

					isFirstItem = false
				} else {
					switch opt := item.Operator; opt {
					case dancok.IsEqual:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + tableName + "." + item.FieldName + " = '" + item.Value.(string) + "'"
					case dancok.IsNotEqual:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + tableName + "." + item.FieldName + " != '" + item.Value.(string) + "'"
					case dancok.IsLessThan:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + tableName + "." + item.FieldName + " < " + item.Value.(string)
					case dancok.IsLessThanOrEqual:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + tableName + "." + item.FieldName + " <= " + item.Value.(string)
					case dancok.IsMoreThan:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + tableName + "." + item.FieldName + " > " + item.Value.(string)
					case dancok.IsMoreThanOrEqual:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + tableName + "." + item.FieldName + " >= " + item.Value.(string)
					}
				}
			}

			filterText = filterText + ")"
		}
	}

	return filterText
}

func (g *SqlGenerator) ParseSort(param dancok.SelectParameter, tableName string) string {
	sortText := " "

	if len(param.SortDescriptors) > 0 {
		isFirstSort := true
		sortText = sortText + "order by"
		for _, sort := range param.SortDescriptors {
			if isFirstSort {
				sortText = sortText + " " + tableName + "." + sort.FieldName
				isFirstSort = false
			} else {
				sortText = sortText + "," + tableName + "." + sort.FieldName
			}

			if sort.SortDirection == dancok.Ascending {
				sortText = sortText + " asc"
			} else {
				sortText = sortText + " desc"
			}
		}
	} else {
		sortText = sortText + " order by " + g.DefaultFieldForSort + " desc"
	}

	return sortText
}

func ParseRangeValues(values []any) string {
	valueText := ""
	if len(values) > 0 {
		isFirstValue := true
		_, isStringType := values[0].(string)
		if isStringType {
			for _, v := range values {
				if isFirstValue {
					valueText = "'" + v.(string) + "'"
					isFirstValue = false
				} else {
					valueText = valueText + ",'" + v.(string) + "'"
				}
			}
		} else {
			for _, v := range values {
				if isFirstValue {
					valueText = fmt.Sprint(v.(int64))
					isFirstValue = false
				} else {
					valueText = valueText + "," + fmt.Sprint(v.(int64))
				}
			}
		}
	}
	return valueText
}
