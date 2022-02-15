package usecases

import (
	"net/http"

	"github.com/AlecSmith96/dnd-scheduler/entities"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type GroupHandler struct {
	DB *gorm.DB
}

// Return all groups with their sessions information
func (handler *GroupHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	var groups entities.GroupList
	if result := handler.DB.Find(&groups.Groups); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}

	if err := render.Render(w, r, &groups); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group entities.Group
	if err := render.Bind(r, &group); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	if result := handler.DB.Create(&group); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}

	// Try to return created player
	if err := render.Render(w, r, &group); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")
	var group entities.Group

	if result := handler.DB.First(&group, "id = ?", groupId); result.Error != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	// Try to return found player
	if err := render.Render(w, r, &group); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler *GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")
	var group entities.Group
	var updatedGroupData entities.Group

	if err := render.Bind(r, &updatedGroupData); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	if result := handler.DB.First(&group, "id = ?", groupId); result.Error != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	if result := handler.DB.Model(&group).Updates(&updatedGroupData); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}
}

// TODO: fix foreign key constraint
func (handler *GroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")

	if result := handler.DB.Delete(&entities.Group{}, "id = ?", groupId); result.Error != nil {
		render.Render(w, r, ErrRender(result.Error))
		return
	}
}

func NewGroupHandler(dbConn *gorm.DB) *GroupHandler {
	return &GroupHandler{
		DB: dbConn,
	}
}
