This project took me a bit longer than I wanted to spend on it (I think I spent ~16 hours), mostly because I went a little crazy with the backend systems. I'm happy with the outcome, but I could have done a better job in restricting what I do on the backend as that just makes implementing the client take a little bit longer. Please feel free to ping me with any questions or issues.

Thanks,
Austin

NOTES
-----------------------------------------------------------------------
- Google was my friend with things I don't often use with Go.
- I haven't done much with Go's http package since most of my work is with AWS lambda, so that was a little bit of a slow down.
- A tiny research went into reading/writing files
- Since Go's http library doesn't support path parameters (GetUserInventory - /users/{userId}/inventory), for the sake of time I hacked around it a bit by having other routes with query parameters (GetUserInventory - /inventory?id=<userId>)
- I tried creating a simple matchmaker. It doesn't consider separating players by skill, board size/type, or anything else.
- I know you wanted me to stick with stdlib, but I imported 1 external package for UUIDs :)
- Long polling works, but websocket implementation would be a much beter/cleaner route.
- Some endpoints are not hack resilient due to time constraints. CreateGame, for example, will let you send a whole game object instead of specific fields allowed from the client.
- Could have taken more time to make the databases a bit "persistant" by saving the data to a txt file
- I also have TODO notes dispersed around the code base for things I could improve at a later date
- I REALLY wanted to spend more time to clean up the client, but I've decided to pass on that and get the project submitted


FLOW
-----------------------------------------------------------------------
Login (get user or create new)
Queue for match (returns a ticket)
Long poll for ticket status update (return ticket)
Join
Long poll game data
Make game actions when it's your move.

COMMANDS
-----------------------------------------------------------------------
-users: lists users that are stored in local mem so you can log in as existing user. NOTE: backend doesn't have persistance so this wont work if the backend is taken down.

-login: logs in as new users

-login=[id]: logs in as an existing user

-queue: queues up for a match, generating a ticket

-join: joins a match based on current ticket stored in mem

-action=<boardIntVal>: submits an action to the backend to play on a specific board piece


EXAMPLE SETUP
-----------------------------------------------------------------------
go run backend.go

go run client.go (x2 for two clients)

client:
-login (x2)

-queue (x2)

-join (x2)

-action=[int] (depending on which player turn it is - The board displays integer values that you can use for making an action)