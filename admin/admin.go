package admin

import (
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/user"
)

// Handler admin
type Handler struct {
	Dao Dao
}

// DeleteUser deletes a user
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	h.Dao.DeleteUser(&user.User{})
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("user deleted"))
}
