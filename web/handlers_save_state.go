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

	states, err := s.getSaveStatesWithDataByGame(context.user, context.game.Id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "save_states", pongo2.Context{
		"user":   context.user,
		"game":   context.game,
		"states": states,
	})
}

func (s *Server) saveStateUpload(c echo.Context) error {
	context := c.(*PlaytimeContext)

	core, err := s.getCoreByGameId(context.user, context.game.Id)
	if err != nil {
		return err
	}

	saveState := storage.SaveState{
		Id:      storage.NewId(),
		UserId:  context.user.Id,
		GameId:  context.game.Id,
		Core:    core,
		Created: time.Now(),
	}

	state, err := c.FormFile("state")
	if err != nil {
		return err
	}
	if err := s.storage.SaveUploadedFile(state, saveState.Id, "sav"); err != nil {
		return err
	}

	screenshot, err := c.FormFile("screenshot")
	if err != nil {
		return err
	}
	if err := s.storage.SaveUploadedFile(screenshot, saveState.Id, "png"); err != nil {
		return err
	}

	if _, err := s.storage.SaveStateSave(saveState); err != nil {
		return err
	}

	saveStateWithData, err := s.getSaveStateWithDataById(context.user, saveState.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, saveStateWithData)
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

	states, err := s.getSaveStatesWithDataByGame(context.user, context.game.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, states)
}
