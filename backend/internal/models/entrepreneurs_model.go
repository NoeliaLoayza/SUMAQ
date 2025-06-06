package models

import (
	"context"
	"fmt"
	"log"

	"github.com/cyarleque/sumaq/internal/database/interfaces"
)

type Entrepreneurs struct {
	Id    int
	Name  string
	Email string
	Url   string
	Role  string
}

type BodyEntrepreneurUpdate struct {
	Url string
}

type FilterEntrepreneurs struct {
	CategoryId     int
	EntrepreneurId int
}

type EntrepreneursModelInterface interface {
	GetEntrepreneurs(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs, querier interfaces.SQLQuerier) ([]*Entrepreneurs, error)
	UpdateUrlEntrepreneur(ctx context.Context, entrepreneurId int, bodyEntrepreneurUpdate BodyEntrepreneurUpdate, querier interfaces.SQLQuerier) error
	GetEntrepreneur(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs, querier interfaces.SQLQuerier) (*Entrepreneurs, error)
}
type entrepreneursModel struct {
	Db *interfaces.SQLConnInterface
}

func NewEntrepreneursModel() EntrepreneursModelInterface {
	return &entrepreneursModel{}
}

func (c *entrepreneursModel) GetEntrepreneurs(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs, querier interfaces.SQLQuerier) ([]*Entrepreneurs, error) {
	queryParams := []interface{}{
		filterEntrepreneurs.CategoryId,
	}

	query := `
		SELECT 
		u.id, 
		u.name, 
		u.email,
		u.url,
		u.role
		FROM users u
		INNER JOIN products p ON p.entrepreneur_id = u.id
		WHERE p.category_id = ?
		AND u.role IN ("entrepreneur", "admin")
		GROUP BY u.id
	`

	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL]", "Error prepare Query:", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, queryParams...)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	entrepreneurs := make([]*Entrepreneurs, 0)
	for rows.Next() {
		entrepreneur := Entrepreneurs{}
		err := rows.Scan(
			&entrepreneur.Id,
			&entrepreneur.Name,
			&entrepreneur.Email,
			&entrepreneur.Url,
			&entrepreneur.Role,
		)
		if err != nil {
			log.Println("Error Scan:", err.Error())
			return nil, err
		}
		entrepreneurs = append(entrepreneurs, &entrepreneur)
	}

	return entrepreneurs, nil
}

func (c *entrepreneursModel) GetEntrepreneur(ctx context.Context, filterEntrepreneurs FilterEntrepreneurs, querier interfaces.SQLQuerier) (*Entrepreneurs, error) {
	queryParams := []interface{}{
		filterEntrepreneurs.EntrepreneurId,
	}

	query := `
		SELECT 
		u.id, 
		u.name, 
		u.email,
		u.url,
		u.role
		FROM users u
		WHERE u.id = ?
		AND u.role IN ("entrepreneur", "admin")
	`

	rows := querier.QueryRowContext(ctx, query, queryParams...)

	entrepreneur := &Entrepreneurs{}
	err := rows.Scan(
		&entrepreneur.Id,
		&entrepreneur.Name,
		&entrepreneur.Email,
		&entrepreneur.Url,
		&entrepreneur.Role,
	)

	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return nil, err
	}

	return entrepreneur, nil
}

func (c *entrepreneursModel) UpdateUrlEntrepreneur(ctx context.Context, entrepreneurId int, bodyEntrepreneurUpdate BodyEntrepreneurUpdate, querier interfaces.SQLQuerier) error {
	queryParams := []interface{}{
		bodyEntrepreneurUpdate.Url,
		entrepreneurId,
	}

	query := `
		UPDATE users 
			SET url = ? 
		WHERE id = ?
	`

	stmt, err := querier.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[FATAL]", "Error prepare Query:", err.Error())
		return err
	}
	defer stmt.Close()

	r, err := stmt.ExecContext(ctx, queryParams...)
	if err != nil {
		log.Printf("[FATAL] Error executing query: %v", err)
		return err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Printf("Error getting affected rows: %v", err)
		return err
	}

	fmt.Printf("UpdateUrlEntrepreneur - Affected lines %d - RequestID %v \n", rowsAffected, ctx.Value("requestId"))
	return nil
}
