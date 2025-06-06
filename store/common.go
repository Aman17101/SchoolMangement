package store

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Aman17101/SchoolMangement/model"
	"github.com/Aman17101/SchoolMangement/util"
	"gorm.io/gorm"
)

func setLimitAndPage(filter map[string]interface{}, query *gorm.DB) {
	// Convert limit and page to integers
	limit := filter[model.DataPerPage]
	page := filter[model.DataPerPage]
	if limit == "" {
		limit = "10" // Default limit is 10
	}
	if page == "" {
		page = "1" // Default page is 1
	}
	limitInt, err := strconv.Atoi(fmt.Sprintf("%v", limit))
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.SetLimitAndPage,
			"error while while converting limit to int", err)
	}
	pageInt, err := strconv.Atoi(fmt.Sprintf("%v", page))
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.SetLimitAndPage,
			"error while while converting page to int", err)
	}

	// Apply pagination
	offset := (pageInt - 1) * limitInt
	query = query.Limit(limitInt).Offset(offset)

}

func setDateRangeFilter(filter map[string]interface{}, query *gorm.DB) {
	// Convert limit and page to integers
	startDatestr := filter[model.StartDate]
	endDatestr := filter[model.EndDate]
	if startDatestr == "" || endDatestr == "" {
		util.Log(model.LogLevelInfo, model.StorePackage, model.SetDateRangeFilter,
			"Date range not provided", nil)
		return
	}
	// Parse the start and end dates
	startDate, err := time.Parse(model.TimeLayout, fmt.Sprintf("%v", startDatestr))
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.SetDateRangeFilter,
			"error in parsing start date", err)
		return
	}
	endDate, err := time.Parse(model.TimeLayout, fmt.Sprintf("%v", endDatestr))
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.SetDateRangeFilter,
			"error in parsing end date", err)
		return
	}
	query.Where("date_column BETWEEN ? AND ?", startDate, endDate)
}
