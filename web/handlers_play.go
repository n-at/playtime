package web

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"net/http"
	"playtime/storage"
)

func (s *Server) play(c echo.Context) error {
	context := c.(*PlaytimeContext)

	game := s.prepareGame(*context.game, *context.user)

	if len(game.Platform) == 0 {
		return errors.New("game platform is undefined")
	}

	settings, err := s.storage.SettingsGetByUserId(context.user.Id)
	if err != nil {
		return err
	}

	emulatorSettings := settings.EmulatorSettings[game.Platform]
	if game.OverrideEmulatorSettings {
		emulatorSettings = game.EmulatorSettings
	}

	saveState := storage.SaveState{}
	saveStateId := c.QueryParam("state")
	if len(saveStateId) != 0 {
		saveState, err = s.storage.SaveStateGetById(saveStateId)
		if err != nil {
			return err
		}
		if saveState.GameId != game.Id {
			return errors.New("save state belongs to different game")
		}
		if saveState.UserId != context.user.Id {
			return errors.New("save state belongs to different user")
		}
	}

	bios := storage.Bios{}
	if len(emulatorSettings.Bios) != 0 {
		for _, item := range storage.Bioses[game.Platform] {
			if item.Name == emulatorSettings.Bios {
				bios = item
			}
		}
	}

	return c.Render(http.StatusOK, "play", pongo2.Context{
		"user":              context.user,
		"game":              game,
		"settings":          settings,
		"emulator_settings": emulatorSettings,
		"bios":              bios,
		"save_state":        prepareSaveState(saveState),
	})
}
