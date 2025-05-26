import os
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route("/record", methods=["POST"])
def record():
    auth_header = request.headers.get("Authorization")
    if not auth_header or auth_header != f"Bearer {os.environ.get('SECRET_TOKEN')}":
        return jsonify({"error": "Unauthorized"}), 401

    data = request.json
    name = data.get("name")
    taijyu = data.get("taijyu")
    print(f"受け取った: {name}, {taijyu}")
    return jsonify({"message": "受け取ったっちゃ！"}), 200

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=int(os.environ.get("PORT", 8080)))
