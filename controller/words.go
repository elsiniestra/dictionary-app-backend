package controller

import (
	"github.com/fallncrlss/dictionary-app-backend/lib/customErrors"
	"github.com/fallncrlss/dictionary-app-backend/lib/enums"
	"github.com/fallncrlss/dictionary-app-backend/logger"
	"github.com/fallncrlss/dictionary-app-backend/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type WordController struct {
	services *service.Manager
	logger   *logger.Logger
}

func NewWordControllers(services *service.Manager, logger *logger.Logger) *WordController {
	return &WordController{
		services: services,
		logger:   logger,
	}
}

func (wc *WordController) Get(ctx echo.Context) error {
	languageCode := ctx.Param("language")
	name := strings.ToLower(ctx.Param("name"))
	language, isCorrect := enums.LanguageCodes[languageCode]
	if !isCorrect {
		return errors.Wrapf(customErrors.IncorrectArgument, "arguments (language='%s')", language)
	}

	wc.logger.Debug().Msgf("Getting word '%s' (language='%s') from database...", name, language)
	wordData, err := wc.services.Word.GetWordWithDB(name, language)
	if err != nil {
		switch errors.Cause(err) {
		case customErrors.UnableFetchInstance, customErrors.FetchedInstanceIsNil:
			wc.logger.Debug().Msgf("Unable to get word from DB, getting word '%s' from web... Error: %s", name, err)
			wordData, err = wc.services.Word.GetWordWithWeb(name, languageCode)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err)
			}
			wc.logger.Debug().Msgf("Saving word '%s' into DB...", name)
			err := wc.services.Word.SaveWordToDB(wordData)
			if err != nil {
				wc.logger.Warn().Err(err)
			}

		case customErrors.IncorrectArgument:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "Could not get word"))
		}
	}
	wc.logger.Debug().Msgf("Return successful response: %s", wordData)
	return ctx.JSON(http.StatusOK, wordData)
}
