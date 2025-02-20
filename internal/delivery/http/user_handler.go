package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/riskibarqy/ihsan-test/datatransfers"
	"github.com/riskibarqy/ihsan-test/internal/constants"
	"github.com/riskibarqy/ihsan-test/internal/domain"
	"github.com/riskibarqy/ihsan-test/internal/usecase"
	utils "github.com/riskibarqy/ihsan-test/pkg"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

// GetUser get user by id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error(err)
		return datatransfers.Error(c, http.StatusBadRequest, constants.UserInvalidID, err.Error())
	}

	user, err := h.usecase.GetUserByID(c.Context(), id)
	if err != nil {
		log.Error(err)
		return datatransfers.Error(c, http.StatusNotFound, constants.UserNotFound, err.Error())
	}

	return datatransfers.Success(c, user)
}

// RegisterUser register user
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		log.Error(err)
		return datatransfers.Error(c, http.StatusBadRequest, constants.UserInvalidRequest, err.Error())
	}

	user.NoRekening = utils.GenerateNoRekening()
	user.NoHp = utils.CleanPhoneNumber(user.NoHp)

	err := h.usecase.CreateUser(c.Context(), user)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), constants.SQLErrorDuplicate) {
			return datatransfers.Error(c, http.StatusBadRequest, constants.UserFailedCreateDuplicated, "user telah terdaftar")
		}
		return datatransfers.Error(c, http.StatusInternalServerError, constants.UserFailedCreate, err.Error())
	}

	return datatransfers.Success(c, user)
}
