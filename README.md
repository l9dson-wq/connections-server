# Game server behavior simualtion
A simulation of a game sever or something similar IDK, I just wanted to start practicing concurrency in go.

The projects run in a localhost for now.

Available endpoints:
* /connect

You must send an ID to this endpoint in other to get a successfull connection to the "server" otherwise you will get a message indicating that the ID is missing from the URL.
