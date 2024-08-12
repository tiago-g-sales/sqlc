package main

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tiago-g-sales/sqlc/internal/db"
  _ "github.com/go-sql-driver/mysql"
)

func main(){
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil{
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	//createCategory(ctx, queries)
	//updateCategory(ctx, queries)
	//deleteCategory(ctx, queries,"e29646b2-d179-4ae7-8442-c6159b6daabf" )

	listCategories(ctx, queries)

}
func createCategory(ctx context.Context, queries *db.Queries){

	err := queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID: uuid.New().String(),
		Name: "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	} )

	if err != nil {
		panic(err)
	}
}

func listCategories(ctx context.Context, queries *db.Queries){
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}	
	for _,category := range categories{
		println(category.ID, category.Name, category.Description.String)
	}
}

func updateCategory(ctx context.Context, queries *db.Queries ){
	err := queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID: "9d7556c6-dd54-4be7-8cf1-5abef48ed074",
		Name: "Backend updated",
		Description: sql.NullString{String: "Backend description updated", Valid: true},
	}) 
	if err != nil {
		panic(err)
	}	

}


func deleteCategory(ctx context.Context, queries *db.Queries, id string ){
	err := queries.DeleteCategory(ctx, id)
	if err != nil {
		panic(err)
	}	
}
