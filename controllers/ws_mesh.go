package controllers

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func MeshWS(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	meshWriterChan := make(chan error)
	meshReaderChan := make(chan error)

	go meshWriter(c, ws, meshWriterChan)
	go meshReader(c, ws, meshReaderChan)

	// QUESTION: Do we want to wait on both routines to error out?
	// Return the error if either the meshReader or the meshWriter encounters a error
	for {
		select {
		case r := <-meshReaderChan:
			c.Logger().Debugf("Recieved meshReaderChan error: %v", r)
			return r
		case w := <-meshWriterChan:
			c.Logger().Debugf("Recieved meshWriterChan error: %v", w)
			return w
		}
	}
}

// Write
func meshWriter(c echo.Context, ws *websocket.Conn, ch chan error) {

forloop:
	for {
		c.Logger().Debug("Writing to the WebSocket")

		c.Logger().Debug("Pushing to the WebSocket")

		jsonSampleData := `{
    "meshes": [
        {
            "vertices": [
                {
                    "x": 0.0,
                    "y": 0.0,
                    "z": 0.0
                },
                {
                    "x": 0.0,
                    "y": 0.0,
                    "z": 64.0
                },
                {
                    "x": 64.0,
                    "y": 0.0,
                    "z": 0.0
                },
                {
                    "x": 64.0,
                    "y": 0.0,
                    "z": 64.0
                }
            ],
            "triangles": [
                0,
                1,
                2,
                2,
                1,
                3
            ],
            "name": "GameObject",
            "layer": "WaterTerrain",
            "transform": {
                "position": {
                    "x": 20.0,
                    "y": 10.0,
                    "z": -130.0
                },
                "scale": {
                    "x": 1.0,
                    "y": 1.0,
                    "z": 1.0
                },
                "rotation": {
                    "x": 0.0,
                    "y": 0.0,
                    "z": 0.0
                }
            }
        }
    ]
}`

		mesh := &MeshesVertices{}
		_ = json.Unmarshal([]byte(jsonSampleData), &mesh)

		err := ws.WriteJSON(mesh)
		if err != nil {

			switch {

			case errors.Is(err, websocket.ErrCloseSent):
				c.Logger().Debug("WEbsocket ErrCloseSent")
				ch <- nil
				close(ch)
				break forloop

			default:
				c.Logger().Error(err)
				ch <- err
				close(ch)
				break forloop
			}
		}
		c.Logger().Debug("Finished writing to the WebSocket Sleeping now")

		// Update Interval NOTE: setting depending on the server and its performance either increase or decrease it.
		time.Sleep(time.Millisecond * 1)

		if viper.GetBool("DEBUG") {
			// Sleep for 1 second in DEBUG mode to not get fludded with data
			time.Sleep(time.Second * 1)
		}
	}
}

// Read
func meshReader(c echo.Context, ws *websocket.Conn, ch chan error) {

forloop:
	for {
		c.Logger().Debug("Reading from the WebSocket")

		// Initializer request player to bind into
		mesh := &MeshesVertices{}
		err := ws.ReadJSON(mesh)

		if err != nil {
			c.Logger().Debug("We get an error from Reading the JSON mesh")
			switch {

			case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
				c.Logger().Debug("Websocket CloseNoStatusReceived")
				ch <- nil
				close(ch)
				break forloop

			default:
				c.Logger().Error(err)
				ch <- err
				close(ch)
				break forloop

			}
		}

		c.Logger().Debugf("mesh from the WebSocket: %+v", mesh)

		c.Logger().Debug("Validating mesh")

		c.Logger().Debug("mesh is valid")

		c.Logger().Debugf("meshRequest: %+v", mesh)

		c.Logger().Debug("meshRequest is valid passing it to the Mesh handler")

		// handlers.Player(*playerModel, c) //TODO: UNCOMNNET and handle errors
	}

}
