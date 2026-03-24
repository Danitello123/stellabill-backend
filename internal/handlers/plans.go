package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"stellarbill-backend/internal/pagination"

	"github.com/gin-gonic/gin"
)

type Plan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Interval    string `json:"interval"`
	Description string `json:"description,omitempty"`
}

// Helper for extracting sort values dynamically in memory
func getPlanSortVal(p Plan, sortBy string) string {
	switch strings.ToLower(sortBy) {
	case "amount":
		return p.Amount
	case "name":
		return p.Name
	default: // id or fallback
		return p.ID
	}
}

func ListPlans(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	offsetStr := c.DefaultQuery("offset", "0")
	offset, _ := strconv.Atoi(offsetStr)

	sortBy := c.DefaultQuery("sort", "id")
	sortOrder := c.DefaultQuery("order", "asc")

	// TODO: load from DB with offset, limit, and order
	// We use an empty slice until the DB connection is implemented
	var mockDB []Plan

	page, meta := pagination.PaginateList(mockDB, offset, limit, sortBy, sortOrder, getPlanSortVal)

	c.JSON(http.StatusOK, gin.H{
		"data":       page,
		"pagination": meta,
	})
}
