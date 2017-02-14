package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/graphql-go/graphql"
	"github.com/mnmtanish/go-graphiql"
	"github.com/NiciiA/GoGraphQL/domain/type/entityType"
	"github.com/NiciiA/GoGraphQL/domain/model/entityModel"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/type/groupType"
	"github.com/NiciiA/GoGraphQL/domain/model/groupModel"
	"github.com/NiciiA/GoGraphQL/domain/type/priorityType"
	"github.com/NiciiA/GoGraphQL/domain/type/categoryType"
	"github.com/NiciiA/GoGraphQL/domain/type/tagType"
	"github.com/NiciiA/GoGraphQL/domain/type/contactType"
	"github.com/NiciiA/GoGraphQL/domain/type/newsType"
	"github.com/NiciiA/GoGraphQL/domain/type/orgUnitType"
	"github.com/NiciiA/GoGraphQL/dataaccess/entityDao"
	"github.com/NiciiA/GoGraphQL/domain/type/activityType"
	"github.com/NiciiA/GoGraphQL/domain/type/fileType"
	"github.com/NiciiA/GoGraphQL/domain/type/jwtType"
	"github.com/NiciiA/GoGraphQL/domain/model/accountModel"
	"github.com/NiciiA/GoGraphQL/dataaccess/categoryDao"
	"strconv"
	"github.com/NiciiA/GoGraphQL/dataaccess/accountDao"
	"github.com/NiciiA/GoGraphQL/domain/model/categoryModel"
	"time"
)

var (
	Schema graphql.Schema
)

