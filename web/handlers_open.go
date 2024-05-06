package web

import (
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"net/http"
	"playtime/storage"
	"playtime/web/gamesession"
)

func (s *Server) open(c echo.Context) error {
	context := c.(*PlaytimeContext)

	user, err := s.findContextSessionUser(context)
	if err != nil {
		return err
	}

	if !s.config.NetplayEnabled {
		return c.Render(http.StatusOK, "open", pongo2.Context{
			"user": user,
		})
	}

	activeGames := s.gameSessions.GetActiveGames()

	activeGamesAssoc := make(map[string]gamesession.SessionGame)
	var gameIds []string
	for _, game := range activeGames {
		gameIds = append(gameIds, game.GameId)
		activeGamesAssoc[game.GameId] = game
	}

	games, err := s.storage.GameGetByIdsWithNetplayOpen(gameIds)
	if err != nil {
		return err
	}

	var userIds []string
	for _, game := range games {
		userIds = append(userIds, game.UserId)
	}

	users, err := s.storage.UserFindByIds(userIds)
	if err != nil {
		return err
	}

	usersAssoc := make(map[string]storage.User)
	for _, user := range users {
		usersAssoc[user.Id] = user
	}

	platforms := platformsNames()

	var openGames []openGame
	for _, game := range games {
		openGames = append(openGames, openGame{
			Game:     game,
			Session:  activeGamesAssoc[game.Id],
			User:     usersAssoc[game.UserId],
			Platform: platforms[game.Platform],
		})
	}

	return c.Render(http.StatusOK, "open", pongo2.Context{
		"user":  user,
		"games": openGames,
	})
}
