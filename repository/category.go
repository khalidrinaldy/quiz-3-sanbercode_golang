package repository

import (
	"database/sql"
	"quiz-3-sanbercode_golang/structs"
)

func GetAllCategories(db *sql.DB) (categories []structs.Category, err error) {
	sql := "select * from category"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for rows.Next() {
		var category structs.Category
		err := rows.Scan(category.ID, category.Name, category.CreatedAt, category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "insert into category (name) values ($1)"
	errs := db.QueryRow(sql, category.Name)
	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "update category set name=$1"
	errs := db.QueryRow(sql, category.Name)
	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "delete from category where id=$1"
	errs := db.QueryRow(sql, category.ID)
	return errs.Err()
}
