package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	db "agendaAPIService/database"
	"agendaAPIService/graph/model"
	"context"
)

// CreateAgenda is de resolver voor de createAgenda field.
func (r *mutationResolver) CreateAgenda(ctx context.Context, input model.CreateAgenda) (*model.Agenda, error) {
	agenda, err := db.CreateAgenda(input.Owner)
	if err != nil {
		return nil, err
	}
	return agenda, nil
}

// DeleteAgenda is the resolver for the deleteAgenda field.
func (r *mutationResolver) DeleteAgenda(ctx context.Context, id string) (bool, error) {
	err := db.DeleteAgenda(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateAgenda is the resolver for the updateAgenda field.
func (r *mutationResolver) UpdateAgenda(ctx context.Context, id string, owner *int) (*model.Agenda, error) {
	agenda, err := db.UpdateAgenda(id, owner)
	if err != nil {
		return nil, err
	}
	return agenda, nil
}

// CreateAgendaItem is the resolver for the createAgendaItem field.
func (r *mutationResolver) CreateAgendaItem(ctx context.Context, agendaID string, input model.CreateAgendaItem) (*model.AgendaItem, error) {
	agenda, err := db.CreateAgendaItem(agendaID, input)
	if err != nil {
		return nil, err
	}
	return agenda, nil
}

// UpdateAgendaItem is the resolver for the updateAgendaItem field.
func (r *mutationResolver) UpdateAgendaItem(ctx context.Context, id string, input model.UpdateAgendaItem) (*model.AgendaItem, error) {
	agenda, err := db.UpdateAgendaItem(id, input)
	if err != nil {
		return nil, err
	}
	return agenda, nil
}

// DeleteAgendaItem is the resolver for the deleteAgendaItem field.
func (r *mutationResolver) DeleteAgendaItem(ctx context.Context, id string) (bool, error) {
	err := db.DeleteAgendaItem(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Agendas is de resolver voor de agendas field.
func (r *queryResolver) Agendas(ctx context.Context) ([]*model.Agenda, error) {
	agendas, err := db.GetAgendas()
	if err != nil {
		return nil, err
	}
	return agendas, nil
}

// Agenda is the resolver for the agenda field.
func (r *queryResolver) Agenda(ctx context.Context, id string) (*model.Agenda, error) {
	agenda, err := db.GetAgenda(id)
	if err != nil {
		return nil, err
	}
	return agenda, nil
}

// AgendaItems is the resolver for the agendaItems field.
func (r *queryResolver) AgendaItems(ctx context.Context, agendaID string) ([]*model.AgendaItem, error) {
	agendaItems, err := db.GetAgendaItems(agendaID)
	if err != nil {
		return nil, err
	}
	return agendaItems, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
