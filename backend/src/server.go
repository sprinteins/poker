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

func NewServer() Server {

	srv := Server{
		store:     store.New(),
		actionMap: make(ActionMap, 0),
	}

	srv.actionMap[store.TypeStartVoting] = srv.StartVoting
	srv.actionMap[store.TypeReveal] = srv.Reveal
	srv.actionMap[store.TypeReset] = srv.Reset
	srv.actionMap[store.TypeChoose] = srv.Choose

	return srv
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		resp := TestResp{
			MSG: "hallo",
		}
		c.JSON(200, resp)
	})

	r.GET("/poker/api/room/:roomid/:playername", func(c *gin.Context) {
		s.wshandler(c.Writer, c.Request, c)
	})

	r.Run("localhost:7788")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (s Server) wshandler(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	roomID := c.Param("roomid")
	playername := c.Param("playername")

	log.Printf("roomid=%s playername=%s", roomID, playername)

	player := store.Player{
		ID:   uuid.NewString(),
		Name: playername,
		Conn: conn,
	}

	s.store.AddPlayer(roomID, player)
	s.ListenUntilDesctonnects(conn, roomID, player)
	s.store.RemovePlayer(roomID, player)
}

func (s Server) ListenUntilDesctonnects(
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

func (s Server) StartVoting(roomID types.ID, _ types.ID, _ []byte) {
	s.store.StartVoting(roomID)
}

func (s Server) Reveal(roomID types.ID, _ types.ID, _ []byte) {
	s.store.Reveal(roomID)
}

func (s Server) Reset(roomID types.ID, _ types.ID, _ []byte) {
	s.store.Reset(roomID)
}

func (s Server) Choose(roomID types.ID, playerID types.ID, payload []byte) {
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
