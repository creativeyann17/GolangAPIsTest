package repository

import (
  "context"
  "fmt"

  "github.com/creativeyann17/GolangAPIsTest/types"
  "github.com/google/uuid"
  "github.com/jackc/pgx/v5"
  "github.com/jackc/pgx/v5/pgxpool"
)

type CityRepository struct {
  pool *pgxpool.Pool
}

func NewCityRepository(pool *pgxpool.Pool) *CityRepository {
  return &CityRepository{pool: pool}
}

func (r *CityRepository) Create(ctx context.Context, city *types.CityDAO) error {
  if city.ID == uuid.Nil {
    city.ID = uuid.New()
  }

  existing, err := r.GetByName(ctx, city.Name)
  if err != nil {
    return fmt.Errorf("failed to check existing city: %w", err)
  }
  if existing != nil {
    return fmt.Errorf("city with name %s already exists", city.Name)
  }

  query := `INSERT INTO city (id, name, population) VALUES ($1, $2, $3)`
  _, err = r.pool.Exec(ctx, query, city.ID, city.Name, city.Population)
  if err != nil {
    return fmt.Errorf("failed to create city: %w", err)
  }
  return nil
}

func (r *CityRepository) GetByID(ctx context.Context, id uuid.UUID) (*types.CityDAO, error) {
  query := `SELECT id, name, population FROM city WHERE id = $1`
  rows, err := r.pool.Query(ctx, query, id)
  if err != nil {
    return nil, fmt.Errorf("failed to query city: %w", err)
  }
  defer rows.Close()

  city, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[types.CityDAO])
  if err != nil {
    if err == pgx.ErrNoRows {
      return nil, nil
    }
    return nil, fmt.Errorf("failed to scan city: %w", err)
  }
  return &city, nil
}

func (r *CityRepository) GetAll(ctx context.Context) ([]types.CityDAO, error) {
  query := `SELECT id, name, population FROM city ORDER BY name`
  rows, err := r.pool.Query(ctx, query)
  if err != nil {
    return nil, fmt.Errorf("failed to query cities: %w", err)
  }
  defer rows.Close()

  cities, err := pgx.CollectRows(rows, pgx.RowToStructByName[types.CityDAO])
  if err != nil {
    return nil, fmt.Errorf("failed to collect cities: %w", err)
  }
  return cities, nil
}

func (r *CityRepository) Update(ctx context.Context, city *types.CityDAO) error {
  query := `UPDATE city SET name = $2, population = $3 WHERE id = $1`
  result, err := r.pool.Exec(ctx, query, city.ID, city.Name, city.Population)
  if err != nil {
    return fmt.Errorf("failed to update city: %w", err)
  }
  if result.RowsAffected() == 0 {
    return fmt.Errorf("city not found")
  }
  return nil
}

func (r *CityRepository) Delete(ctx context.Context, id uuid.UUID) error {
  query := `DELETE FROM city WHERE id = $1`
  result, err := r.pool.Exec(ctx, query, id)
  if err != nil {
    return fmt.Errorf("failed to delete city: %w", err)
  }
  if result.RowsAffected() == 0 {
    return fmt.Errorf("city not found")
  }
  return nil
}

func (r *CityRepository) GetByName(ctx context.Context, name string) (*types.CityDAO, error) {
  query := `SELECT id, name, population FROM city WHERE name = $1`
  rows, err := r.pool.Query(ctx, query, name)
  if err != nil {
    return nil, fmt.Errorf("failed to query city by name: %w", err)
  }
  defer rows.Close()

  city, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[types.CityDAO])
  if err != nil {
    if err == pgx.ErrNoRows {
      return nil, nil
    }
    return nil, fmt.Errorf("failed to scan city: %w", err)
  }
  return &city, nil
}
