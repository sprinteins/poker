package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/sprinteins/poker/src/store"
	"github.com/sprinteins/poker/src/x/types"
)

type ActionHandlerFn func(roomID types.ID, playerID types.ID, actionAsJSON []byte)
type ActionMap map[store.ActionType]ActionHandlerFn

type Server struct {
	store     store.Store
	actionMap ActionMap
}

func NewServer() *Server {

	srv := Server{
		store:     store.New(),
		actionMap: make(ActionMap, 0),
	}

	srv.actionMap[store.TypeStartVoting] = srv.StartVoting
	srv.actionMap[store.TypeReveal] = srv.Reveal
	srv.actionMap[store.TypeReset] = srv.Reset
	srv.actionMap[store.TypeChoose] = srv.Choose
	srv.actionMap[store.TypeReVote] = srv.ReVote
	srv.actionMap[store.TypeSetPlayerType] = srv.SetPlayerType
	srv.actionMap[store.TypeSwitchCardBack] = srv.SwitchCardBack
	srv.actionMap[store.TypeSetCards] = srv.SetCards
	srv.actionMap[store.TypeSetAutoReveal] = srv.SetAutoReveal
	srv.actionMap[store.TypeSetAutoRevealTimer] = srv.SetAutoRevealTimer
	srv.actionMap[store.TypeSetAdmin] = srv.SetAdmin

	return &srv
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/api/room/:roomid/:playername", func(c *gin.Context) {
		s.wshandler(c.Writer, c.Request, c)
	})

	r.Run("localhost:7788")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (s *Server) wshandler(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	roomID := c.Param("roomid")
	playername := c.Param("playername")

	log.Printf("roomid='%s' playername='%s'", roomID, playername)

	player := store.Player{
		ID:   uuid.NewString(),
		Name: playername,
		Conn: conn,
		Type: store.PlayerTypePlayer,
	}

	s.store.AddPlayer(roomID, player)
	s.ListenUntilDesctonnects(conn, roomID, player)
	s.store.RemovePlayer(roomID, player)
}

func (s *Server) ListenUntilDesctonnects(
	conn *websocket.Conn,
	roomID types.ID,
	player store.Player,
) {
	for {
		_, body, err := conn.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}
		action := &store.Action{}
		err = json.Unmarshal(body, action)
		if err != nil {
			log.Error(err)
		}

		handleAction, ok := s.actionMap[action.Type]
		if !ok {
			log.WithFields(log.Fields{
				"type": action.Type,
				"body": fmt.Sprintf("%s", body),
			}).Warn("no handler found")
			continue
		}

		handleAction(roomID, player.ID, body)
	}
}

func (s *Server) StartVoting(roomID types.ID, _ types.ID, _ []byte) {
	s.store.StartVoting(roomID)
}

func (s *Server) Reveal(roomID types.ID, _ types.ID, _ []byte) {
	s.store.Reveal(roomID)
}

func (s *Server) Reset(roomID types.ID, _ types.ID, _ []byte) {
	s.store.Reset(roomID)
}

func (s *Server) Choose(roomID types.ID, playerID types.ID, payload []byte) {
	chooseAction := ActionChoose{}
	err := json.Unmarshal(payload, &chooseAction)
	if err != nil {
		log.WithFields(log.Fields{
			"payload": fmt.Sprintf("%s", payload),
		}).Error("could not parse choose action")
	}

	s.store.Choose(roomID, playerID, chooseAction.Payload)

}

type ActionChoose struct {
	Payload string `json:"payload"`
}

func (s *Server) ReVote(roomID types.ID, playerID types.ID, payload []byte) {
	s.store.ReVote(roomID, playerID)

}

func (s *Server) SetPlayerType(roomID types.ID, playerID types.ID, payload []byte) {
	action := ActionSetPlayerType{}
	err := json.Unmarshal(payload, &action)
	if err != nil {
		log.WithFields(log.Fields{
			"payload": fmt.Sprintf("%s", payload),
		}).Error("could not parse set-player-type action")
	}

	s.store.SetPlayerType(roomID, playerID, action.Payload)
}

type ActionSetPlayerType struct {
	Payload string `json:"payload"`
}

func (s *Server) SwitchCardBack(roomID types.ID, playerID types.ID, payload []byte) {
	action := ActionSetPlayerType{}
	err := json.Unmarshal(payload, &action)
	if err != nil {
		log.WithFields(log.Fields{
			"payload": fmt.Sprintf("%s", payload),
		}).Error("could not parse switch-card-back action")
	}

	s.store.SwitchCardBack(roomID, playerID, action.Payload)
}

type ActionSwitchCardBack struct {
	Payload string `json:"payload"`
}

func (s *Server) SetCards(roomID types.ID, playerID types.ID, payload []byte) {
	action := ActionSetCards{}
	err := json.Unmarshal(payload, &action)
	if err != nil {
		log.WithFields(log.Fields{
			"payload": fmt.Sprintf("%s", payload),
		}).Error("could not parse set-cards action")
		return
	}

	s.store.SetCards(roomID, playerID, action.Payload)
}

type ActionSetCards struct {
	Payload []string `json:"payload"`
}

func (s *Server) SetAutoReveal(roomID types.ID, playerID types.ID, payload []byte) {
	action := ActionSetAutoReveal{}
	err := json.Unmarshal(payload, &action)
	if err != nil {
		log.WithFields(log.Fields{
			"payload": fmt.Sprintf("%s", payload),
		}).Error("could not parse set-auto-reveal action")
		return
	}

	s.store.SetAutoReveal(roomID, action.Payload)
}

type ActionSetAutoReveal struct {
	Payload bool `json:"payload"`
}

func (s *Server) SetAutoRevealTimer(roomID types.ID, playerID types.ID, payload []byte) {
	action := ActionSetAutoRevealTimer{}
	err := json.Unmarshal(payload, &action)
	if err != nil {
		log.WithFields(log.Fields{
			"payload": fmt.Sprintf("%s", payload),
		}).Error("could not parse set-auto-reveal-timer action")
		return
	}

	s.store.SetAutoRevealTimer(roomID, action.Payload)
}

type ActionSetAutoRevealTimer struct {
	Payload int `json:"payload"`
}

func (s *Server) SetAdmin(roomID types.ID, playerID types.ID, payload []byte) {
	action := ActionSetAdmin{}
	err := json.Unmarshal(payload, &action)
	if err != nil {
		log.WithFields(log.Fields{
			"payload": fmt.Sprintf("%d", payload),
		}).Error("could not parse set-admin action")
		return
	}

	s.store.SetAdmin(roomID, action.Payload)
}

type ActionSetAdmin struct {
	Payload types.ID `json:"payload"`
}
