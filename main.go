package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Jidnahn/blog-aggregator/internal/commands"
	"github.com/Jidnahn/blog-aggregator/internal/config"
	"github.com/Jidnahn/blog-aggregator/internal/database"
	"github.com/Jidnahn/blog-aggregator/internal/middlewares"

	_ "github.com/lib/pq"
)

func main() {
	// get arguments
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Not enought arguments given")
		os.Exit(1)
	}
	// get config and set state
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	// open db connection
	db, err := sql.Open("postgres", c.Connection)
	if err != nil {
		fmt.Println("Error opening db connection:", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	state := commands.State{Config: c, Db: dbQueries}
	// register commands
	cmds := commands.Commands{
		Handlers: make(map[string]func(*commands.State, commands.Command) error),
	}
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerUsers)
	cmds.Register("agg", commands.HanlderAgg)
	cmds.Register("addfeed", middlewares.MiddlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("feeds", commands.HandlerFeeds)
	cmds.Register("follow", middlewares.MiddlewareLoggedIn(commands.HandlerFollow))
	cmds.Register("following", middlewares.MiddlewareLoggedIn(commands.HandlerFollowing))
	cmds.Register("unfollow", middlewares.MiddlewareLoggedIn(commands.HandlerUnfollow))
	// create command
	cmdName := args[1]
	loginCmd := commands.Command{
		Name: cmdName,
		Args: args[2:],
	}
	// run command
	if err := cmds.Run(&state, loginCmd); err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
}
