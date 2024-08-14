package controller

import (
	"net/http"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SignUpController struct {
	SignUpService domain.SignUpService
	Env           *config.Env
}

func (sc *SignUpController) SignUp(c *gin.Context){
	var req domain.SignUpRequest

	err := c.ShouldBind(&req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := sc.SignUpService.CheckUsernameExists(c, req.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check username"})
        return
    }

	if exists {
        c.JSON(http.StatusConflict, gin.H{"error": domain.ErrUsernameExists.Error()})
        return
    }

	salt, err := utils.GenerateSalt()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate salt"})
		return
	}
	
	localParam := sc.Env.LocalParam
	
	hashedPassword, err := utils.HashPassword(req.Password, salt, localParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to hash password"})
		return
	}

	author := &domain.Author{
		ID: uuid.New(),
		Username: req.Username,
		Password: hashedPassword,
		Salt:	  salt,
	}

	err = sc.SignUpService.Create(c, author)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, domain.SuccessfullResponse{
		Message: "you signed up successfully!",
	})
}