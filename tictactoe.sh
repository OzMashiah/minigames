#! /bin/bash

# functions

function init_board () {
	row1=("@" "1" "2" "3")
	row2=("A" "." "." ".")
	row3=("B" "." "." ".")
	row4=("C" "." "." ".")
}


function print_board () {
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
	if [[ $1 == "X" ]]; then
		echo "O"
	elif [[ $1 == "O" ]]; then
		echo "X"
	fi
}


function make_move () {
	loc=${1:1:1}
	if [[ ${1:0:1} == "A" ]]; then
		if [ ${row2[$loc]} == "." ]; then
			row2[$loc]=$2
			echo "G"
		else
			echo "!"
		fi
	elif [[ ${1:0:1} == "B" ]]; then
		if [ ${row3[$loc]} == "." ]; then
			row3[$loc]=$2
			echo "G"
		else
			echo "!"
		fi
	elif [[ ${1:0:1} == "C" ]]; then
		if [ ${row4[$loc]} == "." ]; then
			row4[$loc]=$2
			echo "G"
		else
			echo "!"
		fi
	fi
}

function check_taken_square () {
	loc=${1:1:1}
	if [[ ${1:0:1} == "A" ]]; then
                if [ ${row2[$loc]} == "." ]; then
                        echo "G"
                else
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
	loc=${1:1:1}
	col=${1:0:1}
	if [ $loc -ge 1 ] && [ $loc -le 3 ] && { [ $col == "A" ] || [ $col == "B" ] || [ $col == "C" ]; }; then
		echo "G"
	else
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

# initiate empty board
init_board
player_turn=$(choose_starter)
echo "Player "$player_turn" starts."

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

if [ $(check_win) == "T" ]; then
	echo "It's a tie!"
else
	echo "Player "$(check_win)" won!"
fi
print_board

