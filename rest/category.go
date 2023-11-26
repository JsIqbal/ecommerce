package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsiqbal/ecommerce/logger"
	"github.com/jsiqbal/ecommerce/service"
	"github.com/jsiqbal/ecommerce/util"
)

// TreeNode represents a node in the category tree
type TreeNode struct {
	Category service.Category
	Children []*TreeNode `json:"children"`
}

type Category struct {
	ID           string      `json:"id"`
	CategoryName string      `json:"category_name"`
	Children     []*Category `json:"children"`
}

// @Summary Create a new category
// @Description Create a new category with the provided details
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body createCategoryReq true "Category details to create"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/categories [post]
func (s *Server) createCategory(ctx *gin.Context) {
	var req createCategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	ctgry := &service.Category{
		Name:      req.Name,
		ParentID:  req.ParentID,
		StatusID:  req.StatusID,
		CreatedAt: util.GetCurrentTimestamp(),
	}

	newCategory, err := s.svc.AddCategory(ctx, ctgry)
	if err != nil {
		logger.Error(ctx, "cannot add category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", newCategory)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully created", newCategory))
}

// @Summary Get a category by ID
// @Description Get details of a category based on the provided ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID" format "uuid"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/categories/{id} [get]
func (s *Server) getCategory(ctx *gin.Context) {
	var req getCategoryReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	ctgry, err := s.svc.GetCategory(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", ctgry)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched", ctgry))
}

// @Summary Get a list of categories
// @Description Get a paginated list of categories based on the provided parameters
// @Tags Categories
// @Accept json
// @Produce json
// @Param page query int true "Page number (starting from 1)"
// @Param limit query int true "Number of items per page (min: 1, max: 100)"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/categories [get]
func (s *Server) getCategories(ctx *gin.Context) {
	var req getCategoriesReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	result, err := s.svc.GetCategories(ctx, req.Page, req.Limit)
	if err != nil {
		logger.Error(ctx, "cannot get categories", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", result)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched categories", result))
}

// @Summary Get a formatted list of categories
// @Description Get a formatted list of categories with hierarchical structure
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/categories/tree [get]
func (s *Server) getFormattedCategories(ctx *gin.Context) {
	result, err := s.svc.GetCategories(ctx, 1, service.MAX_INF)
	if err != nil {
		logger.Error(ctx, "cannot get categories", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", result)

	treeNodes := buildCategoryTree(result.Categories, "")

	// convert TreeNode data to an array of Category objects
	var categories []*Category
	for _, node := range treeNodes {
		category := convertTreeNodeToCategory(node)
		categories = append(categories, category)
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched categories", categories))
}

// @Summary Update a category
// @Description Update an existing category with the provided details
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID" format "uuid"
// @Param request body updateCategoryReq true "Category details to update"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/categories/{id} [put]
func (s *Server) updateCategory(ctx *gin.Context) {
	var req updateCategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	ctgryID := ctx.Param("id")
	if len(ctgryID) == 0 {
		logger.Error(ctx, "cannot pass validation", ctgryID)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Invalid brand ID", "Bad request"))
		return
	}

	logger.Info(ctx, fmt.Sprintf("req payload for categoryID: %s", ctgryID), req)

	ctgry, err := s.svc.GetCategory(ctx, ctgryID)
	if err != nil {
		logger.Error(ctx, "cannot get category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if ctgry == nil {
		logger.Error(ctx, "category not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Category Not Found", "Not found"))
		return
	}

	// update category
	ctgry.Name = req.Name
	ctgry.StatusID = req.StatusID

	err = s.svc.UpdateCategory(ctx, ctgryID, ctgry)
	if err != nil {
		logger.Error(ctx, "cannot update category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", ctgry)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully updated", ctgry))
}

// @Summary Delete a category
// @Description Delete an existing category based on the provided ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID" format "uuid"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/categories/{id} [delete]
func (s *Server) deleteCategory(ctx *gin.Context) {
	var req deleteCategoryReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	ctgry, err := s.svc.GetCategory(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if ctgry == nil {
		logger.Error(ctx, "category not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Category Not Found", "Not found"))
		return
	}

	err = s.svc.DeleteCategory(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot delete category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", ctgry)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully deleted", ctgry))
}

// buildCategoryTree builds a tree structure from a list of categories
func buildCategoryTree(categories []service.Category, parentID string) []*TreeNode {
	var categoryTree []*TreeNode

	// create a map to quickly look up categories by ID
	categoryMap := make(map[string]*TreeNode)

	// create tree nodes for each category and add them to the map
	for _, category := range categories {
		node := &TreeNode{Category: category}
		categoryMap[category.ID] = node
	}

	// build the tree structure for the specified parent ID
	for _, category := range categories {
		node := categoryMap[category.ID]
		if category.ParentID == parentID {
			// this category has the specified parent, add it to the tree
			node.Children = buildCategoryTree(categories, category.ID)
			categoryTree = append(categoryTree, node)
		}
	}

	// // sort children based on sequence
	// sort.Slice(categoryTree, func(i, j int) bool {
	// 	return categoryTree[i].Sequence < categoryTree[j].Sequence
	// })

	return categoryTree
}

func convertTreeNodeToCategory(node *TreeNode) *Category {
	category := &Category{
		ID:           node.Category.ID,
		CategoryName: node.Category.Name,
	}

	for _, childNode := range node.Children {
		childCategory := convertTreeNodeToCategory(childNode)
		category.Children = append(category.Children, childCategory)
	}

	return category
}
