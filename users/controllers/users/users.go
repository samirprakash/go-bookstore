package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samirprakash/go-bookstore/users/domain/users"
	"github.com/samirprakash/go-bookstore/users/services"
	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

// Create handles incoming request to create a new user
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
		return
	}

	cu, rErr := services.CreateUser(user)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, cu)
}

// Get handles incoming request to get a user
func Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid user id, must be a number"))
		return
	}

	user, rErr := services.GetUser(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update handles incoming request to update a user
func Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid user id, must be a number"))
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
		return
	}

	user.ID = id
	isPatch := c.Request.Method == http.MethodPatch

	u, rErr := services.UpdateUser(isPatch, user)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusOK, u)
}

// Delete handles incoming request to delete a user
func Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid user id, must be a number"))
		return
	}

	if err := services.DeleteUser(id); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search handles incoming request to search for users
func Search(c *gin.Context) {
	status := c.Query("status")
	users, rErr := services.Search(status)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusOK, users)
}
