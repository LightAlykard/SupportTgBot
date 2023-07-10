package info

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
	Data string
	//Permissions int
}

type Deal struct {
	ID     uuid.UUID
	IDUser uuid.UUID
	Data   string
	//Permissions int
}

type InfoStore interface {
	//Create(ctx context.Context, u User) (*uuid.UUID, error)
	ReadUser(ctx context.Context, uid uuid.UUID) (*User, error)
	ReadDeal(ctx context.Context, did uuid.UUID) (*Deal, error)
	//Delete(ctx context.Context, uid uuid.UUID) error
	//SearchUsers(ctx context.Context, s string) (chan User, error)
}

type Infos struct {
	istore InfoStore
}

func NewUsers(istore InfoStore) *Infos {
	return &Infos{
		istore: istore,
	}
}

// func (us *Users) Create(ctx context.Context, u User) (*User, error) {
// 	id, err := us.ustore.Create(ctx, u)
// 	if err != nil {
// 		return nil, fmt.Errorf("create user error: %w", err)
// 	}
// 	u.ID = *id
// 	return &u, nil
// }

func (us *Infos) ReadUser(ctx context.Context, uid uuid.UUID) (*User, error) {
	u, err := us.istore.ReadUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("read user error: %w", err)
	}
	return u, nil
}

func (us *Infos) ReadDeal(ctx context.Context, uid uuid.UUID) (*Deal, error) {
	u, err := us.istore.ReadDeal(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("read user error: %w", err)
	}
	return u, nil
}

// func (us *Users) Delete(ctx context.Context, uid uuid.UUID) (*User, error) {
// 	u, err := us.ustore.Read(ctx, uid)
// 	if err != nil {
// 		return nil, fmt.Errorf("search user error: %w", err)
// 	}
// 	return u, us.ustore.Delete(ctx, uid)
// }
