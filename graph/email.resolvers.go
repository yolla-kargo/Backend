package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	// "fmt"

	"github.com/yolla-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SendTruckDataToEmail(ctx context.Context, email string) (*model.Mail, error) {
	mail := &model.Mail{
		Email: email,
	}
	r.Mails = append(r.Mails, mail)
	return mail, nil
}
