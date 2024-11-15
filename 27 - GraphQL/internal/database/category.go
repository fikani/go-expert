package database

import (
	"database/sql"
	// uuid pkg form google
	"github.com/google/uuid"
)

type Category struct {
	ID          string
	Name        string
	Description string
	db          *sql.DB
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name, description string) (*Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return nil, err
	}
	return &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) GetAll() ([]*Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]*Category, 0)
	for rows.Next() {
		category := &Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) GetById(id string) *Category {
	row := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id)
	category := &Category{}
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil
	}
	return category
}

func (c *Category) FindByCourseId(courseId string) (*Category, error) {
	row := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseId)
	category := &Category{}
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}
	return category, nil
}
