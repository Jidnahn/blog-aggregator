Hello, this is a CLI app to scrape RSS feeds and save them into a PostgesQL database where you can browse the feeds at anytime!

To run the program you will need `Postgres` as well as `Go` installed in your machine.

To install gatorCLI you can either:
    1. Clone the repo
    2. Navigate to the project directory
    3. Run `go install`
Or run the following command:
`go install github.com/Jidnahn/blog-aggregator@latest`

To get the command running you will need to create a `.gatorconfig.json` in your home directory, e.g: `touch ~/.gatorconfig.json` and format it as follows:
```
{
    "db_url":"postgres://localhost:5432/database_name"
    "connection":"postgresql://username:password@localhost:5432/database_name"
}
```

After all is done, you can run it using your terminal with the `gator` prefix!
Try running the command `gator register {username}` to finish cofiguring your gatorconfig file by setting the current_user_name setting.

The following commands are available
    * login
    * register {username}
    * reset (clears the database)
    * users
    * agg {time_interval in minutes} (scrapes the saved feeds according to interval)
    * addfeed {url}
    * feeds
    * follow {url}
    * following
    * unfollow {url}
    * browse {limit} (limits the amount of posts returned)

Your inital workflow should look something like:
    1. Register
    2. Add some feeds
    3. Follow feeds
    4. Run aggregator
    5. Browse posts
