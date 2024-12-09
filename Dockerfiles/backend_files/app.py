from flask import Flask, request, jsonify
from flask_cors import CORS
import random

app = Flask(__name__)

CORS(app)

# Global variables to hold the secret number and guess count
secret_number = random.randint(1, 100)
guess_count = 0

@app.route('/guess', methods=['POST'])
def guess():
    global guess_count, secret_number
    try:
        # Get the guess from the frontend
        data = request.get_json()
        guess = data['guess']
        
        guess_count += 1

        # Check if the guess is correct, too high, or too low
        if guess < secret_number:
            return jsonify({"message": "Too low!", "guess_count": guess_count})
        elif guess > secret_number:
            return jsonify({"message": "Too high!", "guess_count": guess_count})
        else:
            # Correct guess
            return jsonify({"message": f"Correct! It took you {guess_count} guesses.", "guess_count": guess_count})
    except KeyError:
        return jsonify({"error": "Invalid input. Please send a valid 'guess' field."}), 400

@app.route('/reset', methods=['POST'])
def reset_game():
    global secret_number, guess_count
    secret_number = random.randint(1, 100)
    guess_count = 0
    return jsonify({"message": "Game reset. Start guessing!"})

if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0", port=5000)

