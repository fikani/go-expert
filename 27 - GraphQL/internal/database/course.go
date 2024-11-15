package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	ID          string
	Name        string
	Description string
	CategoryId  string
	db          *sql.DB
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

// DDL
// CREATE TABLE courses ( id UUID PRIMARY KEY, name TEXT, description TEXT, category_id UUID, FOREIGN KEY (category_id) REFERENCES categories(id) );

func (c *Course) Create(name, description, categoryId string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryId)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryId:  categoryId,
	}, nil
}

func (c *Course) GetAll() ([]*Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := make([]*Course, 0)
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (c *Course) GetByCategoryId(id string) ([]*Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := make([]*Course, 0)
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}
