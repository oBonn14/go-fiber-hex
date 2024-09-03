package repository

import (
	"context"
	"log"
	"time"

	"github.com/oBonn14/go-fiber-hex/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	database *mongo.Database
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		db,
	}
}

func (pr *ProductRepository) fromEntity(product model.Product) Product {
	return Product{
		Product: product.Product,
		Stock:   product.Stock,
	}
}

type Product struct {
	ID      primitive.ObjectID `bson:"_id"`
	Product string             `bson:"product"`
	Stock   int                `bson:"stock"`
}

func (product *Product) toEntity() *model.Product {
	objId, err := primitive.ObjectIDFromHex(product.ID.Hex())
	if err != nil {
		return nil
	}

	return &model.Product{
		ID:      objId,
		Product: product.Product,
		Stock:   product.Stock,
	}
}

func toEntities(p []Product) []*model.Product {
	products := make([]*model.Product, len(p))
	for i, product := range p {
		products[i] = product.toEntity()
	}

	return products
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, product model.Product) (*model.Product, error) {
	startTime := time.Now()

	data := pr.fromEntity(product)

	res, err := pr.database.Collection("product").InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	data.ID = res.InsertedID.(primitive.ObjectID)

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	log.Printf("Execution Time (Insert New Product): %s\n", executionTime)
	return data.toEntity(), nil
}

func (pr *ProductRepository) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	startTime := time.Now()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objId,
	}

	var product Product
	err = pr.database.Collection("product").FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	log.Printf("Execution Time (Insert New Product) : %s", executionTime)
	return product.toEntity(), nil
}

func (pr *ProductRepository) GetProducts(ctx context.Context) ([]*model.Product, error) {
	starTime := time.Now()

	var products []Product
	res, err := pr.database.Collection("product").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleProduct model.Product
		if err := res.Decode(&singleProduct); err != nil {
			return nil, err
		}

		log.Println(singleProduct)

		products = append(products, Product(singleProduct))
	}

	endTime := time.Now()
	exe := endTime.Sub(starTime)
	log.Printf("Execution Time (Insert New Product): %s\n", exe)

	return toEntities(products), nil
}

func (pr *ProductRepository) UpdateProduct(ctx context.Context, id string, product model.Product) (*model.Product, error) {
	startTime := time.Now()

	data := pr.fromEntity(product)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objId,
	}

	update := bson.M{
		"$set": data,
	}

	res, err := pr.database.Collection("product").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		log.Printf("Product Not Found: %s\n", id)
		return nil, err
	}

	var updatedProduct Product
	if res.MatchedCount == 1 {
		err := pr.database.Collection("product").FindOne(ctx, filter).Decode(&updatedProduct)
		if err != nil {
			return nil, err
		}
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	log.Printf("Execution Time (Update Product): %s\n", executionTime)

	return updatedProduct.toEntity(), nil
}

func (pr *ProductRepository) DeleteProduct(ctx context.Context, id string) (*model.Product, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objId,
	}

	var product Product
	if err := pr.database.Collection("product").FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}

	res, err := pr.database.Collection("product").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	if res.DeletedCount == 0 {
		log.Printf("Product Not Found: %s\n", id)
		return nil, err
	}
	return product.toEntity(), nil
}
