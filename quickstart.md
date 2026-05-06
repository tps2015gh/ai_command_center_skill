# Quick Start Guide: AI Command Center

Get your real-time AI coordination dashboard up and running in less than 60 seconds.

## 1. Installation
Clone the repository into your project root:
```bash
git clone https://github.com/tps2015gh/ai_command_center_skill.git
```

## 2. Initialize Data
Create `kanban.json` in your **project root directory** (the same level as your code):
```json
{"columns":["To Do","In Progress","Review","Done"],"tasks":[],"milestones":[]}
```

## 3. Launch the Dashboard
Run one of these from the project root. They will serve your project files and **automatically open your browser**.

### 🟢 Node.js
```bash
node serve.js
```

### 🐍 Python
```bash
python serve.py
```

### 🐹 Go
```bash
go run serve.go
```

### 🐘 PHP
```bash
php -S localhost:8080
# Then manually open http://localhost:8080/viewer.html
```

## 4. Connect your AI Agent
Copy the contents of **`SKILL.md`** and paste it into your AI agent's chat (Gemini, ChatGPT, Claude).

**Try this prompt:**
> "I have added the AI Command Center skill. Please add a new task 'Initial Project Setup' assigned to TL and set it to In Progress."

## 5. Watch the Magic
Keep your browser tab open. You will see the task appear instantly without refreshing!
