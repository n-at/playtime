package web

import "playtime/storage"

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
