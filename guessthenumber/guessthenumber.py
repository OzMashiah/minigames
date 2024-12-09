# import libraries
import sys
import random

# choose difficulty
if sys.argv[1] == 'Easy':
    floor = 1
    top = 10
elif sys.argv[1] == 'Medium':
    floor = 1
    top = 100
elif sys.argv[1] == 'Hard':
    floor = -500
    top = 500

# generate the number
lucky_num = random.randint(floor,top)

guess = int(input("Please guess a number between " + str(floor) + " and " + str(top) + ": "))
tries = 1
direction = "placeholder"

while guess != lucky_num:
    # add a check that a number was given between the ranges
    tries += 1
    if guess > lucky_num:
        direction = "lower than"
    elif guess < lucky_num:
        direction = "greater than"
    print("The lucky number is " + direction + " " + str(guess) + ".")
    guess = int(input("Please guess a number between " + str(floor) + " and " + str(top) + ": "))

print("The lucky number was " + str(lucky_num) + ", it took you " + str(tries) + " tries.")


