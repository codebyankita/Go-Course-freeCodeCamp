
# 📁 Slack File Bot

`slack-file-bot` is a simple and extensible Slack app written in Go that connects to your Slack workspace via **Socket Mode**, listens for messages or file upload triggers, and interacts with channels using Slack’s File APIs. It demonstrates how to **upload files**, respond to events, and scale interactions securely using `.env` for sensitive data.

---

## 🧰 Tech Stack

- **Language**: Go
- **Slack SDK**: [`github.com/slack-go/slack`](https://github.com/slack-go/slack)
- **Environment Variables Loader**: [`github.com/joho/godotenv`](https://github.com/joho/godotenv)
- **Connection**: Socket Mode via App-Level Token & Web API

---

## 🚀 Features

- 🔗 Uses **Socket Mode** or Slack API for direct interaction
- 📁 Uploads files to Slack channels securely
- 💬 Supports basic message automation or replies
- 🔒 Reads `.env` for tokens & settings, never hardcodes secrets
- 🧱 Easily extendable for more Slack actions

---

## 📂 Project Structure

```

slack-file-bot/
├── .env              # Secrets & config (ignored from git)
├── main.go           # Core app logic
├── go.mod            # Module definition
├── go.sum            # Dependency versions
└── ZIPL.pdf          # Sample file used for upload test

````

---

## 🛠️ Getting Started

### 🔁 Step 1: Clone the Repository

```bash
git clone https://github.com/codebyankita/slack-file-bot.git
cd slack-file-bot
````

### ⚙️ Step 2: Create `.env` File

Create a `.env` file in the root folder with the following:

```env
SLACK_BOT_TOKEN=xoxb-your-bot-token
SLACK_APP_TOKEN=xapp-your-app-token
CHANNEL_ID=channel-id-to-upload
```

**Note**: Never commit this `.env` file to GitHub. It's listed in `.gitignore` already.

---

### 📦 Step 3: Install Dependencies

```bash
go mod tidy
```

---

## 🔧 Slack App Setup (Step-by-Step)

### 📌 Step 1: Create Your Slack App

1. Go to [https://api.slack.com/apps](https://api.slack.com/apps)
2. Click **Create New App** → **From Scratch**
3. Set app name: `file-bot`
4. Choose your workspace

---

### 🔌 Step 2: Enable Socket Mode *(Optional)*

If you're planning real-time events like `message` or `file_shared`:

1. In the app's dashboard, go to **Socket Mode**
2. Toggle `Enable Socket Mode` to **ON**
3. Generate a new token (starts with `xapp-...`) and copy it
4. Add the `connections:write` scope

---

### 🔑 Step 3: Add Bot Scopes

Go to **OAuth & Permissions** > **Bot Token Scopes** and add:

| Scope                | Description                        |
| -------------------- | ---------------------------------- |
| `files:write`        | Upload/edit/delete files           |
| `files:read`         | Read files                         |
| `chat:write`         | Send messages as bot               |
| `channels:read`      | View public channels info          |
| `im:read`            | View DMs info                      |
| `im:write`           | Send messages in DMs               |
| `remote_files:read`  | Read remote files added by the bot |
| `remote_files:write` | Add/edit/delete remote files       |
| `remote_files:share` | Share remote files                 |

---

### ✅ Step 4: Install App to Workspace

1. Go to **Install App** tab
2. Click **Install to Workspace**
3. Authorize the app
4. Copy the **Bot User OAuth Token** (starts with `xoxb-...`)
5. Paste both tokens in your `.env`

---

## 🧪 Running and Testing the Bot

### 🏃 Step 1: Run the Bot

```bash
go run main.go
```

If successful, you will see output like:

```bash
Uploaded File -> Name: ZIPL.pdf, ID: F01ABCDEFGH
```

---

### 💬 Step 2: Invite the Bot in Slack

In your Slack workspace:

```bash
/invite @file-bot
```

Then try uploading a file or triggering interactions (like sending a message, depending on the logic you add).

---

## 🧰 How the Code Works (main.go)

1. Loads the `.env` file using `godotenv`
2. Initializes the Slack client using `slack-go`
3. Prepares a list of files (currently `ZIPL.pdf`)
4. Loops through them and:

   * Opens the file
   * Uses `UploadFileV2` to upload it to the channel specified in `.env`
5. Logs the upload result to the terminal

You can update `fileArr := []string{}` to include more files.

---

## 🔐 Security Best Practices

* Store all secrets in `.env`
* Use `.gitignore` to prevent secrets from being pushed
* Rotate your Slack tokens periodically
* Avoid hardcoding any workspace-specific data

---

## 📅 Future Plans

* Add Slack file metadata replies (uploader, file type, etc.)
* Auto-download and re-share files in other channels
* Build a web dashboard for Slack uploads
* Integrate with Google Drive, Dropbox, or IPFS


---

## 📌 References

* [Slack API Docs](https://api.slack.com/)
* [Slack Go SDK](https://github.com/slack-go/slack)
* [godotenv](https://github.com/joho/godotenv)
* [Socket Mode Overview](https://api.slack.com/apis/connections/socket)

```

