package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Subscription struct {
	ID        string `json:"id"`
	PlanID    string `json:"plan_id"`
	Customer  string `json:"customer"`
	Status    string `json:"status"`
	Amount    string `json:"amount"`
	Interval  string `json:"interval"`
	NextBilling string `json:"next_billing,omitempty"`
}

type SubscriptionQuery struct {
	Page   int    `form:"page" binding:"omitempty,min=1"`
	Limit  int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Status string `form:"status" binding:"omitempty,oneof=active inactive canceled"`
}

type SubscriptionPath struct {
	ID string `uri:"id" binding:"required,uuid4"` // Assuming ID should be a UUID
}

func ListSubscriptions(c *gin.Context) {
	// TODO: load from DB, filter by merchant from JWT/API key
	subscriptions := []Subscription{}
	c.JSON(http.StatusOK, gin.H{"subscriptions": subscriptions})
}

func GetSubscription(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subscription id required"})
		return
	}
	// TODO: load from DB by id
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"status": "placeholder",
	})
}
