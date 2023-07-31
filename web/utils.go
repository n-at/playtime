package web

import (
	"playtime/storage"
	"sort"
)

type platformValue struct {
	Type     string
	Platform storage.Platform
}

func sortedPlatforms() []platformValue {
	var platformValues []platformValue

	for _, system := range storage.Systems {
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

func (s *Server) prepareGamesByPlatform(games []storage.Game) []gameByPlatform {
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
