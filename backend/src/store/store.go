package store

import (
	"log"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/sprinteins/poker/src/x/types"
)

type roomID = string
type rooms = map[roomID]*Room

type Store struct {
	rooms   rooms
	roomMtx sync.Mutex
}

func New() Store {

	return Store{
		rooms: make(rooms, 0),
	}
}

func (s *Store) AddPlayer(roomID types.ID, newPlayer Player) {
	room := s.EnsureRoom(roomID)
	room.AddPlayer(newPlayer)
}

func (s *Store) RemovePlayer(roomID types.ID, player Player) {
	room := s.EnsureRoom(roomID)
	room.RemovePlayer(player)
	s.CleanUpRoomIfNeeded(room)
}

func (s *Store) CleanUpRoomIfNeeded(room *Room) {
	hasPlayers := len(room.players) > 0
	if !hasPlayers {
		delete(s.rooms, room.id)
	}
}

func (s *Store) EnsureRoom(roomID types.ID) *Room {
	s.roomMtx.Lock()
	defer s.roomMtx.Unlock()
	room, ok := s.rooms[roomID]
	if !ok {
		log.Printf("room not found createing, roomid=%s", roomID)
		newRoom := NewRoom(roomID)
		s.rooms[roomID] = &newRoom
		room = s.rooms[roomID]
	} else {
		log.Printf("room found with roomid=%s", roomID)
	}

	return room
}

func (s *Store) searchRoom(roomID types.ID) *Room {
	room, ok := s.rooms[roomID]
	if !ok {
		logrus.WithFields(logrus.Fields{
			"fn":     "StartVoting",
			"roomID": roomID,
		}).Warn("room not found")
		return nil
	}

	return room
}

func (s *Store) StartVoting(roomID types.ID) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.StartVoting()
	go room.AutoRevealIfTimeIsUp()
}

func (s *Store) Reveal(roomID types.ID) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.Reveal()
}

func (s *Store) Reset(roomID types.ID) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.Reset()
}

func (s *Store) Choose(roomID types.ID, playerID types.ID, card string) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.Choose(playerID, card)
	room.AutoRevealIfCan()
}

func (s *Store) ReVote(roomID types.ID, playerID types.ID) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.ReVote(playerID)
}

func (s *Store) SetPlayerType(roomID types.ID, playerID types.ID, playerType PlayerType) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.SetPlayerType(playerID, playerType)
}

func (s *Store) SwitchCardBack(roomID types.ID, playerID types.ID, cardback CardBack) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.SwitchCardBack(playerID, cardback)
}

func (s *Store) SetCards(roomID types.ID, playerID types.ID, cards Cards) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.SetCards(playerID, cards)
}

func (s *Store) SetAutoReveal(roomID types.ID, autoReveal bool) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.SetAutoReveal(autoReveal)
}

func (s *Store) SetAutoRevealTimer(roomID types.ID, autoRevealTimer int) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.SetAutoRevealTimer(autoRevealTimer)
}

func (s *Store) SetAdmin(roomID types.ID, newAdminID types.ID) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.SetAdmin(newAdminID)
}
