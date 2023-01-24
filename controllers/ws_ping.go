package controllers

import (
	"errors"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func PingWS(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	writerChan := make(chan error)
	readerChan := make(chan error)

	go writer(c, ws, writerChan)
	go reader(c, ws, readerChan)

	// QUESTION: Do we want to wait on both routines to error out?
	// Return the error if either the reader or the writer encounters a error
	for {
		select {
		case r := <-readerChan:
			c.Logger().Debugf("Recieved readerChan error: %v", r)
			return r
		case w := <-writerChan:
			c.Logger().Debugf("Recieved writerChan error: %v", w)
			return w
		}
	}
}

// Write
func writer(c echo.Context, ws *websocket.Conn, ch chan error) {

forloop:
	for {
		c.Logger().Debug("Writing to the WebSocket")

		c.Logger().Debug("Pushing to the WebSocket")
		err := ws.WriteJSON(&PingResp{"pong"})
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
func reader(c echo.Context, ws *websocket.Conn, ch chan error) {

forloop:
	for {
		c.Logger().Debug("Reading from the WebSocket")

		// Initializer request player to bind into
		reqPing := &PingResp{}
		err := ws.ReadJSON(reqPing)

		if err != nil {
			c.Logger().Debug("We get an error from Reading the JSON reqPing")
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

		c.Logger().Debugf("reqPing from the WebSocket: %+v", reqPing)

		c.Logger().Debug("Validating reqPing")

		c.Logger().Debug("reqPing is valid")

		c.Logger().Debugf("PingRequest: %+v", reqPing)

		c.Logger().Debug("PingRequest is valid passing it to the Ping handler")

		// handlers.Player(*playerModel, c) //TODO: UNCOMNNET and handle errors
	}

}
