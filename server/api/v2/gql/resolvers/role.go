package resolvers

import (
	"context"
	"fmt"

	"github.com/SevenTV/ServerGo/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type roleResolver struct {
	ctx context.Context
	v   *mongo.Role

	fields map[string]*SelectedField
}

func GenerateRoleResolver(ctx context.Context, pRole *mongo.Role, roleID *primitive.ObjectID, fields map[string]*SelectedField) (*roleResolver, error) {
	if pRole != nil {
		return &roleResolver{
			ctx:    ctx,
			v:      pRole,
			fields: fields,
		}, nil
	}
	if roleID == nil {
		return nil, nil
	}

	role := mongo.GetRole(roleID)
	r := &roleResolver{
		ctx:    ctx,
		v:      &role,
		fields: fields,
	}
	return r, nil
}

func (r *roleResolver) ID() string {
	return r.v.ID.Hex()
}

func (r *roleResolver) Name() string {
	return r.v.Name
}

func (r *roleResolver) Position() int32 {
	return r.v.Position
}

func (r *roleResolver) Color() int32 {
	return r.v.Color
}

func (r *roleResolver) Allowed() string {
	return fmt.Sprint(r.v.Allowed)
}

func (r *roleResolver) Denied() string {
	return fmt.Sprint(r.v.Denied)
}
