package twitterv2

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/2/users/me", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"screen_name": "xkcdComic"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"name": "XKCD Comic", "favourites_count": 2}`)
	})

	client := NewClient(httpClient)
	user, _, err := client.Users.Me(&UserMeParams{})
	expected := &User{Name: "XKCD Comic", FavouritesCount: 2}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}
