package middlewares

import (
	"context"
	"fmt"

	"github.com/Jidnahn/blog-aggregator/internal/commands"
	"github.com/Jidnahn/blog-aggregator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *commands.State, cmd commands.Command, user database.User) error) func(*commands.State, commands.Command) error {
	// function that gets the current user from the DB and passes it to any handler that requires it
	resultingHandler := func(s *commands.State, cmd commands.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.Current_user_name)
		if err != nil {
			return fmt.Errorf("error getting current user: %w", err)
		}

		return handler(s, cmd, user)
	}

	return resultingHandler
}
