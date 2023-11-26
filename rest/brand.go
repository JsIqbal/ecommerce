package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsiqbal/ecommerce/logger"
	"github.com/jsiqbal/ecommerce/service"
	"github.com/jsiqbal/ecommerce/util"
)

func (s *Server) createBrand(ctx *gin.Context) {
	var req createBrandReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	brand := &service.Brand{
		Name:      req.Name,
		StatusID:  req.StatusID,
		CreatedAt: util.GetCurrentTimestamp(),
	}

	newBrand, err := s.svc.AddBrand(ctx, brand)
	if err != nil {
		logger.Error(ctx, "cannot add brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", newBrand)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully created", newBrand))
}

func (s *Server) getBrand(ctx *gin.Context) {
	var req getBrandReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	brand, err := s.svc.GetBrand(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", brand)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched", brand))
}

func (s *Server) getBrands(ctx *gin.Context) {
	var req getBrandsReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	result, err := s.svc.GetBrands(ctx, req.Page, req.Limit)
	if err != nil {
		logger.Error(ctx, "cannot get brands", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", result)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched brands", result))
}

func (s *Server) updateBrand(ctx *gin.Context) {
	var req updateBrandReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	brandID := ctx.Param("id")
	if len(brandID) == 0 {
		logger.Error(ctx, "cannot pass validation", brandID)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Invalid brand ID", "Bad request"))
		return
	}

	logger.Info(ctx, fmt.Sprintf("req payload for brandID: %s", brandID), req)

	brand, err := s.svc.GetBrand(ctx, brandID)
	if err != nil {
		logger.Error(ctx, "cannot get brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if brand == nil {
		logger.Error(ctx, "brand not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Brand Not Found", "Not found"))
		return
	}

	// update brand
	brand.Name = req.Name
	brand.StatusID = req.StatusID

	err = s.svc.UpdateBrand(ctx, brandID, brand)
	if err != nil {
		logger.Error(ctx, "cannot update brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", brand)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully updated", brand))
}

func (s *Server) deleteBrand(ctx *gin.Context) {
	var req deleteBrandReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	brand, err := s.svc.GetBrand(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if brand == nil {
		logger.Error(ctx, "brand not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Brand Not Found", "Not found"))
		return
	}

	err = s.svc.DeleteBrand(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot delete brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", brand)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully deleted", brand))
}
