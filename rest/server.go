package rest

import (
	"net/http"

	"github.com/jsiqbal/ecommerce/config"
	"github.com/jsiqbal/ecommerce/logger"
	"github.com/jsiqbal/ecommerce/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
	svc    service.Service
	appCnf *config.Application
}

func NewServer(svc service.Service, appCnf *config.Application) (*Server, error) {
	server := &Server{
		svc:    svc,
		appCnf: appCnf,
	}

	// custom validators for status id
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validStatusID", validStatusID)
		v.RegisterValidation("validPhone", validPhone)
	}

	// check env-wise mode enabled
	if server.appCnf.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	// setup routers
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// CORS MIDDLEWARE
	router.Use(corsMiddleware)

	// LOG MIDDLEWARE
	router.Use(logger.ModifyContext)

	//------------------------SWAGGER DOCS ROUTE------------------------
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/api/health", server.checkHealth)

	//------------------------BRAND ROUTES------------------------
	router.POST("/api/brands", server.createBrand)
	router.GET("/api/brands", server.getBrands)
	router.GET("/api/brands/:id", server.getBrand)
	router.PUT("/api/brands/:id", server.updateBrand)
	router.DELETE("/api/brands/:id", server.deleteBrand)

	//------------------------CATEGORY ROUTES------------------------
	router.POST("/api/categories", server.createCategory)
	router.GET("/api/categories", server.getCategories)
	router.GET("/api/categories/tree", server.getFormattedCategories)
	router.GET("/api/categories/:id", server.getCategory)
	router.PUT("/api/categories/:id", server.updateCategory)
	router.DELETE("/api/categories/:id", server.deleteCategory)

	//------------------------SUPPLIER ROUTES------------------------
	router.POST("/api/suppliers", server.createSupplier)
	router.GET("/api/suppliers", server.getSuppliers)
	router.GET("/api/suppliers/:id", server.getSupplier)
	router.PUT("/api/suppliers/:id", server.updateSupplier)
	router.DELETE("/api/suppliers/:id", server.deleteSupplier)

	//------------------------PRODUCT ROUTES------------------------
	router.POST("/api/products", server.createProduct)
	router.GET("/api/products", server.getProducts)
	router.GET("/api/products/:id", server.getProduct)
	router.PUT("/api/products/:id", server.updateProduct)
	router.DELETE("/api/products/:id", server.deleteProduct)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
