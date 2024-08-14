package handler

import (
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

func (h *Handler) RenderUsersPage(c *gin.Context) error {
	// Получаем пользователей из базы данных
	users, err := h.r.GetUsers(c)
	if err != nil {
		// Возвращаем JSON-ответ с ошибкой
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return err
	}

	// Парсим HTML-шаблон
	tmpl, err := template.ParseFiles("assets/page.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template"})
		return err
	}

	// Передаем данные в шаблон и рендерим его
	if err := tmpl.Execute(c.Writer, users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
		return err
	}

	return nil
}

func (h *Handler) GetAgents(c *gin.Context) {
	if err := h.r.GetUsers(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
func (h *Handler) UpdateAgent(c *gin.Context) {
	if err := h.r.GetUsers(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
func (h *Handler) DeleteUser(c *gin.Context) {
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
func (h *Handler) AddUser(c *gin.Context) {
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
