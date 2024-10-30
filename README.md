# Game server behavior simualtion
A simulation of a game sever or something similar IDK, I just wanted to start practicing concurrency in go.

The project run in a localhost for now.

Available endpoints:
* /connect?id=[PLAYERID]
* /disconnect?id=[PLAYERID]

You must send the id in both of the endpoints, rember to change "[PLAYERID]" with a number.

You should be able to see something similar like this in the console, indicanting a "Player has joined the session":
* connect
![image](https://github.com/user-attachments/assets/f2668a11-2447-40de-9c85-ef3d460db81b)

* disconnect
![image](https://github.com/user-attachments/assets/032d9ce1-cba8-4c84-92ce-81a7b23939e0)

I added a function to simulate connection and disconnections (not working yet) to the server.
The logic for the disconnection runs alongside with the server thanks to the gorutines and I set a time of 10 seconds to wait the response from the server in case the server takes more than that to start or return a response.

![image](https://github.com/user-attachments/assets/dfe4af91-98ce-4632-ba7a-54c2398b59f2)
