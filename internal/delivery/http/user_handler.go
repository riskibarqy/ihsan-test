package http

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/riskibarqy/ihsan-test/internal/constants"
	"github.com/riskibarqy/ihsan-test/internal/domain"
	"github.com/riskibarqy/ihsan-test/internal/usecase"
	"github.com/riskibarqy/ihsan-test/response"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, constants.UserInvalidID, err.Error())
	}

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		return response.Error(c, http.StatusNotFound, constants.UserNotFound, err.Error())
	}

	return response.Success(c, user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return response.Error(c, http.StatusBadRequest, constants.UserInvalidRequest, err.Error())
	}

	err := h.usecase.CreateUser(user)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, constants.UserFailedCreate, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
