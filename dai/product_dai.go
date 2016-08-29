package dai

import (
	"database/sql"
	"os"

	"github.com/ajanthan/product-api-go/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type ProductDAIMysql struct {
}

const (
	CREATE_TABLE_SQL = "CREATE TABLE IF NOT EXISTS product (\n" +
		"  product_id INT(11) NOT NULL AUTO_INCREMENT,\n" +
		"  product_name VARCHAR(45) DEFAULT NULL,\n" +
		"  product_prize FLOAT DEFAULT 0.0,\n" +
		"  product_rating FLOAT DEFAULT 0.0,\n" +
		"  product_category VARCHAR(200) DEFAULT NULL,\n" +
		"  product_instock BOOLEAN DEFAULT FALSE,\n" +
		"  PRIMARY KEY (product_id)\n" +
		") ENGINE=InnoDB"
	INSERT_PRODUCT = "INSERT INTO product (" +
		"product_name," +
		"product_category," +
		"product_prize," +
		"product_rating," +
		"product_instock)" +
		" values(?,?,?,?,?)"
	DELETE_PRODUCT_BY_NAME = "DELETE FROM product WHERE product_name=?"
	FIND_PRODUCT_BY_NAME   = "SELECT " +
		"product_id," +
		"product_name," +
		"product_category," +
		"product_prize," +
		"product_rating," +
		"product_instock" +
		" FROM product WHERE product_name=?"
	GET_ALL_PRODUCTS = "SELECT " +
		"product_id," +
		"product_name," +
		"product_category," +
		"product_prize," +
		"product_rating," +
		"product_instock" +
		" FROM product"
)

func (productDAIMysql ProductDAIMysql) Init() error {
	db, err := productDAIMysql.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	createTBStatement, err := db.Prepare(CREATE_TABLE_SQL)
	if err != nil {
		return err
	}
	defer createTBStatement.Close()

	if _, err := createTBStatement.Exec(); err != nil {
		return err
	}
	return nil
}
func (productDAIMysql ProductDAIMysql) AddProduct(product model.Product) error {
	db, err := productDAIMysql.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	insertStatement, err := db.Prepare(INSERT_PRODUCT)
	if err != nil {
		return err
	}
	defer insertStatement.Close()

	if _, err := insertStatement.Exec(product.Name, product.Category, product.Prize, product.Rating, product.InStock); err != nil {
		return err
	}

	return nil
}
func (productDAIMysql ProductDAIMysql) GetProductByID(productID string) (model.Product, error) {
	product := model.Product{}
	db, err := productDAIMysql.GetDB()
	defer db.Close()
	if err != nil {
		return product, err
	}
	selectStmt, err := db.Prepare(FIND_PRODUCT_BY_NAME)
	if err != nil {
		return product, err
	}
	defer selectStmt.Close()
	resultRow := selectStmt.QueryRow(productID)

	resultErr := resultRow.Scan(&product.ID, &product.Name, &product.Category, &product.Prize, &product.Rating, &product.InStock)
	if err != nil {

		return product, resultErr
	}
	return product, nil
}
func (productDAIMysql ProductDAIMysql) DeleteProduct(productID string) error {
	db, err := productDAIMysql.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	deleteStmt, err := db.Prepare(DELETE_PRODUCT_BY_NAME)
	if err != nil {
		return err
	}
	defer deleteStmt.Close()

	_, exError := deleteStmt.Exec(productID)
	if exError != nil {
		return exError
	}
	return nil
}
func (productDAIMysql ProductDAIMysql) GetAllProduct() ([]model.Product, error) {
	var products []model.Product

	db, err := productDAIMysql.GetDB()
	defer db.Close()
	if err != nil {
		return products, err
	}
	getAllRows, err := db.Query(GET_ALL_PRODUCTS)
	if err != nil {
		return products, err
	}
	defer getAllRows.Close()

	for getAllRows.Next() {
		product := model.Product{}
		if err := getAllRows.Scan(&product.ID, &product.Name, &product.Category, &product.Prize, &product.Rating, &product.InStock); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (productDAIMysql ProductDAIMysql) GetDB() (*sql.DB, error) {
	var productDB *sql.DB
	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		return productDB, errors.New("DB_USERNAME is not set")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return productDB, errors.New("DB_PASSWORD is not set")
	}
	dbHostPort := os.Getenv("DB_URL")
	if dbHostPort == "" {
		return productDB, errors.New("DB_URL is not set")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return productDB, errors.New("DB_NAME is not set")
	}

	if db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHostPort+")/"+dbName); err != nil {
		return productDB, err
	} else {
		productDB = db
	}

	if err := productDB.Ping(); err != nil {
		return productDB, err
	}
	return productDB, nil
}
