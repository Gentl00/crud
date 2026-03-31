package storage

import (
	"context"
	"fmt"


	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct{
	pool *pgxpool.Pool
	config Configer
}

 type Contact struct{
	ID int
	Nom string
	Prenom string
	Email string
	Numero string
}

type ContactDTO struct{
	ID int
	Nom string
	Prenom string
	Email string
	Numero string
}

type Configer interface {
	EnvConfig()(config *pgxpool.Config, err error)
}

func toDTO(c Contact) ContactDTO{
	return ContactDTO{
		ID: c.ID,
		Nom: c.Nom,
		Prenom: c.Prenom,
		Email: c.Email,
		Numero: c.Numero,
	}
}

func NewStorage(conf Configer) *Storage{
	return &Storage{
		config: conf,
	}
}

func (s *Storage) NewPool(ctx context.Context) error{
	conf, err := s.config.EnvConfig()
	if err != nil {
		return err
	}
	s.pool, err = pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return err
	}
	if err = s.pool.Ping(ctx); err != nil{
		return err
	}
	if err = s.CreateTable(ctx); err != nil{
		return err
	}

	return nil
}

func (s *Storage) close(){
	s.pool.Close()
}

func (s *Storage) CreateTable(ctx context.Context) error{
	query := `
	CREATE TABLE IF NOT EXISTS contacts  (
	id SERIAL PRIMARY KEY,
	nom TEXT NOT NULL,
	prénom TEXT ,
	email TEXT,
	contact TEXT);
	`
	_, err := s.pool.Exec(ctx, query)
	return err
}

func (s *Storage) AddContactName(ctx context.Context, name string) error{
	_, err := s.pool.Exec(ctx, "INSERT INTO contacts (nom) VALUES ($1)", name)
	return err
}


func (s *Storage) UpdateEmail(ctx context.Context, id int, email string) error{
	query := `
	UPDATE contacts
	SET email = $1
	WHERE id = $2;
	`
	cmTag, err := s.pool.Exec(ctx, query, email, id)
	if err != nil {
		return err
	}
	if cmTag.RowsAffected() == 0 {
		return fmt.Errorf("aucun contact avec l'id %d", id)
	}
	return nil
}

func (s *Storage) UpdateName(ctx context.Context, id int, name string) error{
	query := `
	UPDATE contacts
	SET nom = $1
	WHERE id = $2;
	`
	cmTag, err := s.pool.Exec(ctx, query, name, id)
	if err != nil {
		return err
	}
	if cmTag.RowsAffected() == 0 {
		return fmt.Errorf("aucun contact avec l'id %d", id)
	}
	return nil
}

func (s *Storage) UpdateFirstName(ctx context.Context, id int, firstname string) error{
	query := `
	UPDATE contacts
	SET prénom = $1
	WHERE id = $2;
	`
	cmTag, err := s.pool.Exec(ctx, query, firstname, id)
	if err != nil {
		return err
	}
	if cmTag.RowsAffected() == 0 {
		return fmt.Errorf("aucun contact avec l'id %d", id)
	}
	return nil
}

func (s *Storage) UpdateContact(ctx context.Context, id int, contact string) error{
	query := `
	UPDATE contacts
	SET contact = $1
	WHERE id = $2;
	`
	cmTag, err := s.pool.Exec(ctx, query, contact, id)
	if err != nil {
		return err
	}
	if cmTag.RowsAffected() == 0 {
		return fmt.Errorf("aucun contact avec l'id %d", id)
	}
	return nil
}

func (s *Storage) DeleteContact(ctx context.Context, id int) error {
	cmTag, err := s.pool.Exec(ctx, "DELETE FROM contacts WHERE id = $1", id)
	if err != nil {
		return err
	}

	if cmTag.RowsAffected() == 0 {
		return fmt.Errorf("aucun contact avec l'id %d", id)
	}
	return nil
}

func (s *Storage) ListContact(ctx context.Context) ([]ContactDTO, error){
	query := `
	SELECT 
	id,
	COALESCE(nom, 'no_name') AS nom,
	COALESCE(prénom, 'no_lastname') AS prénom,
	COALESCE(email, 'no_email') AS email,
	COALESCE(contact, 'no_phone') AS contact
	FROM contacts ORDER BY id
	`
	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var contacts []Contact
	for rows.Next(){
		var c Contact
		if err := rows.Scan(&c.ID, &c.Nom, &c.Prenom, &c.Email, &c.Numero); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	contactsDTO := make([]ContactDTO, len(contacts))
	for i, v := range contacts {
		contactsDTO[i] = toDTO(v)
	}
	return contactsDTO, nil
}

