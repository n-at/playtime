package web

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"net/http"
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

func (s *Server) saveStateDeleteForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	return c.Render(http.StatusOK, "state_state_delete", pongo2.Context{
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
