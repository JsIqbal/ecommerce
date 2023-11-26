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

	// update spplr
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
