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
	loc=${1:1:2}
	if [[ ${1:0:1} == "A" ]]; then
		row2[$loc]=$2
	elif [[ ${1:0:1} == "B" ]]; then
                row3[$loc]=$2
	elif [[ ${1:0:1} == "C" ]]; then
                row4[$loc]=$2
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
	else
	# No winner yet.
		echo "L"
	fi
}

# initiate empty board
init_board
player_turn=$(choose_starter)
echo "Player "$player_turn" starts."

win_con=$(check_win)

while [ $win_con == "L" ]
do
	print_board
	read -p "please choose a location (for example, A1 for the left top corner): " move
	make_move $move $player_turn
	player_turn=$(rotate_player $player_turn)
	win_con=$(check_win)
done

winner=$(check_win)
echo "Player "$winner" won!"
print_board

