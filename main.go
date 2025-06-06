package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/JCien/gator/internal/config"
	"github.com/JCien/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	myCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	myCommands.register("login", handlerLogin)
	myCommands.register("register", handlerRegister)
	myCommands.register("reset", handlerReset)
	myCommands.register("users", handlerListUsers)
	myCommands.register("agg", handlerAgg)
	myCommands.register("feeds", handlerListFeeds)
	myCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	myCommands.register("follow", middlewareLoggedIn(handlerFollow))
	myCommands.register("following", middlewareLoggedIn(handlerFollowing))
	myCommands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	myCommands.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = myCommands.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
