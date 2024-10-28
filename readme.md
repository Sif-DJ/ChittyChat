To run this program first open up a command promt and navigate to the server folder.
Then run "go run .\server.go" then you may have to give it access as your firewall may detect it

Now open another command promt navigate to the client folder.
Then run "go run .\client.go" this makes a client that has access to the server.'
To make more clients simply do the last 2 steps over again.

When a client has acces to a server if you write "exit" and press enter the client will leave the server and shut down.
If you write anything else and press enter the client will publish that to the server if it is valid.

Valid meaning a sentence which is equal to or shorter than 128 characters