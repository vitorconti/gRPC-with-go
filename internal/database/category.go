package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()
	insertString := "INSERT INTO categories (id,name,description) VALUES ($1,$2,$3)"
	_, err := c.db.Exec(insertString, id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	queryString := "SELECT * FROM categories"
	resultSet, err := c.db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()

	categories := []Category{}
	for resultSet.Next() {
		var id, name, description string
		if err := resultSet.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, nil
}
func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseID).
		Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindByCategoryId(CategoryId string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT * FROM categories WHERE id = $1", CategoryId).
		Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}
