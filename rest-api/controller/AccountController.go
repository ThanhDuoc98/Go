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

func (h *BaseHandler) GetAllAccounts(ctx *gin.Context) {
	var accounts []model.Account
	err := repository.GetAllAccounts(h.db, &accounts)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": accounts})
	}
}

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
	ctx.JSON(http.StatusOK, account)
}

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
	ctx.JSON(http.StatusOK, account)
}
