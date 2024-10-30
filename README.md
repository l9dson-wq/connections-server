# Game server behavior simualtion
A simulation of a game sever or something similar IDK, I just wanted to start practicing concurrency in go.

The project run in a localhost for now.

Available endpoints:
* /connect

You must send an ID to this endpoint in order to get a successfull connection to the "server" otherwise you will get a message indicating that the ID is missing from the URL.

You should be able to see something similar like this in the console, indicanting a "Player has joined the session":
![image](https://github.com/user-attachments/assets/f2668a11-2447-40de-9c85-ef3d460db81b)
