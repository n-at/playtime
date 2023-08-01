package web

import (
	"fmt"
	"playtime/storage"
	"sort"
	"strings"
)

type platformValue struct {
	Type     string
	Platform storage.Platform
}

func sortedPlatforms() []platformValue {
	var platformValues []platformValue

	for _, system := range storage.PlatformIds {
		platformValues = append(platformValues, platformValue{
			Type:     system,
			Platform: storage.Platforms[system],
		})
	}

	return platformValues
}

///////////////////////////////////////////////////////////////////////////////

type gameByPlatform struct {
	Platform storage.Platform
	Games    []storage.Game
}

func prepareGamesByPlatform(games []storage.Game) []gameByPlatform {
	gamesByPlatform := make(map[string]*gameByPlatform)

	for _, game := range games {
		game = prepareGame(game)
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

func prepareGame(game storage.Game) storage.Game {
	uploadPath, err := getUploadPath(game.Id)
	if err != nil {
		uploadPath = ""
	}

	game.DownloadLink = fmt.Sprintf("%s/%s/%s", UploadsWebRoot, uploadPath, game.Id)

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
