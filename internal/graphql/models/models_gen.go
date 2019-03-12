// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"

	"github.com/dictyBase/go-genproto/dictybaseapis/user"
)

type Author struct {
	LastName  *string `json:"last_name"`
	FirstName *string `json:"first_name"`
	Initials  *string `json:"initials"`
	Rank      *string `json:"rank"`
}

type CreatePermissionInput struct {
	Permission  string `json:"permission"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
}

type CreateRoleInput struct {
	Role        string `json:"role"`
	Description string `json:"description"`
}

type CreateUserInput struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Email         string  `json:"email"`
	Organization  *string `json:"organization"`
	GroupName     *string `json:"group_name"`
	FirstAddress  *string `json:"first_address"`
	SecondAddress *string `json:"second_address"`
	City          *string `json:"city"`
	State         *string `json:"state"`
	Zipcode       *string `json:"zipcode"`
	Country       *string `json:"country"`
	Phone         *string `json:"phone"`
	IsActive      bool    `json:"is_active"`
}

type DeleteItem struct {
	Success bool `json:"success"`
}

type Publication struct {
	ID       string     `json:"id"`
	Doi      *string    `json:"doi"`
	Title    *string    `json:"title"`
	Abstract *string    `json:"abstract"`
	Journal  *string    `json:"journal"`
	PubDate  *time.Time `json:"pub_date"`
	Volume   *string    `json:"volume"`
	Pages    *string    `json:"pages"`
	Issn     *string    `json:"issn"`
	PubType  *string    `json:"pub_type"`
	Source   *string    `json:"source"`
	Issue    *string    `json:"issue"`
	Status   *string    `json:"status"`
	Authors  []*Author  `json:"authors"`
}

type UpdatePermissionInput struct {
	Permission  string `json:"permission"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
}

type UpdateRoleInput struct {
	Role        string `json:"role"`
	Description string `json:"description"`
}

type UpdateUserInput struct {
	FirstName     *string `json:"first_name"`
	LastName      *string `json:"last_name"`
	Organization  *string `json:"organization"`
	GroupName     *string `json:"group_name"`
	FirstAddress  *string `json:"first_address"`
	SecondAddress *string `json:"second_address"`
	City          *string `json:"city"`
	State         *string `json:"state"`
	Zipcode       *string `json:"zipcode"`
	Country       *string `json:"country"`
	Phone         *string `json:"phone"`
	IsActive      *bool   `json:"is_active"`
}

type UserList struct {
	Users      []user.User `json:"users"`
	PageNum    *string     `json:"pageNum"`
	PageSize   *string     `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
}
