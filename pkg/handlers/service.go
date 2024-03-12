package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary GetCost
// @Description Получение цены по запросу
// @ID download-file
// @Accept  json
// @Produce  json
// @Param id path string true "File ID"
// @Success 200
// @Router /api/service/cost/:id [get]
func (h *Handler) getCost(c *gin.Context) {
	var data map[string]string

	data, _ = h.services.GetData(100, 25, 15)

	c.JSON(http.StatusOK, data)
}
