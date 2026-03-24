package handlers

import (
	"net/http"

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

type PlanQuery struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

func ListPlans(c *gin.Context) {
	// TODO: load from DB, filter by merchant
	plans := []Plan{}
	c.JSON(http.StatusOK, gin.H{"plans": plans})
}
