package repository

import (
	"database/sql"
	"errors"

	"github.com/gplus1/HouseRentalSystem/entity"
)

// CategoryRepositoryImpl implements the menu.CategoryRepository interface
type HouseCategoryRepositoryImpl struct {
	conn *sql.DB
}

// NewCategoryRepositoryImpl will create an object of PsqlCategoryRepository
func NewCategoryRepositoryImpl(Conn *sql.DB) *HouseCategoryRepositoryImpl {
	return &HouseCategoryRepositoryImpl{conn: Conn}
}

// Categories returns all cateogories from the database
func (HcRi *HouseCategoryRepositoryImpl) Categories() ([]entity.HouseCategory, error) {

	rows, err := HcRi.conn.Query("SELECT * FROM categories;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.HouseCatagory{}

	for rows.Next() {
		HouseCatagory := entity.HouseCatagory{}
		err = rows.Scan(&HouseCatagory.ID, &HouseCatagory.Name, &HouseCatagory.DesHcRiption, &HouseCatagory.Image)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, HouseCatagory)
	}

	return ctgs, nil
}

// HouseCatagory returns a HouseCatagory with a given id
func (HcRi *HouseCatagoryRepositoryImpl) HouseCatagory(id uint) (entity.HouseCatagory, error) {

	row := HcRi.conn.QueryRow("SELECT * FROM categories WHERE id = $1", id)

	c := entity.HouseCatagory{}

	err := row.Scan(&c.ID, &c.Name, &c.DesHcRiption, &c.Image)
	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateHouseCatagory updates a given object with a new data
func (HcRi *HouseCatagoryRepositoryImpl) UpdateCategory(c entity.HouseCategory) error {

	_, err := HcRi.conn.Exec("UPDATE categories SET name=$1,desHcRiption=$2, image=$3 WHERE id=$4", c.Name, c.DesHcRiption, c.Image, c.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteCategory removes a category from a database by its id
func (HcRi *HouseCategoryRepositoryImpl) DeleteCategory(id uint) error {

	_, err := HcRi.conn.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreCategory stores new category information to database
func (HcRi *HouseCategoryRepositoryImpl) StoreCategory(c entity.HouseCategory) error {

	_, err := HcRi.conn.Exec("INSERT INTO categories (name,desHcRiption,image) values($1, $2, $3)", c.Name, c.DesHcRiption, c.Image)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
