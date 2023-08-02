package web

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"net/http"
	"playtime/storage"
	"time"
)

func (s *Server) saveStates(c echo.Context) error {
	context := c.(*PlaytimeContext)

	game := context.game
	states, err := s.storage.SaveStateGetByGameId(game.Id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "save_states", pongo2.Context{
		"user":   context.user,
		"game":   game,
		"states": prepareSaveStates(states),
	})
}

func (s *Server) saveStateUpload(c echo.Context) error {
	context := c.(*PlaytimeContext)

	saveState := storage.SaveState{
		Id:      storage.NewId(),
		UserId:  context.user.Id,
		GameId:  context.game.Id,
		Created: time.Now(),
	}

	state, err := c.FormFile("state")
	if err != nil {
		return err
	}
	if err := s.saveUploadedFile(state, saveState.Id, "sav"); err != nil {
		return err
	}

	screenshot, err := c.FormFile("screenshot")
	if err != nil {
		return err
	}
	if err := s.saveUploadedFile(screenshot, saveState.Id, "png"); err != nil {
		return err
	}

	if _, err := s.storage.SaveStateSave(saveState); err != nil {
		return err
	}

	saveState = prepareSaveState(saveState)

	return c.JSON(http.StatusOK, saveState)
}

func (s *Server) saveStateDeleteForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	return c.Render(http.StatusOK, "save_state_delete", pongo2.Context{
		"user":  context.user,
		"game":  context.game,
		"state": context.saveState,
	})
}

func (s *Server) saveStateDeleteSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	if err := s.storage.SaveStateDeleteById(context.saveState.Id); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/games/save-states/%s", context.game.Id))
}

func (s *Server) saveStateList(c echo.Context) error {
	context := c.(*PlaytimeContext)

	states, err := s.storage.SaveStateGetByGameId(context.game.Id)
	if err != nil {
		return err
	}

	states = prepareSaveStates(states)

	return c.JSON(http.StatusOK, states)
}
