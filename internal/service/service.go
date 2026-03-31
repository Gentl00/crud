package service

import (
	"context"
	"crud/internal/storage"
	"crud/internal/ui"
	"strings"
)
type Contact any

type ContactStorer interface{
	AddContactName(ctx context.Context, name string) error
	UpdateEmail(ctx context.Context, id int, email string) error
	UpdateContact(ctx context.Context, id int, contact string) error
	UpdateName(ctx context.Context, id int, name string) error
	UpdateFirstName(ctx context.Context, id int, firstname string) error
	DeleteContact(ctx context.Context, id int) error
	ListContact(ctx context.Context) ([]storage.ContactDTO, error)
}

type ContactService struct {
	store ContactStorer
}

func NewContactService(store ContactStorer) *ContactService{
	return &ContactService{
		store: store,
	}
}

func (cs *ContactService) AddName(ctx context.Context, name string) error{
	err := cs.store.AddContactName(ctx, name)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) AddFirstName(ctx context.Context, firstname string) error{
	contacts, err := cs.store.ListContact(ctx)
	if err != nil{
		return err
	}
	id := contacts[len(contacts)-1].ID
	if err := cs.UpdateFirstName(ctx, id, firstname); err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) AddMail(ctx context.Context, email string) error{
	contacts, err := cs.store.ListContact(ctx)
	if err != nil{
		return err
	}
	id := contacts[len(contacts)-1].ID
	if err := cs.UpdateEmail(ctx, id, email); err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) AddPhone(ctx context.Context, phone string) error{
	contacts, err := cs.store.ListContact(ctx)
	if err != nil{
		return err
	}
	id := contacts[len(contacts)-1].ID
	if err := cs.UpdateContact(ctx, id, phone); err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) Delete(ctx context.Context, id int) error{
	err := cs.store.DeleteContact(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) UpdateEmail(ctx context.Context, id int, email string) error{
	err := cs.store.UpdateEmail(ctx, id, email)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) UpdateContact(ctx context.Context, id int, contact string) error{
	err := cs.store.UpdateContact(ctx, id, contact)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) UpdateName(ctx context.Context, id int, name string) error{
	err := cs.store.UpdateName(ctx, id, name)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) UpdateFirstName(ctx context.Context, id int, firstname string) error{
	err := cs.store.UpdateFirstName(ctx, id, firstname)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) List(ctx context.Context) error{
	contacts, err := cs.store.ListContact(ctx)
	if err != nil {
		return err
	}
	if err := ui.TaskStyle(contacts); err != nil {
		return err
	}
	return nil
}

func search(i int, name string, contacts []storage.ContactDTO, cont []storage.ContactDTO) ([]storage.ContactDTO, error){
	if i == len(contacts) {
		return cont, nil
	}
	if strings.Contains(strings.ToLower(contacts[i].Nom), strings.ToLower(name)){
		cont = append(cont, contacts[i])
	}
	i++
	return search(i, name, contacts, cont)
}

func (cs *ContactService) ListSearch(ctx context.Context, name string) error{
	contacts, err := cs.store.ListContact(ctx)
	if err != nil {
		return err
	}
	var cont []storage.ContactDTO
	result, err := search(0, name, contacts, cont)
	if err != nil {
		return err
	}
	if err := ui.TaskStyle(result); err != nil {
		return err
	}
	return nil
}