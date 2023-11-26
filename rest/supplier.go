package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jsiqbal/ecommerce/logger"
	"github.com/jsiqbal/ecommerce/service"
	"github.com/jsiqbal/ecommerce/util"
)

// @Summary Create a new supplier
// @Description Create a new supplier with the provided details
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param request body createSupplierReq true "Supplier details to create"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/suppliers [post]
func (s *Server) createSupplier(ctx *gin.Context) {
	var req createSupplierReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	spplrID, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot create supplier id using uuid generator", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Supplier ID generation failed", "failed to generate supplier id"))
		return
	}

	spplr := &service.Supplier{
		ID:                 spplrID.String(),
		Name:               req.Name,
		Email:              req.Email,
		Phone:              req.Phone,
		StatusID:           req.StatusID,
		IsVerifiedSupplier: req.IsVerifiedSupplier,
		CreatedAt:          util.GetCurrentTimestamp(),
	}

	newSpplr, err := s.svc.AddSupplier(ctx, spplr)
	if err != nil {
		logger.Error(ctx, "cannot add supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", newSpplr)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully created", newSpplr))
}

// @Summary Get a supplier by ID
// @Description Get details of a supplier based on the provided ID
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID" format "uuid"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/suppliers/{id} [get]
func (s *Server) getSupplier(ctx *gin.Context) {
	var req getSupplierReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	spplr, err := s.svc.GetSupplier(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", spplr)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched", spplr))
}

// @Summary Get a list of suppliers
// @Description Get a list of suppliers with pagination support
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param page query int true "Page number" minimum 1
// @Param limit query int true "Number of items per page" minimum 1 maximum 100
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/suppliers [get]
func (s *Server) getSuppliers(ctx *gin.Context) {
	var req getSuppliersReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	result, err := s.svc.GetSuppliers(ctx, req.Page, req.Limit)
	if err != nil {
		logger.Error(ctx, "cannot get suppliers", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", result)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched suppliers", result))
}

// @Summary Update a supplier by ID
// @Description Update details of a supplier based on the provided ID
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID" format "uuid"
// @Param request body updateSupplierReq true "Supplier details to update"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/suppliers/{id} [put]
func (s *Server) updateSupplier(ctx *gin.Context) {
	var req updateSupplierReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	spplrID := ctx.Param("id")
	if len(spplrID) == 0 {
		logger.Error(ctx, "cannot pass validation", spplrID)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Invalid supplier ID", "Bad request"))
		return
	}

	logger.Info(ctx, fmt.Sprintf("req payload for supplierID: %s", spplrID), req)

	spplr, err := s.svc.GetSupplier(ctx, spplrID)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if spplr == nil {
		logger.Error(ctx, "supplier not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Supplier Not Found", "Not found"))
		return
	}

	spplr.Name = req.Name
	spplr.Email = req.Email
	spplr.Phone = req.Phone
	spplr.StatusID = req.StatusID
	spplr.IsVerifiedSupplier = req.IsVerifiedSupplier

	err = s.svc.UpdateSupplier(ctx, spplrID, spplr)
	if err != nil {
		logger.Error(ctx, "cannot update supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", spplr)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully updated", spplr))
}

// @Summary Delete a supplier by ID
// @Description Delete a supplier based on the provided ID
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID" format "uuid"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/suppliers/{id} [delete]
func (s *Server) deleteSupplier(ctx *gin.Context) {
	var req deleteSupplierReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	spplr, err := s.svc.GetSupplier(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if spplr == nil {
		logger.Error(ctx, "supplier not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Supplier Not Found", "Not found"))
		return
	}

	err = s.svc.DeleteSupplier(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot delete supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", spplr)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully deleted", spplr))
}
