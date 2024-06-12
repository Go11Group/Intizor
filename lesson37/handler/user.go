package handler

import (
	"github.com/Go11Group/Intizor/lesson37/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	var filter model.Filter

	age, err := strconv.Atoi(c.Query("age"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	filter.Age = age

	filter.Gender = c.Query("gender")
	filter.Nation = c.Query("nation")
	filter.Field = c.Query("field")

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	filter.Limit = limit

	if ofs, ok := c.GetQuery("offset"); ok {
		offset, err := strconv.Atoi(ofs)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}
		filter.Offset = offset
	}
	users, err := h.User.GetAll(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	c.JSON(http.StatusOK, users)
}