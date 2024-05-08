// volunteers.go
package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Volunteer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	Description string `json:"description"`
}

type VolunteerModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (v VolunteerModel) Insert(volunteer *Volunteer) error {
	query := `
		INSERT INTO Volunteers (Name, Age, Description) 
		VALUES ($1, $2, $3) 
		RETURNING id, Name
	`
	args := []interface{}{volunteer.Name, volunteer.Age, volunteer.Description}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return v.DB.QueryRowContext(ctx, query, args...).Scan(&volunteer.ID, &volunteer.Name)
}

func (v VolunteerModel) GetVolunteer(id string) (*Volunteer, error) {
	query := `
		SELECT id, Name, Age, Description
		FROM Volunteers
		WHERE ID = $1
	`
	var volunteer Volunteer
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := v.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&volunteer.ID, &volunteer.Name, &volunteer.Age, &volunteer.Description)
	if err != nil {
		return nil, err
	}
	return &volunteer, nil
}

func (v VolunteerModel) UpdateVolunteer(id string, volunteer *Volunteer) error {
	query := `
		UPDATE Volunteers
		SET Name = $2, Age = $3, Description = $4
		WHERE ID = $1
		RETURNING ID
	`
	args := []interface{}{id, volunteer.Name, volunteer.Age, volunteer.Description}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return v.DB.QueryRowContext(ctx, query, args...).Scan(&volunteer.ID)
}

func (v VolunteerModel) DeleteVolunteer(id string) error {
	query := `
		DELETE FROM Volunteers
		WHERE ID = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := v.DB.ExecContext(ctx, query, id)
	return err
}
