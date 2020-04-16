package controllers

import (
	"calhoun/views"
	"fmt"
	"net/http"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// New - Renders views with the signup form
// Get /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

// Create - Processes data from signup form and creates user
// Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintln(w, form)
}
