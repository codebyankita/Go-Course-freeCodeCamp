
# 🚀 AWS Lambda Function in Go (Golang)

This project demonstrates how to write, build, package, and deploy an AWS Lambda function in **Golang** using the **AWS CLI** (version 2+).

---

## 🗂️ Project Structure

```

lamda-yt-example/
├── go.mod               # Go module definition
├── go.sum               # Go module dependencies
├── main.go              # Lambda source code
├── main                 # Compiled Go binary
├── function.zip         # Zipped binary for Lambda deployment
├── trust-policy.json    # IAM trust policy for Lambda role

````

---

## 🧠 Lambda Code (main.go)

```go
package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"How old are you"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{
		Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age),
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
````

---

## 🛠️ Prerequisites

* ✅ [Go](https://golang.org/dl/) installed (v1.18+ recommended)
* ✅ AWS CLI v2 installed ([install guide](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html))
* ✅ AWS credentials configured using `aws configure`

---

## ⚙️ Step-by-Step Deployment Guide

### 1️⃣ Install Go Dependencies

```bash
go mod tidy
```

---

### 2️⃣ Build and Zip the Go Binary

```bash
go build main.go
zip function.zip main
```

---

### 3️⃣ Install and Verify AWS CLI (if not done)

```bash
curl "https://awscli.amazonaws.com/AWSCLIV2.pkg" -o "AWSCLIV2.pkg"
sudo installer -pkg AWSCLIV2.pkg -target /
aws --version
```

If you see `command not found`, add it to your path:

```bash
export PATH=/usr/local/bin:$PATH
source ~/.zshrc
```

---

### 4️⃣ Configure AWS Credentials

```bash
aws configure
```

You’ll be prompted to enter:

* AWS Access Key ID
* AWS Secret Access Key
* Region (e.g., `us-east-1`)
* Output format (e.g., `json`)

---

### 5️⃣ Create IAM Trust Policy File

Create a file called `trust-policy.json` with the following content:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

---

### 6️⃣ Create IAM Role for Lambda

```bash
aws iam create-role \
--role-name lambda-ex \
--assume-role-policy-document file://trust-policy.json
```

> ⚠️ If you see this error:
>
> ```
> An error occurred (EntityAlreadyExists): Role with name lambda-ex already exists.
> ```
>
> Just skip to the next step — the role is already created.

---

### 7️⃣ Attach Policy to Role

```bash
aws iam attach-role-policy \
--role-name lambda-ex \
--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

---

### 8️⃣ Create Lambda Function

```bash
aws lambda create-function \
--function-name go-lambda-function-3 \
--zip-file fileb://function.zip \
--handler main \
--runtime go1.x \
--role arn:aws:iam::<your-account-id>:role/lambda-ex
```

> 🧠 Replace `<your-account-id>` with your **real AWS Account ID**
> Example from screenshot:
> `arn:aws:iam::419784325734:role/lambda-ex`

---

## ✅ Invoke & Test Your Lambda

### 🔄 Using AWS CLI (v2 compatible)

```bash
aws lambda invoke \
--function-name go-lambda-function-3 \
--cli-binary-format raw-in-base64-out \
--payload '{"What is your name?": "Jim", "How old are you?": 33}' \
output.txt
```

✔️ Then view the output:

```bash
cat output.txt
```

Expected output:

```json
{"Answer:":"Jim is 33 years old!"}
```

---

## 🧯 Common Errors & Fixes

| ❌ Error                        | 💡 Solution                                          |
| ------------------------------ | ---------------------------------------------------- |
| `command not found: aws`       | Install AWS CLI and add to PATH                      |
| `Unable to locate credentials` | Run `aws configure`                                  |
| `EntityAlreadyExists`          | IAM role already exists — skip create, attach policy |
| CLI v2 Payload Issues          | Use `--cli-binary-format raw-in-base64-out`          |

---

## 📜 Reference Commands History

```bash
go mod tidy
go build main.go
zip function.zip main

aws configure

aws iam create-role \
--role-name lambda-ex \
--assume-role-policy-document file://trust-policy.json

aws iam attach-role-policy \
--role-name lambda-ex \
--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

aws lambda create-function \
--function-name go-lambda-function-3 \
--zip-file fileb://function.zip \
--handler main \
--runtime go1.x \
--role arn:aws:iam::419784325734:role/lambda-ex

aws lambda invoke \
--function-name go-lambda-function-3 \
--cli-binary-format raw-in-base64-out \
--payload '{"What is your name?": "Jim", "How old are you?": 33}' \
output.txt
```

---

## 📚 Resources

* [AWS Lambda for Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
* [IAM Roles for Lambda](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html)
* [Install AWS CLI v2](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)

