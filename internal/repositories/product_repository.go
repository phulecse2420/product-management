package repositories

import (
	"database/sql"
	"pm/internal/models"
)

type ProductRepository struct{ db *sql.DB }

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(inp models.CreateProductInput) (*models.Product, error) {
	p := &models.Product{}
	err := r.db.QueryRow(
		`INSERT INTO products (name,description,price,quantity)
         VALUES ($1,$2,$3,$4)
         RETURNING id,name,description,price,quantity,created_at,updated_at`,
		inp.Name, inp.Description, inp.Price, inp.Quantity,
	).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
	return p, err
}

func (r *ProductRepository) List(keyword string) ([]models.Product, error) {
	q := `SELECT id,name,description,price,quantity,created_at,updated_at
          FROM products WHERE ($1='' OR name ILIKE '%'||$1||'%') ORDER BY id`
	rows, err := r.db.Query(q, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Product
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
		list = append(list, p)
	}
	return list, nil
}

func (r *ProductRepository) GetByID(id int64) (*models.Product, error) {
	p := &models.Product{}
	err := r.db.QueryRow(
		`SELECT id,name,description,price,quantity,created_at,updated_at FROM products WHERE id=$1`, id,
	).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return p, err
}

func (r *ProductRepository) Update(id int64, inp models.UpdateProductInput) (*models.Product, error) {
	p := &models.Product{}
	err := r.db.QueryRow(
		`UPDATE products SET name=$1,description=$2,price=$3,quantity=$4,updated_at=NOW()
         WHERE id=$5 RETURNING id,name,description,price,quantity,created_at,updated_at`,
		inp.Name, inp.Description, inp.Price, inp.Quantity, id,
	).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return p, err
}

func (r *ProductRepository) Delete(id int64) (bool, error) {
	res, err := r.db.Exec(`DELETE FROM products WHERE id=$1`, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}
