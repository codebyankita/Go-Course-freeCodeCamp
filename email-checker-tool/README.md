# 📧 Email Domain Checker (Go CLI Tool)

This is a simple command-line tool built in **Go** that checks the email domain's configuration for **MX**, **SPF**, and **DMARC** records. It helps validate if an email domain is correctly set up to handle sending and receiving emails securely.

---

## 🚀 Features

- ✅ Checks if the domain has **MX records** (Mail Servers)
- ✅ Checks for **SPF** (Sender Policy Framework) configuration
- ✅ Checks for **DMARC** (Domain-based Message Authentication) configuration
- 🧾 Provides details of **SPF** and **DMARC** records if available
- 🧑‍💻 Clean and user-friendly terminal output
- 🛡️ Helps in identifying spoofing/misconfiguration issues in emails

---

## 📦 Tech Stack

- **Language**: [Go](https://golang.org/)
- **Standard Library**:
  - `net` for DNS lookups (MX, TXT)
  - `bufio`, `os`, `strings`, `log` for input/output and error handling

---

## 📁 Project Structure

```

email-checker-tool/
├── main.go         # Main CLI application
├── README.md       # Project documentation

````

---

## 🛠️ How It Works

When you input an email like `someone@example.com`, this tool:

1. Extracts the **domain** (`example.com`)
2. Checks if:
   - The domain has an MX record (mail server to receive emails)
   - The domain has an SPF record (used to prevent email spoofing)
   - The domain has a DMARC record (tells receivers what to do if SPF or DKIM fails)
3. Displays a clear ✅/❌ result for each check

---

## 🧪 Sample Output

```bash
$ go run main.go
ankitacode11@gmail.com

📬 Checking domain: gmail.com
✔️  Has MX (Mail Server): YES
✔️  Has SPF (Anti-Spoofing): YES
    ↪️ SPF Record: v=spf1 redirect=_spf.google.com
✔️  Has DMARC: YES
    ↪️ DMARC Record: v=DMARC1; p=none; sp=quarantine; rua=mailto:mailauth-reports@google.com
````

---

## ▶️ How to Run

### 1. Clone the Repository

```bash
git clone https://github.com/codebyankita/email-checker-tool.git
cd email-checker-tool
```

### 2. Run the Go App

```bash
go run main.go
```

### 3. Enter Email(s)

Type a full email address and press **Enter**:

```bash
someone@example.com
```

Repeat for more emails.
Use **Ctrl+D** (on macOS/Linux) or **Ctrl+Z** (on Windows) to stop input.

---

## ❓ Why These Records Matter

| Record    | Purpose                                                                |
| --------- | ---------------------------------------------------------------------- |
| **MX**    | Defines mail servers responsible for accepting email on the domain.    |
| **SPF**   | Helps prevent spoofing by specifying allowed sending servers.          |
| **DMARC** | Works with SPF/DKIM to set policies for handling unauthenticated mail. |

---

## 🔐 Notes

* This tool **only validates domain configurations**, it does not send actual emails.
* Works for **valid email addresses only** (`name@domain.com`). Invalid formats will be skipped.
* Requires an **internet connection** to perform DNS lookups.

---

## ✅ Good For

* System Administrators
* Developers testing email domain configs
* Email marketers verifying their domain security
* Anyone trying to diagnose email delivery or spoofing issues

---

## 🙌 Contributions

Feel free to fork, improve, and contribute to this project via PRs or issues!

---

## 📚 References

* [SPF Overview](https://en.wikipedia.org/wiki/Sender_Policy_Framework)
* [DMARC Overview](https://dmarc.org/)
* [Go net package](https://pkg.go.dev/net)


