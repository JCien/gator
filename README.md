# Gator RSS Feed Aggregator

You will need to have Postgres and Go installed to run the program.

Install Gator by running the following in the root of the project:
```
go install
```

Create a json config file named .gatorconfig.json with the following inside:
```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}

```

Usage:
```
gator <command> <params>

commands:
agg <time_between_reqs>
- This scrapes the rss sites provided from the addfeed command.

register <name>
- Registers a user with the given name and sets them as the current user

login <name>
- Switches current user to the specified name. Name must have been previously registered.

users
- Lists all the registered users and shows the currently logged in user.

addfeed <feed_name> <url>
- Adds an rss feed from a given url.

feeds
- Lists all the added feeds for the currently logged in user.

browse <post_limit>
- Displays info of the post including a description.

follow <feed_url>
- Follows a feed added by other users.

following
- Lists the names of the feeds the current user is following.

unfollow <feed_url>
- Unfollows a feed from the current user.
```
