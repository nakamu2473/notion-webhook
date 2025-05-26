# main.py

from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route("/")
def home():
    return "Webhookは生きてるっちゃ！"

@app.route("/record", methods=["POST"])
def record():
    data = request.json
    name = data.get("name")
    taijyu = data.get("taijyu")
    print(f"受け取った: {name}, {taijyu}")
    return jsonify({"message": "受け取ったっちゃ！"}), 200

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
