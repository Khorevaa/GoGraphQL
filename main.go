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
	"github.com/NiciiA/GoGraphQL/dataaccess/groupDao"
	"github.com/NiciiA/GoGraphQL/domain/type/permissionType"
	"github.com/NiciiA/GoGraphQL/domain/model/orgUnitModel"
	"github.com/NiciiA/GoGraphQL/dataaccess/orgUnitDao"
	"github.com/NiciiA/GoGraphQL/domain/model/priorityModel"
	"github.com/NiciiA/GoGraphQL/dataaccess/priorityDao"
	"github.com/NiciiA/GoGraphQL/domain/model/tagModel"
	"github.com/NiciiA/GoGraphQL/dataaccess/tagDao"
	"github.com/NiciiA/GoGraphQL/domain/model/contactModel"
	"github.com/NiciiA/GoGraphQL/dataaccess/contactDao"
	"github.com/NiciiA/GoGraphQL/domain/model/newsModel"
	"github.com/NiciiA/GoGraphQL/dataaccess/newsDao"
	"github.com/NiciiA/GoGraphQL/domain/model/activityModel"
	"github.com/NiciiA/GoGraphQL/dataaccess/entityActivityDao"
	"github.com/NiciiA/GoGraphQL/dataaccess/newsActivityDao"
	"github.com/NiciiA/GoGraphQL/domain/model/jwtModel"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/dgrijalva/jwt-go"
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
			"createAccount": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"userName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"eMail": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"roles": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"groups": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userName, _ := p.Args["userName"].(string)
					password, _ := p.Args["password"].(string)
					eMail, _ := p.Args["eMail"].(string)
					phone, _ := p.Args["phone"].(string)
					roles, _ := p.Args["roles"].([]string)
					groups, _ := p.Args["groups"].([]string)
					account := accountModel.Account{}
					account.ID = bson.NewObjectId()
					account.Disabled = false
					account.CreatedDate = time.Now().Format(time.RFC3339)
					account.ModifiedDate = time.Now().Format(time.RFC3339)
					account.UserName = userName
					account.Password = password
					account.EMail = eMail
					account.Phone = phone
					account.Roles = roles
					account.Groups = groups
					accountDao.Insert(account)
					return  account, nil
				},
			},
			"updateAccount": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"userName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"eMail": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"roles": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"groups": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					userName, _ := p.Args["userName"].(string)
					password, _ := p.Args["password"].(string)
					eMail, _ := p.Args["eMail"].(string)
					phone, _ := p.Args["phone"].(string)
					roles, _ := p.Args["roles"].([]string)
					groups, _ := p.Args["groups"].([]string)
					if !bson.IsObjectIdHex(id) {
						return nil, nil
					}
					account := accountModel.Account{}
					accountDao.GetById(bson.ObjectIdHex(id)).One(&account)
					account.ModifiedDate = time.Now().Format(time.RFC3339)
					account.UserName = userName
					account.Password = password
					account.EMail = eMail
					account.Phone = phone
					account.Roles = roles
					account.Groups = groups
					accountDao.Update(account)
					return  account, nil
				},
			},
			"removeAccount": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(id) {
						return nil, nil
					}
					account := accountModel.Account{}
					accountDao.GetById(bson.ObjectIdHex(id)).One(&account)
					accountDao.Delete(account)
					return  account, nil
				},
			},
			"disableAccount": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					disable, _ := p.Args["disable"].(bool)
					if !bson.IsObjectIdHex(id) {
						return nil, nil
					}
					account := accountModel.Account{}
					accountDao.GetById(bson.ObjectIdHex(id)).One(&account)
					account.ModifiedDate = time.Now().Format(time.RFC3339)
					account.Disabled = disable
					accountDao.Update(account)
					return  account, nil
				},
			},
			"createCategory": &graphql.Field{
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
					category.ID = bson.NewObjectId()
					category.Disabled = false
					category.CreatedDate = time.Now().Format(time.RFC3339)
					category.ModifiedDate = time.Now().Format(time.RFC3339)
					category.Name = name
					category.Type = typeCat
					categoryDao.AddCategory(category)
					return  category, nil
				},
			},
			"updateCategory": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					name, _ := p.Args["name"].(string)
					typeCat, _ := p.Args["type"].(string)
					category := categoryDao.GetByKey(bson.ObjectIdHex(idQuery))
					category.ModifiedDate = time.Now().Format(time.RFC3339)
					category.Name = name
					category.Type = typeCat
					categoryDao.UpdateCategory(category)
					return  category, nil
				},
			},
			"removeCategory": &graphql.Field{
				Type: categoryType.Type,
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
					category := categoryDao.GetByKey(bson.ObjectIdHex(idQuery))
					categoryDao.Delete(category)
					return category, nil
				},
			},
			"disableCategory": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					disable, _ := p.Args["disable"].(bool)
					category := categoryDao.GetByKey(bson.ObjectIdHex(idQuery))
					category.ModifiedDate = time.Now().Format(time.RFC3339)
					category.Disabled = disable
					categoryDao.UpdateCategory(category)
					return  category, nil
				},
			},
			"createContact": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"street": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"village": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"orgUnit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"accounts": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					firstName, _ := p.Args["firstName"].(string)
					lastName, _ := p.Args["lastName"].(string)
					street, _ := p.Args["street"].(string)
					village, _ := p.Args["village"].(string)
					orgUnit, _ := p.Args["orgUnit"].(string)
					accounts, _ := p.Args["accounts"].([]string)
					contact := contactModel.Contact{}
					contact.ID = bson.NewObjectId()
					contact.Disabled = false
					contact.CreatedDate = time.Now().Format(time.RFC3339)
					contact.ModifiedDate = time.Now().Format(time.RFC3339)
					contact.FirstName = firstName
					contact.LastName = lastName
					contact.Street = street
					contact.Village = village
					contact.OrgUnit = orgUnit
					contact.Accounts = accounts
					contactDao.Insert(contact)
					return  contact, nil
				},
			},
			"updateContact": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"street": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"village": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"orgUnit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"accounts": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					firstName, _ := p.Args["firstName"].(string)
					lastName, _ := p.Args["lastName"].(string)
					street, _ := p.Args["street"].(string)
					village, _ := p.Args["village"].(string)
					orgUnit, _ := p.Args["orgUnit"].(string)
					accounts, _ := p.Args["accounts"].([]string)
					if !bson.IsObjectIdHex(id) {
						return nil, nil
					}
					contact := contactModel.Contact{}
					contactDao.GetById(bson.ObjectIdHex(id)).One(&contact)
					contact.ModifiedDate = time.Now().Format(time.RFC3339)
					contact.FirstName = firstName
					contact.LastName = lastName
					contact.Street = street
					contact.Village = village
					contact.OrgUnit = orgUnit
					contact.Accounts = accounts
					contactDao.Update(contact)
					return  contact, nil
				},
			},
			"removeContact": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(id) {
						return nil, nil
					}
					contact := contactModel.Contact{}
					contactDao.GetById(bson.ObjectIdHex(id)).One(&contact)
					contactDao.Delete(contact)
					return  contact, nil
				},
			},
			"disableContact": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					disable, _ := p.Args["disable"].(bool)
					if !bson.IsObjectIdHex(id) {
						return nil, nil
					}
					contact := contactModel.Contact{}
					contactDao.GetById(bson.ObjectIdHex(id)).One(&contact)
					contact.ModifiedDate = time.Now().Format(time.RFC3339)
					contact.Disabled = disable
					contactDao.Update(contact)
					return  contact, nil
				},
			},
			"createEntity": &graphql.Field{
				Type: entityType.Type,
				Args: graphql.FieldConfigArgument{
					"subject": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"closed": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
					"tags": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"groups": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"priority": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"creator": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					subject, _ := p.Args["subject"].(string)
					description, _ := p.Args["description"].(string)
					longitude, _ := p.Args["longitude"].(string)
					latitude, _ := p.Args["latitude"].(string)
					closed, _ := p.Args["closed"].(bool)
					tags, _ := p.Args["tags"].([]string)
					groups, _ := p.Args["groups"].([]string)
					priority, _ := p.Args["priority"].(string)
					category, _ := p.Args["category"].(string)
					creator, _ := p.Args["creator"].(string)
					entity := entityModel.Entity{}
					entity.ID = bson.NewObjectId()
					entity.Disabled = false
					entity.CreatedDate = time.Now().Format(time.RFC3339)
					entity.ModifiedDate = time.Now().Format(time.RFC3339)
					entity.Subject = subject
					entity.Description = description
					entity.Longitude = longitude
					entity.Latitude = latitude
					entity.Closed = closed
					entity.Tags = tags
					entity.Groups = groups
					entity.Priority = priority
					entity.Category = category
					entity.CreatedDate = creator
					entityDao.Insert(entity)
					return  entity, nil
				},
			},
			"updateEntity": &graphql.Field{
				Type: entityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"subject": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"closed": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
					"tags": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"groups": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"priority": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"creator": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					subject, _ := p.Args["subject"].(string)
					description, _ := p.Args["description"].(string)
					longitude, _ := p.Args["longitude"].(string)
					latitude, _ := p.Args["latitude"].(string)
					closed, _ := p.Args["closed"].(bool)
					tags, _ := p.Args["tags"].([]string)
					groups, _ := p.Args["groups"].([]string)
					priority, _ := p.Args["priority"].(string)
					category, _ := p.Args["category"].(string)
					creator, _ := p.Args["creator"].(string)
					entity := entityModel.Entity{}
					entityDao.GetById(bson.ObjectIdHex(id)).One(&entity)
					entity.ModifiedDate = time.Now().Format(time.RFC3339)
					entity.Subject = subject
					entity.Description = description
					entity.Longitude = longitude
					entity.Latitude = latitude
					entity.Closed = closed
					entity.Tags = tags
					entity.Groups = groups
					entity.Priority = priority
					entity.Category = category
					entity.CreatedDate = creator
					entityDao.Insert(entity)
					return  entity, nil
				},
			},
			"removeEntity": &graphql.Field{
				Type: entityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					entity := entityModel.Entity{}
					entityDao.GetById(bson.ObjectIdHex(id)).One(&entity)
					entityDao.Delete(entity)
					return  entity, nil
				},
			},
			"disableEntity": &graphql.Field{
				Type: entityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					disable, _ := p.Args["disable"].(bool)
					entity := entityModel.Entity{}
					entityDao.GetById(bson.ObjectIdHex(id)).One(&entity)
					entity.ModifiedDate = time.Now().Format(time.RFC3339)
					entity.Disabled = disable
					entityDao.Update(entity)
					return  entity, nil
				},
			},
			"pushEntityFile": &graphql.Field{
				Type: fileType.Type,
			},
			"removeEntityFile": &graphql.Field{
				Type: fileType.Type,
			},
			"pushEntityActivity": &graphql.Field{
				Type: activityType.Type,
				Args: graphql.FieldConfigArgument{
					"referenceId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"comment": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"intern": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
					"creator": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					referenceId, _ := p.Args["referenceId"].(string)
					comment, _ := p.Args["comment"].(string)
					intern, _ := p.Args["intern"].(bool)
					creator, _ := p.Args["creator"].(string)
					activity := activityModel.Activity{}
					activity.ID = bson.NewObjectId()
					activity.Disabled = false
					activity.CreatedDate = time.Now().Format(time.RFC3339)
					activity.ModifiedDate = time.Now().Format(time.RFC3339)
					activity.ReferenceClass = "Entity"
					activity.ReferenceId = referenceId
					activity.Comment = comment
					activity.Intern = intern
					activity.Creator = creator
					entityActivityDao.Insert(activity)
					return  activity, nil
				},
			},
			"removeEntityActivity": &graphql.Field{
				Type: activityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					activity := activityModel.Activity{}
					entityActivityDao.GetById(bson.ObjectIdHex(id)).One(&activity)
					entityActivityDao.Delete(activity)
					return  activity, nil
				},
			},
			"createGroup": &graphql.Field{
				Type: groupType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					group := groupModel.Group{}
					group.ID = bson.NewObjectId()
					group.Disabled = false
					group.CreatedDate = time.Now().Format(time.RFC3339)
					group.ModifiedDate = time.Now().Format(time.RFC3339)
					group.Name = name
					groupDao.AddGroup(group)
					return  group, nil
				},
			},
			"updateGroup": &graphql.Field{
				Type: groupType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					name, _ := p.Args["name"].(string)
					group := groupDao.GetByKey(bson.ObjectIdHex(idQuery))
					group.ModifiedDate = time.Now().Format(time.RFC3339)
					group.Name = name
					groupDao.UpdateGroup(group)
					return  group, nil
				},
			},
			"removeGroup": &graphql.Field{
				Type: groupType.Type,
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
					group := groupDao.GetByKey(bson.ObjectIdHex(idQuery))
					groupDao.Delete(group)
					return group, nil
				},
			},
			"disableGroup": &graphql.Field{
				Type: groupType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					disable, _ := p.Args["disable"].(bool)
					group := groupDao.GetByKey(bson.ObjectIdHex(idQuery))
					group.ModifiedDate = time.Now().Format(time.RFC3339)
					group.Disabled = disable
					groupDao.UpdateGroup(group)
					return  group, nil
				},
			},
			"createNews": &graphql.Field{
				Type: newsType.Type,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"intern": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
					"tags": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"groups": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"important": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"creator": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					title, _ := p.Args["title"].(string)
					text, _ := p.Args["text"].(string)
					intern, _ := p.Args["intern"].(bool)
					tags, _ := p.Args["tags"].([]string)
					groups, _ := p.Args["groups"].([]string)
					important, _ := p.Args["important"].(bool)
					category, _ := p.Args["category"].(string)
					creator, _ := p.Args["creator"].(string)
					news := newsModel.News{}
					news.ID = bson.NewObjectId()
					news.Disabled = false
					news.CreatedDate = time.Now().Format(time.RFC3339)
					news.ModifiedDate = time.Now().Format(time.RFC3339)
					news.Title = title
					news.Text = text
					news.Intern = intern
					news.Tags = tags
					news.Groups = groups
					news.Important = important
					news.Category = category
					news.Creator = creator
					newsDao.Insert(news)
					return  news, nil
				},
			},
			"updateNews": &graphql.Field{
				Type: newsType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"intern": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
					"tags": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"groups": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
					"important": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"creator": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					title, _ := p.Args["title"].(string)
					text, _ := p.Args["text"].(string)
					intern, _ := p.Args["intern"].(bool)
					tags, _ := p.Args["tags"].([]string)
					groups, _ := p.Args["groups"].([]string)
					important, _ := p.Args["important"].(bool)
					category, _ := p.Args["category"].(string)
					creator, _ := p.Args["creator"].(string)
					news := newsModel.News{}
					newsDao.GetById(bson.ObjectIdHex(id)).One(&news)
					news.ModifiedDate = time.Now().Format(time.RFC3339)
					news.Title = title
					news.Text = text
					news.Intern = intern
					news.Tags = tags
					news.Groups = groups
					news.Important = important
					news.Category = category
					news.Creator = creator
					newsDao.Update(news)
					return  news, nil
				},
			},
			"removeNews": &graphql.Field{
				Type: newsType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					news := newsModel.News{}
					newsDao.GetById(bson.ObjectIdHex(id)).One(&news)
					newsDao.Delete(news)
					return  news, nil
				},
			},
			"disableNews": &graphql.Field{
				Type: newsType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					disable, _ := p.Args["disable"].(bool)
					news := newsModel.News{}
					newsDao.GetById(bson.ObjectIdHex(id)).One(&news)
					news.ModifiedDate = time.Now().Format(time.RFC3339)
					news.Disabled = disable
					newsDao.Update(news)
					return  news, nil
				},
			},
			"pushNewsFile": &graphql.Field{
				Type: fileType.Type,
			},
			"removeNewsFile": &graphql.Field{
				Type: fileType.Type,
			},
			"pushNewsActivity": &graphql.Field{
				Type: activityType.Type,
				Args: graphql.FieldConfigArgument{
					"referenceId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"comment": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"intern": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
					"creator": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					referenceId, _ := p.Args["referenceId"].(string)
					comment, _ := p.Args["comment"].(string)
					intern, _ := p.Args["intern"].(bool)
					creator, _ := p.Args["creator"].(string)
					activity := activityModel.Activity{}
					activity.ID = bson.NewObjectId()
					activity.Disabled = false
					activity.CreatedDate = time.Now().Format(time.RFC3339)
					activity.ModifiedDate = time.Now().Format(time.RFC3339)
					activity.ReferenceClass = "News"
					activity.ReferenceId = referenceId
					activity.Comment = comment
					activity.Intern = intern
					activity.Creator = creator
					newsActivityDao.Insert(activity)
					return  activity, nil
				},
			},
			"removeNewsActivity": &graphql.Field{
				Type: activityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					activity := activityModel.Activity{}
					newsActivityDao.GetById(bson.ObjectIdHex(id)).One(&activity)
					newsActivityDao.Delete(activity)
					return  activity, nil
				},
			},
			"createOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					orgUnit := orgUnitModel.OrgUnit{}
					orgUnit.ID = bson.NewObjectId()
					orgUnit.Disabled = false
					orgUnit.CreatedDate = time.Now().Format(time.RFC3339)
					orgUnit.ModifiedDate = time.Now().Format(time.RFC3339)
					orgUnit.Name = name
					orgUnitDao.AddOrgUnit(orgUnit)
					return  orgUnit, nil
				},
			},
			"updateOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					name, _ := p.Args["name"].(string)
					orgUnit := orgUnitDao.GetByKey(bson.ObjectIdHex(idQuery))
					orgUnit.ModifiedDate = time.Now().Format(time.RFC3339)
					orgUnit.Name = name
					orgUnitDao.UpdateOrgUnit(orgUnit)
					return  orgUnit, nil
				},
			},
			"removeOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
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
					orgUnit := orgUnitDao.GetByKey(bson.ObjectIdHex(idQuery))
					orgUnitDao.Delete(orgUnit)
					return orgUnit, nil
				},
			},
			"disableOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					disable, _ := p.Args["disable"].(bool)
					orgUnit := orgUnitDao.GetByKey(bson.ObjectIdHex(idQuery))
					orgUnit.ModifiedDate = time.Now().Format(time.RFC3339)
					orgUnit.Disabled = disable
					orgUnitDao.UpdateOrgUnit(orgUnit)
					return  orgUnit, nil
				},
			},
			"createPriority": &graphql.Field{
				Type: priorityType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					priority := priorityModel.Priority{}
					priority.ID = bson.NewObjectId()
					priority.Disabled = false
					priority.CreatedDate = time.Now().Format(time.RFC3339)
					priority.ModifiedDate = time.Now().Format(time.RFC3339)
					priority.Name = name
					priorityDao.AddPriority(priority)
					return  priority, nil
				},
			},
			"updatePriority": &graphql.Field{
				Type: priorityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					name, _ := p.Args["name"].(string)
					priority := priorityDao.GetByKey(bson.ObjectIdHex(idQuery))
					priority.ModifiedDate = time.Now().Format(time.RFC3339)
					priority.Name = name
					priorityDao.UpdatePriority(priority)
					return  priority, nil
				},
			},
			"removePriority": &graphql.Field{
				Type: priorityType.Type,
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
					priority := priorityDao.GetByKey(bson.ObjectIdHex(idQuery))
					priorityDao.Delete(priority)
					return priority, nil
				},
			},
			"disablePriority": &graphql.Field{
				Type: priorityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					disable, _ := p.Args["disable"].(bool)
					priority := priorityDao.GetByKey(bson.ObjectIdHex(idQuery))
					priority.ModifiedDate = time.Now().Format(time.RFC3339)
					priority.Disabled = disable
					priorityDao.UpdatePriority(priority)
					return  priority, nil
				},
			},
			"createPermission": &graphql.Field{
				Type: permissionType.Type,
			},
			"updatePermission": &graphql.Field{
				Type: permissionType.Type,
			},
			"removePermission": &graphql.Field{
				Type: permissionType.Type,
			},
			"disablePermission": &graphql.Field{
				Type: permissionType.Type,
			},
			"createTag": &graphql.Field{
				Type: tagType.Type,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"style": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					style, _ := p.Args["style"].(string)
					tag := tagModel.Tag{}
					tag.ID = bson.NewObjectId()
					tag.Disabled = false
					tag.CreatedDate = time.Now().Format(time.RFC3339)
					tag.ModifiedDate = time.Now().Format(time.RFC3339)
					tag.Name = name
					tag.Style = style
					tagDao.AddTag(tag)
					return  tag, nil
				},
			},
			"updateTag": &graphql.Field{
				Type: tagType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"style": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					name, _ := p.Args["name"].(string)
					style, _ := p.Args["style"].(string)
					tag := tagDao.GetByKey(bson.ObjectIdHex(idQuery))
					tag.ModifiedDate = time.Now().Format(time.RFC3339)
					tag.Name = name
					tag.Style = style
					tagDao.UpdateTag(tag)
					return  tag, nil
				},
			},
			"removeTag": &graphql.Field{
				Type: tagType.Type,
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
					tag := tagDao.GetByKey(bson.ObjectIdHex(idQuery))
					tagDao.Delete(tag)
					return tag, nil
				},
			},
			"disableTag": &graphql.Field{
				Type: tagType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disable": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					disable, _ := p.Args["disable"].(bool)
					tag := tagDao.GetByKey(bson.ObjectIdHex(idQuery))
					tag.ModifiedDate = time.Now().Format(time.RFC3339)
					tag.Disabled = disable
					tagDao.UpdateTag(tag)
					return  tag, nil
				},
			},
		},
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"categoryList": &graphql.Field{
				Type: graphql.NewList(categoryType.Type),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					offset, _ := p.Args["offset"].(int)
					limit, _ := p.Args["limit"].(int)
					// type == news oder entity
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					categoryList := []categoryModel.Category{}
					categoryDao.GetAll().Skip(offset).Limit(limit).All(&categoryList)
					return categoryList, nil
				},
			},
			"contactList": &graphql.Field{
				Type: graphql.NewList(contactType.Type),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					offset, _ := p.Args["offset"].(int)
					limit, _ := p.Args["limit"].(int)
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					contactList := []contactModel.Contact{}
					contactDao.GetAll(bson.M{}).Skip(offset).Limit(limit).All(&contactList)
					return contactList, nil
				},
			},
			"entityList": &graphql.Field{
				Type: graphql.NewList(entityType.Type),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					offset, _ := p.Args["offset"].(int)
					limit, _ := p.Args["limit"].(int)
					entityList := []entityModel.Entity{}
					entityDao.GetAll(bson.M{}).Skip(offset).Limit(limit).All(&entityList)
					return entityList, nil
					// curl -g 'http://localhost:8080/graphql?query={entityList{id}}'
					// return entityModel.Entity{CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"groupList": &graphql.Field{
				Type: graphql.NewList(groupType.Type),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					offset, _ := p.Args["offset"].(int)
					limit, _ := p.Args["limit"].(int)
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					groupList := []groupModel.Group{}
					groupDao.GetAll().Skip(offset).Limit(limit).All(&groupList)
					return groupList, nil
				},
			},
			"newsList": &graphql.Field{
				Type: graphql.NewList(newsType.Type),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					offset, _ := p.Args["offset"].(int)
					limit, _ := p.Args["limit"].(int)
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					newsList := []newsModel.News{}
					newsDao.GetAll(bson.M{}).Skip(offset).Limit(limit).All(&newsList)
					return newsList, nil
				},
			},
			"orgUnitList": &graphql.Field{
				Type: graphql.NewList(orgUnitType.Type),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					offset, _ := p.Args["offset"].(int)
					limit, _ := p.Args["limit"].(int)
					orgUnitList := []orgUnitModel.OrgUnit{}
					orgUnitDao.GetAll().Skip(offset).Limit(limit).All(&orgUnitList)
					return orgUnitList, nil
				},
			},
			"priorityList": &graphql.Field{
				Type: graphql.NewList(priorityType.Type),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					offset, _ := p.Args["offset"].(int)
					limit, _ := p.Args["limit"].(int)
					priorityList := []priorityModel.Priority{}
					priorityDao.GetAll().Skip(offset).Limit(limit).All(&priorityList)
					return priorityList, nil
				},
			},
			"tagList": &graphql.Field{
				Type: graphql.NewList(tagType.Type),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					offset, _ := p.Args["offset"].(int)
					limit, _ := p.Args["limit"].(int)
					tagList := []tagModel.Tag{}
					tagDao.GetAll().Skip(offset).Limit(limit).All(&tagList)
					return tagList, nil
				},
			},
			"category": &graphql.Field{
				Type: categoryType.Type,
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
					return categoryDao.GetByKey(bson.ObjectIdHex(idQuery)), nil
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
					contact := contactModel.Contact{}
					contactDao.GetById(bson.ObjectIdHex(idQuery)).One(&contact)
					return contact, nil
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
						return nil, nil
					}
					// curl -g 'http://localhost:8080/graphql?query={category(name:"catName"){name}}'
					return groupDao.GetByKey(bson.ObjectIdHex(idQuery)), nil
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
						return nil, nil
					}
					news := newsModel.News{}
					newsDao.GetById(bson.ObjectIdHex(idQuery)).One(&news)
					return news, nil
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
						return nil, nil
					}
					return orgUnitDao.GetByKey(bson.ObjectIdHex(idQuery)), nil
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
						return nil, nil
					}
					return priorityDao.GetByKey(bson.ObjectIdHex(idQuery)), nil
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
						return nil, nil
					}
					return tagDao.GetByKey(bson.ObjectIdHex(idQuery)), nil
				},
			},
			"entityActivites": &graphql.Field{
				Type: graphql.NewList(activityType.Type),
				Args: graphql.FieldConfigArgument{
					"referenceId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					referenceId, _ := p.Args["referenceId"].(string)
					if !bson.IsObjectIdHex(referenceId) {
						return nil, nil
					}
					activityList := []activityModel.Activity{}
					entityActivityDao.GetAll(referenceId).All(&activityList)
					return activityList, nil
				},
			},
			"newsActivites": &graphql.Field{
				Type: graphql.NewList(activityType.Type),
				Args: graphql.FieldConfigArgument{
					"referenceId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					referenceId, _ := p.Args["referenceId"].(string)
					if !bson.IsObjectIdHex(referenceId) {
						return nil, nil
					}
					activityList := []activityModel.Activity{}
					newsActivityDao.GetAll(referenceId).All(&activityList)
					return activityList, nil
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
	http.HandleFunc("/graphql", log(serveGraphQL(Schema)))
	http.HandleFunc("/graphiql", graphiql.ServeGraphiQL)
	http.HandleFunc("/auth", RestAuth())
	http.HandleFunc("/register", RestRegister())
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func RestAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			var account accountModel.Account
			err := decoder.Decode(&account)
			if err != nil {
				panic(err)
			}
			defer r.Body.Close()
			accountDao.GetAll(account).One(&account)
			jwt := jwtModel.JWT{}
			jwt.JWT = authHandler.CreateJWT(account)
			jwt.Account = account
			json.NewEncoder(w).Encode(&jwt)
			cookie := http.Cookie{HttpOnly:true, Name:"jwt",Value:jwt.JWT,MaxAge:0}
			http.SetCookie(w, &cookie)
		}
	}
}

func RestRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			var account accountModel.Account
			err := decoder.Decode(&account)
			if err != nil {
				panic(err)
			}
			defer r.Body.Close()
			account.ID = bson.NewObjectId()
			accountDao.Insert(account)
			account.Roles = []string{"customer"}
			account.Groups = []string{"customer"}
			jwt := jwtModel.JWT{}
			jwt.JWT = authHandler.CreateJWT(account)
			jwt.Account = account
			json.NewEncoder(w).Encode(&jwt)
		}
	}
}

func log(fn http.HandlerFunc) http.HandlerFunc {
	/*
	return func(w http.ResponseWriter, r *http.Request) {
    if len(r.URL.Query().Get("key") == 0) {
      http.Error(w, "missing key", http.StatusUnauthorized)
      return // don't call original handler
    }
    fn(w, r)
  }
	 */
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString != "" {
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("secretKeyM8"), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println("Before", claims["accountId"])
				account := accountModel.Account{}
				account.ID = bson.ObjectIdHex(claims["accountId"].(string))
				account.Groups = claims["groups"].(string)
				account.Roles = claims["roles"].(string)
				authHandler.CurrentAccount = account
				fn(w, r)
			} else {
				fmt.Println("After")
				http.Error(w, "missing key", http.StatusUnauthorized)
			}
		} else {
			http.Error(w, "missing key", http.StatusUnauthorized)
		}
	}
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
