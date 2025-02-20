package http

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/riskibarqy/ihsan-test/datatransfers"
	"github.com/riskibarqy/ihsan-test/internal/constants"
	"github.com/riskibarqy/ihsan-test/internal/usecase"
)

type UserBalanceHistoryHandler struct {
	usecase *usecase.UserBalanceHistoryUsecase
}

func NewUserBalanceHistoryHandler(uc *usecase.UserBalanceHistoryUsecase) *UserBalanceHistoryHandler {
	return &UserBalanceHistoryHandler{usecase: uc}
}

// AddBalance add funds to user balance
func (h *UserBalanceHistoryHandler) AddBalance(c *fiber.Ctx) error {
	userBalanceHistoryRequest := new(datatransfers.UserBalanceHistoryTXRequest)
	if err := c.BodyParser(userBalanceHistoryRequest); err != nil {
		log.Error(err)
		return datatransfers.Error(c, http.StatusBadRequest, constants.UserBalanceHistoryInvalidRequest, err.Error())
	}

	userBalanceHistoryRequest.TransactionType = constants.TransactionTypeAdd

	response, err := h.usecase.CreateBalance(c.Context(), userBalanceHistoryRequest)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), constants.RecordNotFound) {
			return datatransfers.Error(c, http.StatusBadRequest, constants.UserBalanceHistoryNotFound, "data no rekening tidak ditemukan")
		}
		return datatransfers.Error(c, http.StatusInternalServerError, constants.InternalServerError, err.Error())
	}
	return datatransfers.Success(c, response)
}

// WithdrawBalance withdraw funds from user balance
func (h *UserBalanceHistoryHandler) WithdrawBalance(c *fiber.Ctx) error {
	userBalanceHistoryRequest := new(datatransfers.UserBalanceHistoryTXRequest)
	if err := c.BodyParser(userBalanceHistoryRequest); err != nil {
		log.Error(err)
		return datatransfers.Error(c, http.StatusBadRequest, constants.UserBalanceHistoryInvalidRequest, err.Error())
	}

	userBalanceHistoryRequest.TransactionType = constants.TransactionTypeWithdraw

	response, err := h.usecase.CreateBalance(c.Context(), userBalanceHistoryRequest)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), constants.RecordNotFound) {
			return datatransfers.Error(c, http.StatusBadRequest, constants.UserBalanceHistoryNotFound, "data no rekening tidak ditemukan")
		}
		return datatransfers.Error(c, http.StatusInternalServerError, constants.InternalServerError, err.Error())
	}
	return datatransfers.Success(c, response)
}
