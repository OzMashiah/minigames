# minigames
This repository contains the following minigames that can be played from your bash shell:
* Guess The Number written in python with three difficulties [Easy, Medium, Hard]. 
    Should be ran using 'python3 guessthenumber.py <difficulty>', for example, python3 guessthenumber.py Medium.
* Tic Tac Toe written in bash.
    Should be ran using './tictactoe.sh'
* Submarines written in golang.
    Should be ran using './submarines'

This repository also contains an option to play Guess The Number on a browser.
This option is built on containers, with a python backend and nginx frontend.
The Dockerfiles directory contains the backend's and frontend's dockerfiles and the rest of the code.
To run this use 'sudo docker compose up -d --build' from the main directory, and head over to 'http://localhost:8080'.
