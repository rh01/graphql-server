//go:generate go run ../../../scripts/gqlgen.go
package resolver

import (
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/order"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/publication"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver/user"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
)

type Resolver struct {
	registry.Registry
	Logger *logrus.Entry
}

type MutationResolver struct {
	registry.Registry
	Logger *logrus.Entry
}

type QueryResolver struct {
	registry.Registry
	Logger *logrus.Entry
}

func NewResolver(nr registry.Registry, l *logrus.Entry) *Resolver {
	return &Resolver{Registry: nr, Logger: l}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &MutationResolver{
		Registry: r.Registry,
		Logger:   r.Logger,
	}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &QueryResolver{
		Registry: r.Registry,
		Logger:   r.Logger,
	}
}
func (r *Resolver) User() generated.UserResolver {
	return &user.UserResolver{
		Client: r.GetUserClient(registry.USER),
		Logger: r.Logger,
	}
}
func (r *Resolver) Role() generated.RoleResolver {
	return &user.RoleResolver{
		Client: r.GetRoleClient(registry.ROLE),
		Logger: r.Logger,
	}
}
func (r *Resolver) Permission() generated.PermissionResolver {
	return &user.PermissionResolver{
		Client: r.GetPermissionClient(registry.PERMISSION),
		Logger: r.Logger,
	}
}
func (r *Resolver) Publication() generated.PublicationResolver {
	return &publication.PublicationResolver{
		Logger: r.Logger,
	}
}
func (r *Resolver) Author() generated.AuthorResolver {
	return &publication.AuthorResolver{
		Logger: r.Logger,
	}
}
func (r *Resolver) Strain() generated.StrainResolver {
	return &stock.StrainResolver{
		Client:     r.GetStockClient(registry.STOCK),
		UserClient: r.GetUserClient(registry.USER),
		Registry:   r.Registry,
		Logger:     r.Logger,
	}
}
func (r *Resolver) Plasmid() generated.PlasmidResolver {
	return &stock.PlasmidResolver{
		Client:     r.GetStockClient(registry.STOCK),
		UserClient: r.GetUserClient(registry.USER),
		Registry:   r.Registry,
		Logger:     r.Logger,
	}
}

func (r *Resolver) Order() generated.OrderResolver {
	return &order.OrderResolver{
		Client:      r.GetOrderClient(registry.ORDER),
		StockClient: r.GetStockClient(registry.STOCK),
		UserClient:  r.GetUserClient(registry.USER),
		Logger:      r.Logger,
	}
}
