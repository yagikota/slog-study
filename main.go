package main

import (
	"os"

	"golang.org/x/exp/slog"
)

// User does not implement `LogValuer` here
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// implement the `LogValuer` interface
func (u *User) LogValue() slog.Value {
    return slog.GroupValue(
        slog.String("id", u.ID),
        slog.String("name", u.FirstName+" "+u.LastName),
    )
}


func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	u := &User{
		ID:        "user-12234",
		FirstName: "Jan",
		LastName:  "Doe",
		Email:     "jan@example.com",
		Password:  "pass-12334",
	}

	logger.Info("info", "user", u)
}
