package commands

import (
	"fmt"

	"github.com/Jidnahn/blog-aggregator/internal/config"
	"github.com/Jidnahn/blog-aggregator/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("command does not exist")
	}
	if err := f(s, cmd); err != nil {
		return err
	}
	return nil
}
