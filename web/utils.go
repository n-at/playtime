package web

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"playtime/storage"
	"sort"
	"strings"
)

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

///////////////////////////////////////////////////////////////////////////////

type gameByPlatform struct {
	Platform storage.Platform
	Games    []storage.Game
}

func (s *Server) prepareGamesByPlatform(games []storage.Game) []gameByPlatform {
	gamesByPlatform := make(map[string]*gameByPlatform)

	for _, game := range games {
		game = s.prepareGame(game)
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

func (s *Server) prepareGame(game storage.Game) storage.Game {
	uploadPath, err := getUploadPath(game.Id)
	if err != nil {
		uploadPath = ""
	}
	game.DownloadLink = fmt.Sprintf("%s/%s/%s", UploadsWebRoot, uploadPath, game.Id)

	saveState, err := s.storage.SaveStateGetLatestByGameId(game.Id)
	if err != nil {
		log.Warnf("prepareGame unable to get latest save state for %s: %s", game.Id, err)
		saveState = storage.SaveState{}
	}
	game.LatestSaveState = prepareSaveState(saveState)

	return game
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

func prepareSaveStates(states []storage.SaveState) []storage.SaveState {
	for i := 0; i < len(states); i++ {
		states[i] = prepareSaveState(states[i])
	}
	return states
}

func prepareSaveState(state storage.SaveState) storage.SaveState {
	if len(state.Id) == 0 {
		return state
	}

	uploadUrl, err := getUploadPath(state.Id)
	if err != nil {
		uploadUrl = ""
	}

	state.StateFileDownloadLink = fmt.Sprintf("%s/%s/%s.sav", UploadsWebRoot, uploadUrl, state.Id)
	state.ScreenshotDownloadLink = fmt.Sprintf("%s/%s/%s.png", UploadsWebRoot, uploadUrl, state.Id)

	return state
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
