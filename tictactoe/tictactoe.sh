#! /bin/bash

# functions

function init_board () {
	# This function initiate an empty board.
	row1=("@" "1" "2" "3")
	row2=("A" "." "." ".")
	row3=("B" "." "." ".")
	row4=("C" "." "." ".")
}


function print_board () {
	# This functions print the current board.
	echo ${row1[*]}
	echo ${row2[*]}
	echo ${row3[*]}
	echo ${row4[*]}
	echo ""
}

function choose_starter () {
	# randomly choosing the starter, 1 = X, 2 = O
	min=1
	max=2
	randomized=$(($RANDOM%($max-$min+1)+$min))
	if [[ $randomized -eq 1 ]]; then
        	echo "X"
	elif [[ $randomized -eq 2 ]]; then
        	echo "O"
	fi
}

function rotate_player () {
	# This function changes the turn to the other player.
	# Gets the which player took a turn and returns the other player.
	if [[ $1 == "X" ]]; then
		echo "O"
	elif [[ $1 == "O" ]]; then
		echo "X"
	fi
}


function make_move () {
	# This function makes a move.
	# Gets the user's move as the first input and which player's turn is it as the second argument.
	loc=${1:1:1}
	if [[ ${1:0:1} == "A" ]]; then
		row2[$loc]=$2
	elif [[ ${1:0:1} == "B" ]]; then
		row3[$loc]=$2
	elif [[ ${1:0:1} == "C" ]]; then
		row4[$loc]=$2
	fi
}

function check_taken_square () {
	# Checks if the chosen square is already taken.
	# Gets the user's move as an input.
	loc=${1:1:1}
	if [[ ${1:0:1} == "A" ]]; then
                if [ ${row2[$loc]} == "." ]; then
                        # Empty square, G is good.
			echo "G"
                else
                        # Taken sqaure, B is bad.
			echo "B"
                fi
        elif [[ ${1:0:1} == "B" ]]; then
                if [ ${row3[$loc]} == "." ]; then
                        echo "G"
                else
                        echo "B"
                fi
        elif [[ ${1:0:1} == "C" ]]; then
                if [ ${row4[$loc]} == "." ]; then
                        echo "G"
                else
                        echo "B"
                fi
        fi
}

function check_move_bounds () {
	# Checks that the chosen move is valid (A/B/C and 1/2/3).
	# Gets the user's move as an input.
	loc=${1:1:1}
	col=${1:0:1}
	if [ $loc -ge 1 ] && [ $loc -le 3 ] && { [ $col == "A" ] || [ $col == "B" ] || [ $col == "C" ]; }; then
		# Valid move, G is good.
		echo "G"
	else
		# Out of bounds, B is bad.
		echo "B"
	fi
}


function check_win () {
	# Check the board for a winner in all possible combinations.
	# First row condition.
	if [ ${row2[1]} == ${row2[2]} ] && [ ${row2[2]} == ${row2[3]} ] && [ ${row2[1]} != "." ]; then
		echo ${row2[1]}
	# Second row condition.
	elif [ ${row3[1]} == ${row3[2]} ] && [ ${row3[2]} == ${row3[3]} ] && [ ${row3[1]} != "." ]; then
                echo ${row3[1]}
	# Third row condition.
	elif [ ${row4[1]} == ${row4[2]} ] && [ ${row4[2]} == ${row4[3]} ] && [ ${row4[1]} != "." ]; then
                echo ${row4[1]}
	# First column condition.	
	elif [ ${row2[1]} == ${row3[1]} ] && [ ${row3[1]} == ${row4[1]} ] && [ ${row2[1]} != "." ]; then
                echo ${row2[1]}
	# Second column condition.
	elif [ ${row2[2]} == ${row3[2]} ] && [ ${row3[2]} == ${row4[2]} ] && [ ${row2[2]} != "." ]; then
                echo ${row2[2]}
	# Third column condition.
	elif [ ${row2[3]} == ${row3[3]} ] && [ ${row3[3]} == ${row4[3]} ] && [ ${row2[3]} != "." ]; then
                echo ${row2[3]}
	# Top left to bottom right diagonal (\).
	elif [ ${row2[1]} == ${row3[2]} ] && [ ${row3[2]} == ${row4[3]} ] && [ ${row2[1]} != "." ]; then
                echo ${row2[1]}
	# Bottom left to top right diagonal (/).
	elif [ ${row2[3]} == ${row3[2]} ] && [ ${row3[2]} == ${row4[1]} ] && [ ${row2[3]} != "." ]; then
                echo ${row2[3]}
	# Tie condition
	elif [ ${row2[1]} != "." ] && [ ${row2[2]} != "." ] && [ ${row2[3]} != "." ] && \
		 [ ${row3[1]} != "." ] && [ ${row3[2]} != "." ] && [ ${row3[3]} != "." ] && \
		 [ ${row4[1]} != "." ] && [ ${row4[2]} != "." ] && [ ${row4[3]} != "." ]; then
				echo "T"
	# No winner yet
	else
		echo "L"
	fi
}


# Main code

# Initiate game
init_board
player_turn=$(choose_starter)
echo "Player "$player_turn" starts."

# Play the game
while [ $(check_win) == "L" ]
do
	print_board
	read -p "please choose a location (for example, A1 for the left top corner): " move
	while [ $(check_move_bounds $move) == "B" ]
        do
                read -p "Move is out of bounds (A/B/C and 1/2/3 only), please try again: " move
        done
	while [ $(check_taken_square $move) == "B" ]
	do
		read -p "This square is already taken, please try again: " move
	done
	make_move $move $player_turn > /dev/null
	player_turn=$(rotate_player $player_turn)
done

# Find out the result
if [ $(check_win) == "T" ]; then
	echo "It's a tie!"
else
	echo "Player "$(check_win)" won!"
fi
print_board

