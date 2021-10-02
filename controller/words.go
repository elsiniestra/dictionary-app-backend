package controller

import (
	"net/http"
	"strings"

	"github.com/fallncrlss/dictionary-app-backend/lib/customerrors"
	"github.com/fallncrlss/dictionary-app-backend/lib/enums"
	"github.com/fallncrlss/dictionary-app-backend/logger"
	"github.com/fallncrlss/dictionary-app-backend/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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
	language, isCorrect := enums.GetLanguageCodes()[languageCode]

	if !isCorrect {
		return errors.Wrapf(customerrors.ErrIncorrectArgument, "arguments (language='%s')", language)
	}

	wc.logger.Debug().Msgf("Getting word '%s' (language='%s') from database...", name, language)

	wordData, err := wc.services.Word.GetWordWithDB(name, language)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrUnableFetchInstance), errors.Is(err, customerrors.ErrFetchedInstanceIsNil):
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

		case errors.Is(err, customerrors.ErrIncorrectArgument):
			return echo.NewHTTPError(http.StatusBadRequest, err)

		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "Could not get word"))
		}
	}

	wc.logger.Debug().Msgf("Return successful response: %s", wordData)

	err = ctx.JSON(http.StatusOK, wordData)
	if err != nil {
		return errors.Wrap(err, "Ssending a JSON response with status code failed")
	}
	return nil
}
