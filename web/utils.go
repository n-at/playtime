package web

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"playtime/storage"
	"playtime/web/gamesession"
	"sort"
	"strings"
)

type gameByPlatform struct {
	Platform storage.Platform
	Games    []storage.GameWithData
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

func (s *Server) saveStateWithData(saveState storage.SaveState) (storage.SaveStateWithData, error) {
	uploadPath, err := storage.GetUploadPath(saveState.Id)
	if err != nil {
		return storage.SaveStateWithData{}, err
	}

	saveStateWithData := storage.SaveStateWithData{
		SaveState:              saveState,
		StateFileDownloadLink:  fmt.Sprintf("%s/%s/%s.sav", UploadsWebRoot, uploadPath, saveState.Id),
		ScreenshotDownloadLink: fmt.Sprintf("%s/%s/%s.png", UploadsWebRoot, uploadPath, saveState.Id),
	}

	return saveStateWithData, nil
}

func (s *Server) getSaveStatesWithDataByGame(user *storage.User, gameId string) ([]storage.SaveStateWithData, error) {
	core, err := s.getCoreByGameId(user, gameId)
	if err != nil {
		return []storage.SaveStateWithData{}, err
	}

	saveStates, err := s.storage.SaveStateGetByGameIdAndCore(gameId, core)
	if err != nil {
		return []storage.SaveStateWithData{}, err
	}

	var states []storage.SaveStateWithData

	for i := 0; i < len(saveStates); i++ {
		saveStateWithData, err := s.saveStateWithData(saveStates[i])
		if err != nil {
			return []storage.SaveStateWithData{}, err
		}
		states = append(states, saveStateWithData)
	}

	return states, nil
}

func (s *Server) getSaveStateWithDataById(user *storage.User, stateId string) (storage.SaveStateWithData, error) {
	saveState, err := s.storage.SaveStateGetById(stateId)
	if err != nil {
		return storage.SaveStateWithData{}, err
	}

	core, err := s.getCoreByGameId(user, saveState.GameId)
	if err != nil {
		return storage.SaveStateWithData{}, err
	}

	if core != saveState.Core {
		return storage.SaveStateWithData{}, errors.New("save state from different core")
	}

	saveStateWithData, err := s.saveStateWithData(saveState)
	if err != nil {
		return storage.SaveStateWithData{}, err
	}

	return saveStateWithData, nil
}

func (s *Server) getLatestSaveStateWithDataByGameId(user *storage.User, gameId string) (storage.SaveStateWithData, error) {
	core, err := s.getCoreByGameId(user, gameId)
	if err != nil {
		return storage.SaveStateWithData{}, err
	}

	saveState, err := s.storage.SaveStateGetLatestByGameIdAndCore(gameId, core)
	if err != nil {
		return storage.SaveStateWithData{}, err
	}
	if len(saveState.Id) == 0 {
		return storage.SaveStateWithData{}, nil
	}

	saveStateWithData, err := s.saveStateWithData(saveState)
	if err != nil {
		return storage.SaveStateWithData{}, err
	}

	return saveStateWithData, nil
}

///////////////////////////////////////////////////////////////////////////////

func (s *Server) gameWithData(user *storage.User, game storage.Game) (storage.GameWithData, error) {
	uploadPath, err := storage.GetUploadPath(game.Id)
	if err != nil {
		return storage.GameWithData{}, err
	}

	gameWithData := storage.GameWithData{
		Game:         game,
		DownloadLink: fmt.Sprintf("%s/%s/%s", UploadsWebRoot, uploadPath, game.Id),
	}

	core, err := s.getCoreByGameId(user, game.Id)
	if err != nil {
		log.Warnf("getGameWithDataById unable to get game core for %s: %s", game.Id, err)
		core = ""
	}
	if len(core) > 0 {
		latestSaveState, err := s.getLatestSaveStateWithDataByGameId(user, game.Id)
		if err != nil {
			log.Warnf("getGameWithDataById unable to get latest save state for %s: %s", game.Id, err)
			latestSaveState = storage.SaveStateWithData{}
		}
		gameWithData.LatestSaveState = latestSaveState
	}

	return gameWithData, nil
}

func (s *Server) getGameWithDataById(user *storage.User, gameId string) (storage.GameWithData, error) {
	game, err := s.storage.GameGetById(gameId)
	if err != nil {
		return storage.GameWithData{}, err
	}

	gameWithData, err := s.gameWithData(user, game)
	if err != nil {
		return storage.GameWithData{}, err
	}

	return gameWithData, nil
}

func (s *Server) getGamesWithDataByUser(user *storage.User) ([]storage.GameWithData, error) {
	games, err := s.storage.GameGetByUserId(user.Id)
	if err != nil {
		return []storage.GameWithData{}, err
	}

	var gamesWithData []storage.GameWithData

	for i := 0; i < len(games); i++ {
		gameWithData, err := s.gameWithData(user, games[i])
		if err != nil {
			return []storage.GameWithData{}, err
		}
		gamesWithData = append(gamesWithData, gameWithData)
	}

	return gamesWithData, nil
}

func (s *Server) groupGamesByPlatform(games []storage.GameWithData) []gameByPlatform {
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

func (s *Server) findNetplayControls(context *PlaytimeContext) storage.EmulatorControls {
	game := context.game

	if game == nil || len(game.Platform) == 0 {
		return storage.EmulatorControls{}
	}

	if context.session != nil && len(context.session.UserId) != 0 {
		userSettings, err := s.storage.SettingsGetByUserId(context.session.UserId)
		if err != nil {
			log.Warnf("unable to get current user %s settings: %s", context.settings.UserId, err)
		} else {
			return userSettings.EmulatorSettings[game.Platform].Controls[0]
		}
	}

	userSettings, err := s.storage.SettingsGetByUserId(game.UserId)
	if err != nil {
		log.Warnf("unable to get user %s settings: %s", game.UserId, err)
	} else {
		return userSettings.EmulatorSettings[game.Platform].Controls[0]
	}

	return storage.DefaultEmulatorSettings(game.Platform).Controls[0]
}

func (s *Server) collectNetplayCurrentSessionClients(session *gamesession.GameSession) []gamesession.MessageGreetingClient {
	var greetingClients []gamesession.MessageGreetingClient

	for _, clientId := range session.GetClients() {
		client := session.GetClient(clientId)
		if client == nil {
			continue
		}
		greetingClients = append(greetingClients, gamesession.MessageGreetingClient{
			Id:     client.GetId(),
			Name:   client.GetName(),
			Player: client.GetPlayer(),
		})
	}

	return greetingClients
}

func (s *Server) sendWebSocketError(ws *websocket.Conn, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), gamesession.SendTimeout)
	defer cancel()

	if err := wsjson.Write(ctx, ws, gamesession.MessageError(message)); err != nil {
		log.Warnf("unable to send error message: %s", err)
	}
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

func getFileExtension(name string) string {
	if len(name) == 0 {
		return ""
	}

	parts := strings.Split(name, ".")
	if len(parts) == 1 {
		return ""
	}

	return parts[len(parts)-1]
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
