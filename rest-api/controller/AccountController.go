package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"rest-api/model"
	"rest-api/repository"
	"strconv"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

// ListAccounts godoc
// @Summary List accounts
// @Description List accounts
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Account
// @Failure 401 {object} ErrResponse
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /accounts [get]
func (h *BaseHandler) GetAllAccounts(ctx *gin.Context) {
	var accounts []model.Account
	err := repository.GetAllAccounts(h.db, &accounts)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, DataResponse(accounts))
	}
}

// GetAccount godoc
// @Summary Get Account Information
// @Description Get Account Information
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 401 {object} ErrResponse
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /accounts/:id [get]
func (h *BaseHandler) GetAccount(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 8)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	var account model.Account
	err = repository.GetAccount(h.db, &account, uint(id))
	if err != nil && err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, ErrorResponse(err))
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, DataResponse(account))
}

// CreateAccount godoc
// @Summary Create Account
// @Description Create Account
// @Accept  json
// @Produce  json
// @Param account body model.Account true "New Account"
// @Success 200 {object} model.Account
// @Failure 401 {object} ErrResponse
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /accounts [post]
func (h *BaseHandler) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	var account = model.Account{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	err := repository.CreateAccount(h.db, &account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, DataResponse(account))
}
