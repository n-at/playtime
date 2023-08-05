package web

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"playtime/storage"
	"sort"
	"strings"
)

type gameByPlatform struct {
	Platform storage.Platform
	Games    []storage.Game
}

///////////////////////////////////////////////////////////////////////////////

func (s *Server) getCoreByGameId(user *storage.User, gameId string) (string, error) {
	settings, err := s.storage.SettingsGetByUserId(user.Id)
	if err != nil {
		return "", err
	}

	game, err := s.storage.GameGetById(gameId)
	if err != nil {
		return "", err
	}

	if game.OverrideEmulatorSettings {
		return game.EmulatorSettings.Core, nil
	} else {
		return settings.EmulatorSettings[game.Platform].Core, nil
	}
}

///////////////////////////////////////////////////////////////////////////////

func (s *Server) saveStateWithData(saveState *storage.SaveState) error {
	uploadPath, err := getUploadPath(saveState.Id)
	if err != nil {
		return err
	}

	saveState.StateFileDownloadLink = fmt.Sprintf("%s/%s/%s.sav", UploadsWebRoot, uploadPath, saveState.Id)
	saveState.ScreenshotDownloadLink = fmt.Sprintf("%s/%s/%s.png", UploadsWebRoot, uploadPath, saveState.Id)

	return nil
}

func (s *Server) getSaveStatesWithDataByGame(user *storage.User, gameId string) ([]storage.SaveState, error) {
	core, err := s.getCoreByGameId(user, gameId)
	if err != nil {
		return []storage.SaveState{}, err
	}

	saveStates, err := s.storage.SaveStateGetByGameIdAndCore(gameId, core)
	if err != nil {
		return []storage.SaveState{}, err
	}

	for i := 0; i < len(saveStates); i++ {
		if err := s.saveStateWithData(&saveStates[i]); err != nil {
			return []storage.SaveState{}, err
		}
	}

	return saveStates, nil
}

func (s *Server) getSaveStateWithDataById(user *storage.User, stateId string) (storage.SaveState, error) {
	saveState, err := s.storage.SaveStateGetById(stateId)
	if err != nil {
		return storage.SaveState{}, err
	}

	core, err := s.getCoreByGameId(user, saveState.GameId)
	if err != nil {
		return storage.SaveState{}, err
	}

	if core != saveState.Core {
		return storage.SaveState{}, errors.New("save state from different core")
	}
	if err := s.saveStateWithData(&saveState); err != nil {
		return storage.SaveState{}, err
	}

	return saveState, nil
}

func (s *Server) getLatestSaveStateWithDataByGameId(user *storage.User, gameId string) (storage.SaveState, error) {
	core, err := s.getCoreByGameId(user, gameId)
	if err != nil {
		return storage.SaveState{}, err
	}

	saveState, err := s.storage.SaveStateGetLatestByGameIdAndCore(gameId, core)
	if err != nil {
		return storage.SaveState{}, err
	}
	if len(saveState.Id) == 0 {
		return storage.SaveState{}, nil
	}
	if err := s.saveStateWithData(&saveState); err != nil {
		return storage.SaveState{}, err
	}

	return saveState, nil
}

///////////////////////////////////////////////////////////////////////////////

func (s *Server) gameWithData(user *storage.User, game *storage.Game) error {
	uploadPath, err := getUploadPath(game.Id)
	if err != nil {
		return err
	}

	game.DownloadLink = fmt.Sprintf("%s/%s/%s", UploadsWebRoot, uploadPath, game.Id)

	core, err := s.getCoreByGameId(user, game.Id)
	if err != nil {
		log.Warnf("getGameWithDataById unable to get game core for %s: %s", game.Id, err)
		core = ""
	}
	if len(core) > 0 {
		latestSaveState, err := s.getLatestSaveStateWithDataByGameId(user, game.Id)
		if err != nil {
			log.Warnf("getGameWithDataById unable to get latest save state for %s: %s", game.Id, err)
			latestSaveState = storage.SaveState{}
		}
		game.LatestSaveState = latestSaveState
	}

	return nil
}

func (s *Server) getGameWithDataById(user *storage.User, gameId string) (storage.Game, error) {
	game, err := s.storage.GameGetById(gameId)
	if err != nil {
		return storage.Game{}, err
	}

	if err := s.gameWithData(user, &game); err != nil {
		return storage.Game{}, err
	}

	return game, nil
}

func (s *Server) getGamesWithDataByUser(user *storage.User) ([]storage.Game, error) {
	games, err := s.storage.GameGetByUserId(user.Id)
	if err != nil {
		return []storage.Game{}, err
	}

	for i := 0; i < len(games); i++ {
		if err := s.gameWithData(user, &games[i]); err != nil {
			return []storage.Game{}, err
		}
	}

	return games, nil
}

func (s *Server) groupGamesByPlatform(games []storage.Game) []gameByPlatform {
	gamesByPlatform := make(map[string]*gameByPlatform)

	for _, game := range games {
		_, ok := gamesByPlatform[game.Platform]
		if !ok {
			gamesByPlatform[game.Platform] = &gameByPlatform{
				Platform: storage.Platforms[game.Platform],
			}
		}
		gamesByPlatform[game.Platform].Games = append(gamesByPlatform[game.Platform].Games, game)
	}

	var platforms []gameByPlatform
	for _, platform := range gamesByPlatform {
		platforms = append(platforms, *platform)
	}

	sort.Slice(platforms, func(i, j int) bool {
		return platforms[i].Platform.Name < platforms[j].Platform.Name
	})

	return platforms
}

///////////////////////////////////////////////////////////////////////////////

func sortedPlatforms() []storage.Platform {
	var platforms []storage.Platform

	for _, platform := range storage.Platforms {
		platforms = append(platforms, platform)
	}

	sort.Slice(platforms, func(i, j int) bool {
		return platforms[i].Name < platforms[j].Name
	})

	return platforms
}

func guessGameProperties(games []storage.Game) []storage.Game {
	var output []storage.Game

	for _, game := range games {
		game.Name = cleanupName(game.Name)
		game.Platform = guessGamePlatform(game.OriginalFileExtension)
		output = append(output, game)
	}

	return output
}

func cleanupName(name string) string {
	if len(name) == 0 {
		return name
	}

	parts := strings.Split(name, "_")
	name = strings.Join(parts, " ")
	parts = strings.Split(name, ".")
	if len(parts) > 1 {
		parts = parts[0 : len(parts)-1]
	}
	return strings.Join(parts, "")
}

func guessGamePlatform(ext string) string {
	for _, platform := range storage.Platforms {
		for _, extension := range platform.Extensions {
			if extension == ext {
				return platform.Id
			}
		}
	}
	return ""
}

func startsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}

	return true
}
