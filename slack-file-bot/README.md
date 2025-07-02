
# 📁 Slack File Bot

`slack-file-bot` is a simple Slack app written in Go that connects to your Slack workspace via **Socket Mode**, listens for messages and file events, and interacts using Slack APIs. It's designed to work with file-related actions such as uploading and reading files from channels and conversations.

---

## 🧰 Tech Stack

- **Language**: Go
- **Slack SDK**: [`github.com/slack-go/slack`](https://github.com/slack-go/slack)
- **Connection**: Socket Mode via App-Level Token

---

## 🚀 Features

- Uses **Socket Mode** to receive real-time Slack events
- Handles file operations like **reading**, **writing**, and **sharing files**
- Supports public and private conversations
- Easy to extend for more commands and automation

---

## 📂 Project Structure

```

slack-file-bot/
├── main.go
├── go.mod
└── go.sum

````

---

## 🛠️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/codebyankita/slack-file-bot.git
cd slack-file-bot
````

### 2. Initialize Go Modules

```bash
go mod tidy
```

---

## 🔧 Slack App Setup (Step-by-Step)

### 📌 Step 1: Create App

1. Visit: [https://api.slack.com/apps](https://api.slack.com/apps)
2. Click **Create New App** → **From Scratch**
3. App Name: `file-bot`
4. Pick your development workspace

---

### 🔌 Step 2: Enable Socket Mode

1. In your app's dashboard, go to **Socket Mode**
2. Toggle `Enable Socket Mode` to **ON**
3. Click **Generate Token**

   * Token Name: `socket-token`
   * Scope: `connections:write`
4. Copy the **App-Level Token** (starts with `xapp-...`)

---

### 🎯 Step 3: Add OAuth Scopes

Go to **OAuth & Permissions** and add the following under **Bot Token Scopes**:

| Scope                | Description                             |
| -------------------- | --------------------------------------- |
| `channels:read`      | View public channels info               |
| `chat:write`         | Send messages as the bot                |
| `files:read`         | Read files from channels and DMs        |
| `files:write`        | Upload/edit/delete files                |
| `im:read`            | View info about DMs the bot is added to |
| `im:write`           | Start/send messages to DMs              |
| `remote_files:read`  | View remote files added by the app      |
| `remote_files:write` | Add/edit/delete remote files            |
| `remote_files:share` | Share remote files                      |

---

### ✅ Step 4: Install the App to Workspace

* Go to **Install App** tab
* Click **Install to Workspace**
* Authorize the app
* Copy the **Bot User OAuth Token** (starts with `xoxb-...`)

---

## 🧪 Running the Bot

Update your `main.go` with the following environment variables:

```go
os.Setenv("SLACK_BOT_TOKEN", "xoxb-...")
os.Setenv("SLACK_APP_TOKEN", "xapp-...")
```

### Run your bot

```bash
go run main.go
```

---

## 💬 Testing in Slack

1. Open your Slack workspace
2. Invite the bot to a channel:

```
/invite @file-bot
```

3. Interact with the bot:

   * Upload a file
   * Mention the bot
   * Watch it respond (depends on your code logic)

---

## 🔐 Best Practices

* Store tokens securely in a `.env` file and use `github.com/joho/godotenv` to load them.
* Never hardcode secrets in code or push to public repositories.

---

## 📅 Future Plans

* Respond with file metadata (file name, size, uploader)
* Automatically download and re-upload files
* Integrate cloud storage like Google Drive or IPFS

---

## 🧾 License

MIT © [Ankita Virani](https://github.com/codebyankita)

---

## 📌 References

* [Slack API Docs](https://api.slack.com/)
* [Slack Go SDK](https://github.com/slack-go/slack)
* [Socket Mode Overview](https://api.slack.com/apis/connections/socket)

```

---

Let me know once your `main.go` is ready and I’ll help you update the README further based on your bot logic (e.g., uploading or reading files).
```


