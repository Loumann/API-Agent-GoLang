package handler

import (
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateUserInput struct {
	ID        int    `json:"id" db:"id"`
	AgentName string `json:"agentname" db:"agentname"`
	Status    string `json:"status" db:"status"`
}

func (h *Handler) GetAgents(c *gin.Context) {
	agents, err := h.r.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if agents == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "no agents found"})
	} else {
		c.JSON(http.StatusOK, agents)
	}
}
func (h *Handler) DeleteAgent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.r.DeleteUser(id); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, gin.H{"message": "user deleted"})
}
func (h *Handler) AddAgent(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.r.Create(
		input.AgentName,
		input.Status,
	); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(201, gin.H{"message": "user added"})

}
func (h *Handler) UpdateAgent(c *gin.Context) {
	var agent models.Agent
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid agent ID"})
		return
	}

	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data(or json body)"})
		return
	}
	agent.ID = id

	updated, err := h.r.UpdateAgent(agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update agent"})
		return
	}
	if !updated {
		c.JSON(404, gin.H{"message": "agent not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "agent updated successfully",
			"agent":   agent,
		})
	}
}

func (h *Handler) CreateQuest(c *gin.Context) {
	var inputQuest models.Quest

	if err := c.ShouldBindJSON(&inputQuest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "quest created successfully"})
}
