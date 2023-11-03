package scripts

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/BloomGameStudio/EnvironmentService/controllers/ws/errorHandlers"
	privateModels "github.com/BloomGameStudio/EnvironmentService/models/private"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func scriptsReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:

			// Initializer model to bind into
			// NOTE: We are using a private model here TODO: Change to public model in production or handle this case
			reqModelArr := &[]privateModels.Scripts{}

			err := ws.ReadJSON(reqModelArr)

			if err != nil {

				switch err.(type) {

				case *json.UnmarshalTypeError:
					c.Logger().Info(err)

				default:
					errorHandlers.HandleReadError(c, ch, err)
					return
				}
			}

			for _, reqModel := range *reqModelArr {

				// TODO: Clean this up
				if reqModel.ID <= 0 {
					ch <- errors.New("missing/invalid ID")
					return
				}

				if !reqModel.IsValid() {
					// NOTE: no Chan Timeout used
					ch <- errors.New("reqModel Validation failed")
					return
				}

				// Use dot annotation for promoted aka embedded fields.
				model := &privateModels.Scripts{}
				// TODO: Handle ID and production mode

				if viper.GetBool("DEBUG") {
					// Accept client provided ID in DEBUG mode
					model.ID = reqModel.ID

				}

				if !model.IsValid() {
					// NOTE: no Chan Timeout used
					ch <- errors.New("Model Validation failed")
					return
				}

				// handlers.Level(*levelModel, c)
			}
		}
	}
}
