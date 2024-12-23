package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Jidnahn/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("error: expected at least one argument")
	}

	id := uuid.New()
	created_at := time.Now()
	updated_at := time.Now()
	userName := strings.Join(cmd.Args, " ")

	newUser, err := s.Db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        id,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			Name:      userName,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	if err := s.Config.SetUser(userName); err != nil {
		return err
	}

	fmt.Printf("The user %s has been registered\n", userName)
	fmt.Println(newUser)

	return nil
}
