package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adibhauzan/sekretaris_online_backend/models"
	"github.com/adibhauzan/sekretaris_online_backend/repository"
	"github.com/gin-gonic/gin"
)

type JadwalController struct {
	JadwalRepo repository.JadwalRepository
}

func NewJadwalController(repo repository.JadwalRepository) *JadwalController {
	return &JadwalController{JadwalRepo: repo}
}

func (c *JadwalController) CreateJadwal(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}

	if InvalidatedTokens[tokenString] {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}

	var jadwal models.Jadwal
	if err := ctx.BindJSON(&jadwal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jadwal.ExpiryDate = jadwal.Date.AddDate(0, 0, 7)

	jadwal.Date = time.Now()

	if err := c.JadwalRepo.CreateJadwal(&jadwal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, jadwal)

	ctx.JSON(http.StatusCreated, jadwal)
}

func (c *JadwalController) GetAllJadwal(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if InvalidatedTokens[tokenString] {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}
	jadwals, err := c.JadwalRepo.GetAllJadwal()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jadwals)
}

func (c *JadwalController) GetJadwalByID(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if InvalidatedTokens[tokenString] {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}

	jadwal, err := c.JadwalRepo.GetJadwalByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Jadwal not found"})
		return
	}

	ctx.JSON(http.StatusOK, jadwal)
}

func (c *JadwalController) UpdateJadwal(ctx *gin.Context) {

	tokenString := ctx.GetHeader("Authorization")

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jadwal models.Jadwal
	if err := ctx.BindJSON(&jadwal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if InvalidatedTokens[tokenString] {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}

	jadwal.ID = uint(id)

	jadwal.ExpiryDate = jadwal.Date.AddDate(0, 0, 7)

	jadwal.Date = time.Now()

	if err := c.JadwalRepo.UpdateJadwal(&jadwal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jadwal)
}

func (c *JadwalController) DeleteJadwal(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.JadwalRepo.DeleteJadwal(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if InvalidatedTokens[tokenString] {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *JadwalController) GetJadwalByDatetime(ctx *gin.Context) {

	tokenString := ctx.GetHeader("Authorization")

	datetimeStr := ctx.Query("date")
	if datetimeStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'datetime' harus disertakan"})
		return
	}
	if InvalidatedTokens[tokenString] {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}
	datetime, err := time.Parse("2006-01-02 15:04", datetimeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal dan waktu tidak valid. Gunakan format 'YYYY-MM-DD HH:mm'"})
		return
	}

	jadwals, err := c.JadwalRepo.GetJadwalByDatetime(datetime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jadwals)
}
