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

	cu, rErr := services.UsersService.CreateUser(user)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, cu.Marshal(c.GetHeader("X-Public") == "true"))
}

// Get handles incoming request to get a user
func Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid user id, must be a number"))
		return
	}

	user, rErr := services.UsersService.GetUser(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
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

	u, rErr := services.UsersService.UpdateUser(isPatch, user)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusOK, u.Marshal(c.GetHeader("X-Public") == "true"))
}

// Delete handles incoming request to delete a user
func Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid user id, must be a number"))
		return
	}

	if err := services.UsersService.DeleteUser(id); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search handles incoming request to search for users
func Search(c *gin.Context) {
	status := c.Query("status")
	users, rErr := services.UsersService.Search(status)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}

// Login handles incoming request to login a user
func Login(c *gin.Context) {
	var user users.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
		return
	}

	u, rErr := services.UsersService.LoginUser(user)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusOK, u.Marshal(c.GetHeader("X-Public") == "true"))
}
