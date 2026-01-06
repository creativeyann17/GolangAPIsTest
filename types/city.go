package types

import "github.com/google/uuid"

// CityDAO is used by pgx.RowToStructByName to map SQL query results
type CityDAO struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Population int64     `db:"population"`
}

// City is the API response model with JSON annotations
type City struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Population int64  `json:"population"`
}

// ToCity converts CityDAO to City for API responses
func (dao *CityDAO) ToCity() City {
	return City{
		ID:         dao.ID.String(),
		Name:       dao.Name,
		Population: dao.Population,
	}
}

// ToCityDAO converts City to CityDAO for database operations
func (c *City) ToCityDAO() CityDAO {
	id, _ := uuid.Parse(c.ID)
	return CityDAO{
		ID:         id,
		Name:       c.Name,
		Population: c.Population,
	}
}
