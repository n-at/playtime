package web

import (
	"errors"
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"playtime/storage"
)

func (s *Server) games(c echo.Context) error {
	context := c.(*PlaytimeContext)

	games, err := s.storage.GameGetByUserId(context.user.Id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "games", pongo2.Context{
		"user":              context.user,
		"games_by_platform": s.prepareGamesByPlatform(games),
	})
}

func (s *Server) gameUpload(c echo.Context) error {
	context := c.(*PlaytimeContext)

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files, ok := form.File["games"]
	if !ok || len(files) == 0 {
		return errors.New("no games uploaded")
	}

	var gameIds []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		game := storage.Game{
			Id:                       storage.NewId(),
			UserId:                   context.user.Id,
			Name:                     file.Filename,
			OriginalFileName:         file.Filename,
			OriginalFileExtension:    getFileExtension(file.Filename),
			Platform:                 "",
			OverrideEmulatorSettings: false,
			EmulatorSettings:         storage.DefaultEmulatorSettings(""),
		}

		uploadPath, err := s.prepareUploadPath(game.Id)
		if err != nil {
			return err
		}

		uploadFilename := fmt.Sprintf("%s%c%s", uploadPath, os.PathSeparator, game.Id)
		dst, err := os.Create(uploadFilename)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		if _, err := s.storage.GameSave(game); err != nil {
			return err
		}

		gameIds = append(gameIds, game.Id)
	}

	uploadBatch := storage.UploadBatch{
		Id:      storage.NewId(),
		UserId:  context.user.Id,
		GameIds: gameIds,
	}

	if _, err := s.storage.UploadBatchSave(uploadBatch); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/upload-batch/%s", uploadBatch.Id))
}

func (s *Server) gameUploadBatchForm(c echo.Context) error {
	return nil
}

func (s *Server) gameUploadBatchSubmit(c echo.Context) error {
	return nil
}

func (s *Server) gameEditForm(c echo.Context) error {
	return nil
}

func (s *Server) gameEditSubmit(c echo.Context) error {
	return nil
}

func (s *Server) gameDeleteForm(c echo.Context) error {
	return nil
}

func (s *Server) gameDeleteSubmit(c echo.Context) error {
	return nil
}
