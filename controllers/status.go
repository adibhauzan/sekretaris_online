package controllers

import (
    "net/http"
    "strconv"

    "github.com/adibhauzan/sekretaris_online_backend/models"
    "github.com/adibhauzan/sekretaris_online_backend/repository"
    "github.com/gin-gonic/gin"
)

type StatusController struct {
    StatusRepo repository.StatusRepository
}

func NewStatusController(repo repository.StatusRepository) *StatusController {
    return &StatusController{StatusRepo: repo}
}

func (c *StatusController) CreateStatus(ctx *gin.Context) {
    var status models.Status
    if err := ctx.BindJSON(&status); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.StatusRepo.CreateStatus(&status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, status)
}

func (c *StatusController) GetAllStatus(ctx *gin.Context) {
    statuses, err := c.StatusRepo.GetAllStatus()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, statuses)
}

func (c *StatusController) GetStatusByID(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    status, err := c.StatusRepo.GetStatusByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
        return
    }

    ctx.JSON(http.StatusOK, status)
}

func (c *StatusController) UpdateStatus(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var status models.Status
    if err := ctx.BindJSON(&status); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    status.ID = uint(id)
    if err := c.StatusRepo.UpdateStatus(&status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, status)
}

func (c *StatusController) DeleteStatus(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := c.StatusRepo.DeleteStatus(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusNoContent, nil)
}
