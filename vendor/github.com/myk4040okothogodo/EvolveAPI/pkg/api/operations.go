package api

import (
	"net/http"

	"github.com/go-chi/render"

	um "github.com/myk4040okothogodo/EvolveAPI/pkg/middleware"
	"github.com/myk4040okothogodo/EvolveAPI/pkg/types"
)

// GetUser renders the user from the context
// @Summary Get user by id
// @Description GetUser returns a single user by id
// @Tags Users
// @Produce json
// @Param id path string true "user id"
// @Router /users/{id} [get]
// @Success 200 {object} types.User
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(um.UserCtxKey).(*types.User)

	if err := render.Render(w, r, user); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

// PutUser writes a user to the database
// @Summary Add a User to the database
// @Description PutUser writes a user to the database
// @Description To write a new User, leave the id empty. To update an existing one, use the id of the article to be updated
// @Tags Users
// @Produce json
// @Router /users [put]
// @Success 200 {object} types.User
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func PutUser(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}
	if err := render.Bind(r, user); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequest(err))
		return
	}

	if err := DBClient.SetUser(user); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequest(err))
		return
	}

	if err := render.Render(w, r, user); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

// ListUsers returns all Users in the database
// @Summary List all articles
// @Description Get all articles stored in the database
// @Tags Users
// @Produce json
// @Param page_id query string false "id of the page to be retrieved"
// @Router /users [get]
// @Success 200 {object} types.UserList
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func ListUsers(w http.ResponseWriter, r *http.Request) {
	pageID := r.Context().Value(um.PageIDKey)
	if err := render.Render(w, r, DBClient.GetUsers(pageID.(int))); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}


