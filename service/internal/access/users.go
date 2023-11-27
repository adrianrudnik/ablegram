package access

import (
	"github.com/google/uuid"
	"sync"
)

// User represents a configured end-user of the system. This is based on the unique ID given
// after saying hello via POST /api/auth. The role is controlled by the system. The display name
// can be configured by the user.
type User struct {
	UserID      uuid.UUID `json:"id"`
	DisplayName string    `json:"display_name"`
	Role        string    `json:"role"`
}

func NewUser(userID uuid.UUID, displayName string, role string) *User {
	return &User{
		UserID:      userID,
		DisplayName: displayName,
		Role:        role,
	}
}

// UserList represents a list of users.
// Currently used to keep a list of all connected users, based on their websocket connection.
type UserList struct {
	mutex sync.RWMutex
	users map[uuid.UUID]*User
}

func NewUserList() *UserList {
	return &UserList{
		users: make(map[uuid.UUID]*User),
	}
}

func (s *UserList) All() []User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var ul []User

	for _, users := range s.users {
		ul = append(ul, *users)
	}

	return ul
}

func (s *UserList) Add(u *User) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.users[u.UserID] = u
}

func (s *UserList) Get(id uuid.UUID) *User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if user, ok := s.users[id]; ok {
		return user
	}

	return nil
}

func (s *UserList) Delete(id uuid.UUID) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.users, id)
}
