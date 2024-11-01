package db

import (
	"agendaAPIService/graph/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// db is de databaseverbinding die we in andere delen van de code zullen gebruiken
var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Ping de database om te controleren of de verbinding werkt
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

func CreateAgenda(owner int) (*model.Agenda, error) {
	var agenda model.Agenda
	err := db.QueryRow("INSERT INTO agendas (owner) VALUES ($1) RETURNING id", owner).Scan(&agenda.ID)
	if err != nil {
		return nil, err
	}
	agenda.Owner = owner
	return &agenda, nil
}

func GetAgendas() ([]*model.Agenda, error) {
	rows, err := db.Query("SELECT id, owner FROM agendas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []*model.Agenda
	for rows.Next() {
		var agenda model.Agenda
		if err := rows.Scan(&agenda.ID, &agenda.Owner); err != nil {
			return nil, err
		}
		agendas = append(agendas, &agenda)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return agendas, nil
}

func DeleteAgenda(id string) error {
	_, err := db.Exec("DELETE FROM agendas WHERE id = $1", id)
	return err
}

func UpdateAgenda(id string, owner *int) (*model.Agenda, error) {
	var agenda model.Agenda

	if owner != nil {
		err := db.QueryRow("UPDATE agendas SET owner = $1 WHERE id = $2 RETURNING id, owner", *owner, id).Scan(&agenda.ID, &agenda.Owner)
		if err != nil {
			return nil, fmt.Errorf("failed to update agenda: %v", err)
		}
	} else {
		return nil, fmt.Errorf("input.owner is required")
	}

	return &agenda, nil
}

func CreateAgendaItem(agendaID string, input model.CreateAgendaItem) (*model.AgendaItem, error) {
	var agendaItem model.AgendaItem

	if input.Date == nil {
		return nil, fmt.Errorf("date input is required")
	}

	if _, err := GetAgenda(agendaID); err != nil {
		return nil, fmt.Errorf("agenda with ID %s does not exist: %v", agendaID, err)
	}

	date, err := CreateOrGetDate(*input.Date)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow("INSERT INTO agenda_items (agenda_id, title, description, duration, date_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		agendaID, input.Title, input.Description, input.Duration, date.ID).Scan(&agendaItem.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create agenda item: %v", err)
	}

	agendaItem.Title = input.Title
	agendaItem.Description = input.Description
	agendaItem.Duration = input.Duration
	agendaItem.Date = date

	return &agendaItem, nil
}

func CreateOrGetDate(input model.DateInput) (*model.Date, error) {
	var date model.Date

	err := db.QueryRow("SELECT id FROM dates WHERE (day, month, year, hour, minute) = ($1, $2, $3, $4, $5)",
		input.Day, input.Month, input.Year, input.Hour, input.Minute).Scan(&date.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = db.QueryRow("INSERT INTO dates (day, month, year, hour, minute) VALUES ($1, $2, $3, $4, $5) RETURNING id",
				input.Day, input.Month, input.Year, input.Hour, input.Minute).Scan(&date.ID)
			if err != nil {
				return nil, fmt.Errorf("failed to insert date: %v", err)
			}
		} else {
			return nil, err
		}
	}

	date.Day = input.Day
	date.Month = input.Month
	date.Year = input.Year
	date.Hour = input.Hour
	date.Minute = input.Minute

	return &date, nil
}

func GetDate(dateID string) (*model.Date, error) {
	var date model.Date
	err := db.QueryRow("SELECT id, day, month, year, hour, minute FROM dates WHERE id = $1", dateID).
		Scan(&date.ID, &date.Day, &date.Month, &date.Year, &date.Hour, &date.Minute)
	if err != nil {
		return nil, err
	}
	return &date, nil
}

func UpdateAgendaItem(id string, input model.UpdateAgendaItem) (*model.AgendaItem, error) {
	var agendaItem model.AgendaItem

	if input.Date != nil {
		date, err := CreateOrGetDate(*input.Date)
		if err != nil {
			return nil, err
		}

		err = db.QueryRow("UPDATE agenda_items SET date_id = $1 WHERE id = $2 RETURNING id, title, description, duration",
			date.ID, id).Scan(&agendaItem.ID, &agendaItem.Title, &agendaItem.Description, &agendaItem.Duration)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Input.date is required")
	}

	return &agendaItem, nil
}

func DeleteAgendaItem(id string) error {
	_, err := db.Exec("DELETE FROM agenda_items WHERE id = $1", id)
	return err
}

func GetAgendaItems(agendaID string) ([]*model.AgendaItem, error) {
	rows, err := db.Query("SELECT id, title, description, duration, date_id FROM agenda_items WHERE agenda_id = $1", agendaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendaItems []*model.AgendaItem
	for rows.Next() {
		var agendaItem model.AgendaItem
		var dateID string

		if err := rows.Scan(&agendaItem.ID, &agendaItem.Title, &agendaItem.Description, &agendaItem.Duration, &dateID); err != nil {
			return nil, err
		}

		date, err := GetDate(dateID)
		if err != nil {
			return nil, err
		}

		agendaItem.Date = date
		agendaItems = append(agendaItems, &agendaItem)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return agendaItems, nil
}

func GetAgenda(id string) (*model.Agenda, error) {
	var agenda model.Agenda
	err := db.QueryRow("SELECT id, owner FROM agendas WHERE id = $1", id).
		Scan(&agenda.ID, &agenda.Owner)
	if err != nil {
		return nil, err
	}
	return &agenda, nil
}
