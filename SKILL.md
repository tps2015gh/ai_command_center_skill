# AI Command Center Skill (Auto-Start Enabled)

Manage tasks/milestones via real-time JSON dashboard.

## 🛠 Resources
- `kanban.json`: Project state (Stored in project root).
- `viewer.html`: Web interface.
- `serve.js/py/go`: One-click server scripts.

## 📋 SOP

### 1. Auto-Initialization (FOR AGENT)
As soon as you are prompted to use this skill:
1.  **Check/Create State**: Ensure `kanban.json` exists in project root. If not, create it: `{"columns":["To Do","In Progress","Review","Done"],"tasks":[],"milestones":[]}`.
2.  **Start Server**: Execute `node serve.js`, `python serve.py`, or `go run serve.go` in the background if no server is running on port 8080.
3.  **Open Dashboard**: Use your browser tool to open `http://localhost:8080/viewer.html`.

### 2. Default Team (Compact IDs)
- **TL**: Tech Lead | **SA**: Sys Analyst | **FE**: Frontend | **BE**: Backend | **SEC**: Security | **OPS**: DevOps | **FIX**: Elite Debugger.

### 3. Live Updates
On every milestone or task change:
1. **Read** project root `kanban.json`.
2. **Update** fields. **Mandatory**: Set `updated_at` (YYYY-MM-DD HH:MM:SS).
3. **Log**: Add latest action to `logs` array.
4. **Write** back to project root.

## 📐 Data Schema
- Task: `{id, title, description, assignee(ID), status, updated_at, logs:[]}`
- Milestone: `{id, title, progress(0-100), status, due_date}`

## 🤖 MANDATORY AGENT BEHAVIOR
You are the owner of this dashboard. Do NOT wait for the user to ask you to "update the board." Whenever you perform a task (coding, research, bug fix), you MUST autonomously update the `kanban.json` to reflect your current state.
