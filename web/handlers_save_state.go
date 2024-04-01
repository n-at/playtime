package web

import (
	"errors"
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
		"_csrf_token": c.Get("csrf"),
		"user":        context.user,
		"game":        context.game,
		"states":      states,
	})
}

func (s *Server) saveStateUpload(c echo.Context) error {
	context := c.(*PlaytimeContext)
	user := context.user
	game := context.game

	core, err := s.getCoreByGameId(user, game.Id)
	if err != nil {
		return err
	}

	saveState := storage.SaveState{
		Id:      storage.NewId(),
		UserId:  user.Id,
		GameId:  game.Id,
		Core:    core,
		Created: time.Now(),
		IsAuto:  c.FormValue("auto") == "1",
	}

	if saveState.IsAuto && !game.AutoSaveEnabled {
		return errors.New("auto saves disabled")
	}

	state, err := c.FormFile("state")
	if err != nil {
		return err
	}
	if err := s.storage.SaveUploadedFile(state, saveState.Id, storage.FileExtensionSaveState); err != nil {
		return err
	}

	screenshot, err := c.FormFile("screenshot")
	if err != nil {
		return err
	}
	if err := s.storage.SaveUploadedFile(screenshot, saveState.Id, storage.FileExtensionScreenshot); err != nil {
		return err
	}

	saveState.Size = state.Size + screenshot.Size
	if user.Quota > 0 && user.GetQuotaUsed()+saveState.Size > user.Quota {
		return errors.New("disk quota exceeded")
	}

	if err := s.storage.SaveStateUpload(saveState); err != nil {
		return err
	}

	if saveState.IsAuto {
		if err := s.storage.SaveStateDeleteAutoByGameId(game.Id, game.AutoSaveCapacity); err != nil {
			return err
		}
	}

	saveStateWithData, err := s.getSaveStateWithDataById(user, saveState.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, saveStateWithData)
}

func (s *Server) saveStateDeleteForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	return c.Render(http.StatusOK, "save_state_delete", pongo2.Context{
		"_csrf_token": c.Get("csrf"),
		"user":        context.user,
		"game":        context.game,
		"state":       context.saveState,
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
