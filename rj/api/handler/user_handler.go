package handler

import (
	"api/model"
	"api/repository"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) CreateUser (c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name: req.Name,
		Email: req.Email,
		Bio:   req.Bio,
	}

	if err := h.userRepo.CreateUser(&user); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			log.Println(pgErr)
			c.JSON(http.StatusConflict, gin.H{"error": "Email address already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetAllUsers (c *gin.Context) {
	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserById (c *gin.Context) {
	id := c.Param("id")

	user, err := h.userRepo.GetUserById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser (c *gin.Context) {
	id := c.Param("id")

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userRepo.UpdateUser(id, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) HardDeleteUser (c *gin.Context) {
	id := c.Param("id")

	if err := h.userRepo.HardDeleteUser(id); err != nil {
		if (errors.Is(err, sql.ErrNoRows)) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) SoftDeleteUser (c *gin.Context) {
	id := c.Param("id")

	if err := h.userRepo.SoftDeleteUser(id); err != nil {
		if (errors.Is(err, sql.ErrNoRows)) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}