package permissionType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/model/permissionModel"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Permission",
	Description: "A Permission response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if permission, ok := p.Source.(permissionModel.Permissions); ok {
					return permission.ID.Hex(), nil
				}
				return nil, nil
			},
		},
		"disabled": &graphql.Field{
			Type: graphql.Boolean,
		},
		"createdDate": &graphql.Field{
			Type: graphql.String,
		},
		"modifiedDate": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"ticketPermission": &graphql.Field{
			Type: TicketType,
		},
		"newsPermission": &graphql.Field{
			Type: NewsType,
		},
		"usertaskPermission": &graphql.Field{
			Type: UsertaskType,
		},
		"generalPermission": &graphql.Field{
			Type: GeneralType,
		},
	},
})

var TicketType *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PermissionT",
	Description: "A PermissionT response",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type: graphql.Boolean,
		},
		"read": &graphql.Field{
			Type: graphql.Boolean,
		},
		"edit": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var NewsType *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PermissionN",
	Description: "A PermissionN response",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type: graphql.Boolean,
		},
		"read": &graphql.Field{
			Type: graphql.Boolean,
		},
		"edit": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var UsertaskType *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PermissionU",
	Description: "A PermissionU response",
	Fields: graphql.Fields{
		"edit": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var GeneralType *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PermissionG",
	Description: "A PermissionG response",
	Fields: graphql.Fields{
		"admin": &graphql.Field{
			Type: graphql.Boolean,
		},
		"internal": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})