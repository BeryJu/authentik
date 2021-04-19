package ak

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gorilla/websocket"
	"github.com/recws-org/recws"
	"goauthentik.io/outpost/pkg"
)

func (ac *APIController) initWS(pbURL url.URL, outpostUUID strfmt.UUID) {
	pathTemplate := "%s://%s/ws/outpost/%s/"
	scheme := strings.ReplaceAll(pbURL.Scheme, "http", "ws")

	authHeader := fmt.Sprintf("Bearer %s", ac.token)

	header := http.Header{
		"Authorization": []string{authHeader},
		"User-Agent":    []string{fmt.Sprintf("authentik-proxy@%s", pkg.VERSION)},
	}

	value, set := os.LookupEnv("AUTHENTIK_INSECURE")
	if !set {
		value = "false"
	}

	ws := &recws.RecConn{
		NonVerbose: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: strings.ToLower(value) == "true",
		},
	}
	ws.Dial(fmt.Sprintf(pathTemplate, scheme, pbURL.Host, outpostUUID.String()), header)

	ac.logger.WithField("logger", "authentik.outpost.ak-ws").WithField("outpost", outpostUUID.String()).Debug("connecting to authentik")

	ac.wsConn = ws
	// Send hello message with our version
	msg := websocketMessage{
		Instruction: WebsocketInstructionHello,
		Args: map[string]interface{}{
			"version": pkg.VERSION,
		},
	}
	err := ws.WriteJSON(msg)
	if err != nil {
		ac.logger.WithField("logger", "authentik.outpost.ak-ws").WithError(err).Warning("Failed to hello to authentik")
	}
}

// Shutdown Gracefully stops all workers, disconnects from websocket
func (ac *APIController) Shutdown() {
	// Cleanly close the connection by sending a close message and then
	// waiting (with timeout) for the server to close the connection.
	err := ac.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		ac.logger.Println("write close:", err)
		return
	}
}

func (ac *APIController) startWSHandler() {
	logger := ac.logger.WithField("loop", "ws-handler")
	for {
		if !ac.wsConn.IsConnected() {
			continue
		}
		var wsMsg websocketMessage
		err := ac.wsConn.ReadJSON(&wsMsg)
		if err != nil {
			logger.Println("read:", err)
			ac.wsConn.CloseAndReconnect()
			continue
		}
		if wsMsg.Instruction == WebsocketInstructionTriggerUpdate {
			time.Sleep(ac.reloadOffset)
			logger.Debug("Got update trigger...")
			err := ac.Server.Refresh()
			if err != nil {
				logger.WithError(err).Debug("Failed to update")
			}
		}
	}
}

func (ac *APIController) startWSHealth() {
	ticker := time.NewTicker(time.Second * 10)
	for ; true; <-ticker.C {
		if !ac.wsConn.IsConnected() {
			continue
		}
		aliveMsg := websocketMessage{
			Instruction: WebsocketInstructionHello,
			Args: map[string]interface{}{
				"version": pkg.VERSION,
			},
		}
		err := ac.wsConn.WriteJSON(aliveMsg)
		ac.logger.WithField("loop", "ws-health").Trace("hello'd")
		if err != nil {
			ac.logger.WithField("loop", "ws-health").Println("write:", err)
			ac.wsConn.CloseAndReconnect()
			continue
		}
	}
}
