# 🧠 Slack Age Bot

A simple Slack bot built with Go using the [`slacker`](https://github.com/shomali11/slacker) framework. This bot calculates a user's age when they send a message like:

```

my yob is 1995

```

and responds with:

```

Your age is 30

````

---

## 🚀 Features

- Calculates age based on the input year
- Handles invalid input gracefully
- Listens for real-time Slack messages via the Slack Events API
- Logs incoming commands with timestamp and parameters

---

## 🛠️ Tech Stack

- Language: **Go**
- Bot Framework: **slacker**
- Slack App: **Socket Mode** & **Bot Token Authentication**

---

## 📦 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/slack-age-bot.git
cd slack-age-bot
````

### 2. Install Dependencies

Make sure you have Go installed (`go version`).

Then run:

```bash
go mod tidy
```

---

## 🔐 Slack App Configuration

### Step-by-step guide to creating a Slack App

1. **Go to Slack API Console**:
   [https://api.slack.com/apps](https://api.slack.com/apps)

2. **Create a New App**

   * Choose: `From scratch`
   * App name: `age-bot`
   * Workspace: Select your Slack workspace

3. **Enable Socket Mode**

   * Go to **Settings > Socket Mode**
   * Enable it
   * Create an App-Level Token with `connections:write` scope
   * Save the token (`xapp-...`)

4. **Bot Token Scopes**
   Navigate to **OAuth & Permissions**, under **Bot Token Scopes**, add the following:

   * `app_mentions:read`
   * `chat:write`
   * `chat:write.public`
   * `channels:read`
   * `channels:history`
   * `im:read`
   * `im:write`
   * `im:history`
   * `mpim:read`
   * `mpim:write`
   * `mpim:history`

5. **Install App to Workspace**

   * Click on **Install App**
   * Copy your **Bot User OAuth Token** (`xoxb-...`)

---

## 🧪 Testing the Bot

### Add your secrets as environment variables

Replace your tokens in `main.go` with `os.Getenv(...)`:

```go
os.Setenv("SLACK_BOT_TOKEN", "xoxb-...")
os.Setenv("SLACK_APP_TOKEN", "xapp-...")
```

Or even better, use a `.env` file and the `godotenv` package.

---

### Run the Bot

```bash
go run main.go
```

---

### Try It in Slack

1. **Invite the Bot to a Channel**

Type in any Slack channel:

```
/invite @age-bot
```

2. **Trigger the Command**

Say:

```
@age-bot my yob is 1995
```

💬 The bot will reply:

```
Your age is 30
```

---

## 🧾 Command Example

| Command Pattern    | Description    | Example Input    | Bot Reply        |
| ------------------ | -------------- | ---------------- | ---------------- |
| `my yob is <year>` | Age calculator | `my yob is 2000` | `Your age is 25` |

---

## 🧼 Cleanup Tokens (Optional but Recommended)

Instead of hardcoding tokens in your code, add them via environment variables or use `.env`:

```bash
export SLACK_BOT_TOKEN="xoxb-..."
export SLACK_APP_TOKEN="xapp-..."
```

---

## 💡 Future Improvements

* Automatically fetch current year (`time.Now().Year()`)
* Add more commands (e.g., `days until birthday`)
* Add user-specific analytics

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

## ✨ Made with ❤️ by [Ankita Virani](https://github.com/codebyankita)

```

---

Let me know if you'd like the `.env` file support added or want this published to GitHub with a CI setup!
```
