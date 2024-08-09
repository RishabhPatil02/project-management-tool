package graph

import "github.com/rishabhpatil02/project-management-tool/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	users []*model.User
}
