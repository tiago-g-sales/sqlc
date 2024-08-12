package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/tiago-g-sales/sqlc/internal/db"
)

type CourseDB struct{
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB{
	return &CourseDB{
		dbConn: dbConn,
		Queries: db.New(dbConn),
	}

}

type CourseParams struct{
	ID 			string
	Name 		string
	Description sql.NullString
	Price 		float64
}

type CategoryParams struct{ 
	ID 			string
	Name 		string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error{

	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil{
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil{
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}



func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams )  error{

	err := c.callTx(ctx, func(q *db.Queries) error{	
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID: argsCategory.ID,
			Name: argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil{
			return err
		}	
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID: argsCourse.ID,
			Name: argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID: argsCategory.ID,
			Price: argsCourse.Price,
		})
		if err != nil{
			return err
		}	
		return nil
	})
	if err != nil{
		return err
	}	
	return nil
}

func main(){
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil{
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)


	courses, err := queries.ListCourses(ctx)
	if err != nil{
		panic(err)	
	}
	for _, couse := range courses{
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f \n",
				couse.CategoryName, couse.ID, couse.Name, couse.Description.String, couse.Price )
	}


	// courseArgs := CourseParams{
	// 	ID: uuid.New().String(),
	// 	Name: "Go",
	// 	Description: sql.NullString{String: "Go Course", Valid: true},
	// 	Price: 10.95,
	// }
	// categoryArgs := CategoryParams{
	// 	ID: uuid.New().String(),
	// 	Name: "Backend",
	// 	Description: sql.NullString{String: "Backend Course", Valid: true},
	// }

	//courseDB := NewCourseDB(dbConn)
	//err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	//if err != nil{
	//	panic(err)
	//}	

	//createCategory(ctx, queries)
	//updateCategory(ctx, queries)
	//deleteCategory(ctx, queries,"e29646b2-d179-4ae7-8442-c6159b6daabf" )

	//listCategories(ctx, queries)

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
