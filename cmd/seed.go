package cmd

import (
	"context"
	"log"

	"github.com/jsiqbal/ecommerce/config"
	database "github.com/jsiqbal/ecommerce/db"
	"github.com/jsiqbal/ecommerce/repo"
	"github.com/jsiqbal/ecommerce/service"
	"github.com/jsiqbal/ecommerce/util"
	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seed",
	Short: "seeds database initally",
	RunE:  seed,
}

func seed(cmd *cobra.Command, args []string) error {
	dbCnf := config.GetDB()

	// connect to db
	db, err := database.Connect(dbCnf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("---------------Connected to database---------------")

	// create the brands table if it doesn't exist
	_, err = db.Exec(database.DbSchema)
	if err != nil {
		log.Fatal("can not create table: ", err)
	}

	// initialize the repos
	brandRepo := repo.NewBrandRepo(db)
	ctgryRepo := repo.NewCategoryRepo(db)
	spplrRepo := repo.NewSupplierRepo(db)
	productRepo := repo.NewProductRepo(db)

	// -------------------- brand --------------------
	// create a new brand
	newBrand := &service.Brand{Name: "Lenovo", StatusID: 1, CreatedAt: util.GetCurrentTimestamp()}
	brand, err := brandRepo.Add(context.Background(), newBrand)
	if err != nil {
		log.Fatal("can not create brand: ", err)
	}

	// --------------------- category --------------------
	// create a new category
	newCategory := &service.Category{
		Name:      "Laptop",
		StatusID:  1,
		CreatedAt: util.GetCurrentTimestamp(),
	}
	ctgry, err := ctgryRepo.Add(context.Background(), newCategory)
	if err != nil {
		log.Fatal("can not create category: ", err)
	}

	// --------------------- supplier --------------------
	// create a new supplier
	newSpplr := &service.Supplier{
		Name:               "Z Studio",
		Email:              "zstudio@gmail.com",
		Phone:              "1234567895",
		StatusID:           1,
		IsVerifiedSupplier: true,
		CreatedAt:          util.GetCurrentTimestamp(),
	}
	spplr, err := spplrRepo.Add(context.Background(), newSpplr)
	if err != nil {
		log.Fatal("can not create category: ", err)
	}

	createProducts(context.Background(), productRepo, brand.ID, ctgry.ID, spplr.ID)

	log.Println("---------------seed completed successfully---------")

	return nil
}

func createProducts(ctx context.Context, prdRepo service.ProductRepo, brandID, ctgryID, spplrID string) {
	for i := 1; i <= 20; i++ {
		addProduct(ctx, prdRepo, &service.Product{
			Brand: service.Brand{
				ID: brandID,
			},
			Category: service.Category{
				ID: ctgryID,
			},
			Supplier: service.Supplier{
				ID: spplrID,
			},
			ProductStock: service.ProductStock{
				StockQuantity: util.RandomQuantity(),
			},
			Name:           util.RandomOwner(),
			Description:    util.RandomString(20),
			Specifications: util.RandomString(30),
			UnitPrice:      float64(util.RandomMoney()),
			DiscountPrice:  5,
			Tags:           []string{"Laptop"},
			StatusID:       1,
			CreatedAt:      util.GetCurrentTimestamp(),
		})
	}
}

func addProduct(ctx context.Context, prdRepo service.ProductRepo, product *service.Product) {
	_, err := prdRepo.Add(context.Background(), product)
	if err != nil {
		log.Fatal("can not create product: ", err)
	}
}
