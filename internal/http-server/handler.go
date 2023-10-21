package httpserver

import (
	"emobletest/internal/http-server/response"
	"emobletest/internal/lib/logger"
	"emobletest/internal/storage/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (api *API) endpoints() {
	api.r.Post("/api/v1/users/", api.createUser)
	api.r.Delete("/api/v1/users/{id}", api.deleteUser)
	api.r.Put("/api/v1/users/{id}", api.updateUser)
	api.r.Post("/api/v1/users/get/", api.getUser)
}

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	var usr model.User

	err := render.DecodeJSON(r.Body, &usr)
	if err != nil {
		api.logger.Debug("error while decoding json body", logger.Err(err))
		render.JSON(w, r, response.Error("enter correct data"))
		return
	}
	if usr.Name == "" || usr.Surname == "" {
		api.logger.Debug("incorrect input")
		render.JSON(w, r, response.Error("enter correct data"))
		return
	}

	id, err := api.db.CreateUser(usr)
	if err != nil {
		api.logger.Debug("error while creating user", logger.Err(err))
		render.JSON(w, r, response.Error("cannot create user"))
		return
	}
	api.logger.Info("user created")
	render.JSON(w, r, response.OK(map[string]int{"id": id}))
}

func (api *API) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		api.logger.Debug("error while converting id", logger.Err(err))
		render.JSON(w, r, response.Error("enter correct id"))
		return
	}

	err = api.db.DeleteUser(newId)
	if err != nil {
		api.logger.Debug("error while deleting user", logger.Err(err))
		render.JSON(w, r, response.Error("cannot delete user"))
		return
	}
	api.logger.Info("user deleted")
	render.JSON(w, r, response.OK(map[string]string{"Message": "User deleted"}))
}

func (api *API) updateUser(w http.ResponseWriter, r *http.Request) {
	var ui model.UpdateInput

	id := chi.URLParam(r, "id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		api.logger.Debug("error while converting id", logger.Err(err))
		render.JSON(w, r, response.Error("enter correct id"))
		return
	}
	err = render.DecodeJSON(r.Body, &ui)
	if err != nil {
		api.logger.Debug("error while decoding json body", logger.Err(err))
		render.JSON(w, r, response.Error("enter correct data"))
		return
	}
	err = api.db.UpdateUser(newId, ui)
	if err != nil {
		api.logger.Debug("error while updating user", logger.Err(err))
		render.JSON(w, r, response.Error("cannot update user"))
		return
	}
	api.logger.Info("user updated")
	render.JSON(w, r, response.OK(map[string]string{"Message": "User updated"}))
}

func (api *API) getUser(w http.ResponseWriter, r *http.Request) {
	var gi model.GetInput
	err := render.DecodeJSON(r.Body, &gi)
	if err != nil {
		api.logger.Debug("error while decoding json body", logger.Err(err))
		render.JSON(w, r, response.Error("enter correct data"))
		return
	}
	usrs, err := api.db.GetUser(gi)
	if err != nil {
		api.logger.Debug("error while getting user", logger.Err(err))
		render.JSON(w, r, response.Error("cannot get user"))
		return
	}
	api.logger.Info("requested users info")
	render.JSON(w, r, usrs)

}
