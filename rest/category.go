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

func (s *Server) createCategory(ctx *gin.Context) {
	var req createCategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	// TODO: need to check whether the parent category id exists

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
