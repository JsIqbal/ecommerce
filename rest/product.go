package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsiqbal/ecommerce/logger"
	"github.com/jsiqbal/ecommerce/service"
	"github.com/jsiqbal/ecommerce/util"
)

func (s *Server) createProduct(ctx *gin.Context) {
	var req createProductReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	// check supplier exists
	spplr, err := s.svc.GetSupplier(ctx, req.SupplierID)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if spplr == nil {
		logger.Error(ctx, "Supplier id not found", nil)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Supplier not found", nil))
		return
	}

	// check category exists
	ctgry, err := s.svc.GetCategory(ctx, req.CategoryID)
	if err != nil {
		logger.Error(ctx, "cannot get category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if ctgry == nil {
		logger.Error(ctx, "Category id not found", nil)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Category not found", nil))
		return
	}

	// check brand exists
	brand, err := s.svc.GetBrand(ctx, req.BrandID)
	if err != nil {
		logger.Error(ctx, "cannot get brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if brand == nil {
		logger.Error(ctx, "Brand id not found", nil)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Brand not found", nil))
		return
	}

	// check supplier wise product name uniqueness
	existProduct, err := s.svc.GetProducts(ctx, service.FilterProductsParams{
		Name:       req.Name,
		SupplierID: req.SupplierID,
		Limit:      1,
	})
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if existProduct != nil && len(existProduct.Products) > 0 {
		logger.Error(ctx, "Product name already exists for same supplier", existProduct)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Product name already exists for this supplier", nil))
		return
	}

	product := &service.Product{
		Name:           req.Name,
		Description:    req.Description,
		Specifications: req.Specifications,
		Brand:          *brand,
		Category:       *ctgry,
		Supplier:       *spplr,
		ProductStock: service.ProductStock{
			StockQuantity: req.StockQuantity,
		},
		UnitPrice:     req.UnitPrice,
		DiscountPrice: req.DiscountPrice,
		Tags:          req.Tags,
		StatusID:      req.StatusID,
		CreatedAt:     util.GetCurrentTimestamp(),
	}

	newProduct, err := s.svc.AddProduct(ctx, product)
	if err != nil {
		logger.Error(ctx, "cannot add product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", newProduct)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully created", newProduct))
}

func (s *Server) getProduct(ctx *gin.Context) {
	var req getProductReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	product, err := s.svc.GetProduct(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", product)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched", product))
}

func (s *Server) getProducts(ctx *gin.Context) {
	var req getProductsReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	result, err := s.svc.GetProducts(ctx, service.FilterProductsParams{
		Name:       req.Name,
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		BrandIDs:   req.BrandIDs,
		CategoryID: req.CategoryID,
		SupplierID: req.SupplierID,
		Page:       req.Page,
		Limit:      req.Limit,
	})
	if err != nil {
		logger.Error(ctx, "cannot filter products", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "Res payload", result)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched Products", result))
}

func (s *Server) updateProduct(ctx *gin.Context) {
	var req updateProductReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	productID := ctx.Param("id")
	if len(productID) == 0 {
		logger.Error(ctx, "cannot pass validation", productID)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Invalid product ID", "Bad request"))
		return
	}

	logger.Info(ctx, fmt.Sprintf("req payload for productID: %s", productID), req)

	product, err := s.svc.GetProduct(ctx, productID)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if product == nil {
		logger.Error(ctx, "product not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "product Not Found", "Not found"))
		return
	}

	// update product
	product.Name = req.Name
	product.Description = req.Description
	product.Specifications = req.Specifications
	product.Brand = service.Brand{
		ID: req.BrandID,
	}
	product.Category = service.Category{
		ID: req.CategoryID,
	}
	product.Supplier = service.Supplier{
		ID: req.SupplierID,
	}
	product.ProductStock = service.ProductStock{
		StockQuantity: req.StockQuantity,
	}
	product.UnitPrice = req.UnitPrice
	product.DiscountPrice = req.DiscountPrice
	product.Tags = req.Tags
	product.StatusID = req.StatusID

	err = s.svc.UpdateProduct(ctx, productID, product)
	if err != nil {
		logger.Error(ctx, "cannot update product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", product)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully updated", product))
}

func (s *Server) deleteProduct(ctx *gin.Context) {
	var req deleteProductReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	product, err := s.svc.GetProduct(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if product == nil {
		logger.Error(ctx, "product not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Product Not Found", "Not found"))
		return
	}

	err = s.svc.DeleteProduct(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot delete product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", product)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully deleted", product))
}
