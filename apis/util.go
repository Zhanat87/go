package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/Zhanat87/go/util"
)

const (
	DEFAULT_PAGE_SIZE int = 10
	MAX_PAGE_SIZE int = 100
	WITHOUT_LIMIT_PAGE_SIZE int = 100000000
)

func getPaginatedListFromRequest(c *routing.Context, count int) *util.PaginatedList {
	page := parseInt(c.Query("page"), 1)
	perPage := parseInt(c.Query("per_page"), DEFAULT_PAGE_SIZE)
	if (c.Query("limit") == "false") {
		perPage = WITHOUT_LIMIT_PAGE_SIZE
	} else {
		if perPage <= 0 {
			perPage = DEFAULT_PAGE_SIZE
		}
		if perPage > MAX_PAGE_SIZE {
			perPage = MAX_PAGE_SIZE
		}
	}
	return util.NewPaginatedList(page, perPage, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
