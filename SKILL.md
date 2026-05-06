# AI Command Center Skill

Manage tasks/milestones via real-time JSON dashboard.

## 🛠 Resources
- `kanban.json`: Project state (Stored in the **target application's root directory**).
- `viewer.html`: Web interface.

## 📋 SOP

### 1. Init
If missing, create `kanban.json` in the **project root**:
`{"columns":["To Do","In Progress","Review","Done"],"tasks":[],"milestones":[]}`

### 2. Default Team (Compact Roles)
- **TL (Tech Lead)**: Arch, reviews, guides.
- **SA (Sys Analyst)**: Requirements, data design.
- **FE (Frontend)**: UI/UX, CSS, dashboard.
- **BE (Backend)**: API, DB (CI4), logic.
- **SEC (Security)**: BCRYPT, auth, audits.
- **OPS (DevOps)**: Envs (Node/Py/Go/PHP), Git.
- **FIX (Elite Debugger)**: Deep investigation; expert in finding/fixing complex bugs.

### 3. Updates
On every action:
1. **Read** `kanban.json` from the project root.
2. **Update** task/milestone. 
3. **Mandatory**: Set `updated_at` (YYYY-MM-DD HH:MM:SS).
4. **Log**: Add to `logs` array.
5. **Write** file back to the project root.

### 4. Serve
- Place `viewer.html` and a server script in the project root.
- Run: `node serve.js` | `python serve.py` | `go run serve.go` | `php -S localhost:8080`
- URL: `http://localhost:8080/viewer.html`

## 📐 Data Schema

### Task
`{id, title, description, assignee(Short Name), status, updated_at, logs:[]}`

### Milestone
`{id, title, progress(0-100), status, due_date}`

## 🤖 Agent Instructions
1. **New Objective?** Add task to `kanban.json` in project root.
2. **Starting?** Move to "In Progress".
3. **Finished?** Move to "Done" after validation.
4. **Matrix?** Always set a clear `assignee` short name.
