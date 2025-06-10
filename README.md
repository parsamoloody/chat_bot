# Chat bot (openai based)

A chat bot built with **OpenAI** for your assistant 🤖💀

## 1- 📦 Installation

```bash
git https://github.com/parsamoloody/chat_bot
cd chat_bot
go build
```

## 2- ⚙️ Environment Variables

set a environment variable on your machine
```bash
PORT=PORT
OPENAI_API_KEY=your_openai_api_key_here
```
## 3- ▶️ Running the Server

```bash
go run main.go
```
---

# API endpoint

POST /api/chat

Request Body:
```json
{
  "prompt": "Hello"
}
````

Response:
```json
{
    "answer":"Hello! How can I assist you today?"
}   
```

# example usage

request:
```bash
curl -X POST http://localhost:3030/api/chat \
 -H 'Content-Type: application/json' \
 -d '{"prompt":"Hello"}'
```
response:
```json
{"answer":"Hello! How can I assist you today?"}
```