/*
	News TODO - REST Client
	Permission / Roles

	And Account Managment with REST

 */

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func init() {
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"auth": &graphql.Field{
				Type: jwtType.Type,
				Args: graphql.FieldConfigArgument{
					"auth": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					auth, _ := p.Args["auth"].(string)
					password, _ := p.Args["password"].(string)
					var account = accountModel.Account{}
					account.Password = password
					if validateEmail(auth) {
						account.EMail = auth
						accountDao.GetAll(account).One(&account)
						return account, nil
					} else if len(auth) == 10 {
						if _, err := strconv.Atoi(auth); err == nil {
							account.Phone = auth
							accountDao.GetAll(account).One(&account)
							return account, nil
						}
					}

					return nil, nil
				},
			},
			"register": &graphql.Field{
				Type: jwtType.Type,
				Args: graphql.FieldConfigArgument{
					"auth": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					auth, _ := p.Args["auth"].(string)
					password, _ := p.Args["password"].(string)
					var account accountModel.Account = accountModel.Account{}
					account.ID = bson.NewObjectId()
					account.Password = password
					if validateEmail(auth) {
						account.EMail = auth
						accountDao.Insert(account)
						return account, nil
					} else if len(auth) == 10 {
						if _, err := strconv.Atoi(auth); err == nil {
							account.Phone = auth
							accountDao.Insert(account)
							return account, nil
						}
					}
					return nil, nil
				},
			},
			"createCategory": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					category := categoryModel.Category{}
					category.ID = bson.NewObjectId()
					category.Disabled = false
					category.CreatedDate = time.Now().Format(time.RFC3339)
					category.ModifiedDate = time.Now().Format(time.RFC3339)
					category.Name = name
					categoryDao.AddCategory(category)
					return  category, nil
				},
			},
			"updateCategory": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					typeCat, _ := p.Args["type"].(string)
					category := categoryModel.Category{}
					category.Name = name
					category.Type = typeCat
					categoryDao.UpdateCategory(category)
					return  category, nil
				},
			},
			"removeCategory": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					category := categoryModel.Category{}
					category.Name = name
					categoryDao.Delete(category)
					return category, nil
				},
			},
			"disableCategory": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					disable, _ := p.Args["disable"].(bool)
					category := categoryModel.Category{}
					category.Name = name
					category.Disabled = disable
					categoryDao.UpdateCategory(category)
					return  category, nil
				},
			},
			"createContact": &graphql.Field{
				Type: contactType.Type,
			},
			"updateContact": &graphql.Field{
				Type: contactType.Type,
			},
			"removeContact": &graphql.Field{
				Type: contactType.Type,
			},
			"disableContact": &graphql.Field{
				Type: contactType.Type,
			},
			"createEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"updateEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"removeEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"disableEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"pushEntityFile": &graphql.Field{
				Type: fileType.Type,
			},
			"removeEntityFile": &graphql.Field{
				Type: fileType.Type,
			},
			"pushEntityActivity": &graphql.Field{
				Type: activityType.Type,
			},
			"removeEntityActivity": &graphql.Field{
				Type: activityType.Type,
			},
			"createGroup": &graphql.Field{
				Type: groupType.Type,
			},
			"updateGroup": &graphql.Field{
				Type: groupType.Type,
			},
			"removeGroup": &graphql.Field{
				Type: groupType.Type,
			},
			"disableGroup": &graphql.Field{
				Type: groupType.Type,
			},
			"createNews": &graphql.Field{
				Type: newsType.Type,
			},
			"updateNews": &graphql.Field{
				Type: newsType.Type,
			},
			"removeNews": &graphql.Field{
				Type: newsType.Type,
			},
			"disableNews": &graphql.Field{
				Type: newsType.Type,
			},
			"pushNewsFile": &graphql.Field{
				Type: fileType.Type,
			},
			"removeNewsFile": &graphql.Field{
				Type: fileType.Type,
			},
			"removeNewsComment": &graphql.Field{
				Type: activityType.Type,
			},
			"pushNewsComment": &graphql.Field{
				Type: activityType.Type,
			},
			"createOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
			},
			"updateOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
			},
			"removeOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
			},
			"disableOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
			},
			"createPriority": &graphql.Field{
				Type: priorityType.Type,
			},
			"updatePriority": &graphql.Field{
				Type: priorityType.Type,
			},
			"removePriority": &graphql.Field{
				Type: priorityType.Type,
			},
			"disablePriority": &graphql.Field{
				Type: priorityType.Type,
			},
			"createTag": &graphql.Field{
				Type: tagType.Type,
			},
			"updateTag": &graphql.Field{
				Type: tagType.Type,
			},
			"removeTag": &graphql.Field{
				Type: tagType.Type,
			},
			"disableTag": &graphql.Field{
				Type: tagType.Type,
			},
		},
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"categoryList": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// type == news oder entity
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return categoryDao.GetAll(), nil
				},
			},
			"contactList": &graphql.Field{
				Type: contactType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"entityList": &graphql.Field{
				Type: graphql.NewList(entityType.Type),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					entityList := []entityModel.Entity{}
					entityDao.GetAll(bson.M{}).All(&entityList)
					return entityList, nil
					// curl -g 'http://localhost:8080/graphql?query={entityList{id}}'
					// return entityModel.Entity{CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"groupList": &graphql.Field{
				Type: groupType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"newsList": &graphql.Field{
				Type: newsType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"orgUnitList": &graphql.Field{
				Type: orgUnitType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"priorityList": &graphql.Field{
				Type: priorityType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"tagList": &graphql.Field{
				Type: tagType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"category": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					catName, _ := p.Args["name"].(string)
					// curl -g 'http://localhost:8080/graphql?query={category(name:"catName"){name}}'
					return categoryDao.GetByKey(catName), nil
				},
			},
			"contact": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"entity": &graphql.Field{
				Type: entityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					// return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					entity := entityModel.Entity{}
					entityDao.GetById(bson.ObjectIdHex(idQuery)).One(&entity)
					return entity, nil
				},
			},
			"group": &graphql.Field{
				Type: groupType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"news": &graphql.Field{
				Type: newsType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"orgUnit": &graphql.Field{
				Type: orgUnitType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"priority": &graphql.Field{
				Type: priorityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"tag": &graphql.Field{
				Type: tagType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
		},
	})
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Mutation: mutationType,
	})
}

func main() {
	http.HandleFunc("/graphql", serveGraphQL(Schema))
	http.HandleFunc("/", graphiql.ServeGraphiQL)
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func serveGraphQL(s graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendError := func(err error) {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		req := &graphiql.Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			sendError(err)
			return
		}

		res := graphql.Do(graphql.Params{
			Schema:        s,
			RequestString: req.Query,
		})

		if err := json.NewEncoder(w).Encode(res); err != nil {
			sendError(err)
		}
	}
}
