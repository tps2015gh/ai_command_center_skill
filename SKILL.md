# AI Command Center Skill

This skill allows an AI Agent to manage project tasks, milestones, and team workload using a real-time JSON-based dashboard.

## 🛠 Available Resources
- `kanban.json`: The source of truth for the project state.
- `viewer.html`: The real-time web interface.

## 📋 Standard Operating Procedure (SOP)

### 1. Initializing the Board
If `kanban.json` does not exist, create it using the standard schema:
```json
{
  "columns": ["To Do", "In Progress", "Review", "Done"],
  "tasks": [],
  "milestones": []
}
```

### 2. Default Team Structure
When starting a new project, the following agent roles are assumed by default. Any agent using this skill should assign tasks to these specific roles:
-   **Tech Lead**: Oversees architecture, technical guides, and code reviews.
-   **Frontend Dev**: Handles UI/UX, CSS, and real-time dashboard enhancements.
-   **Backend/System Architect**: Manages API logic, database schemas (CI4), and performance.
-   **Security Engineer**: Responsible for BCRYPT hashing, authentication, and audits.
-   **DevOps**: Manages server environments (Python/Go/PHP), Git, and deployments.

### 3. Updating Tasks
Whenever you (the agent) start, update, or finish a task:
1.  **Read** the current `kanban.json`.
2.  **Update** the relevant task object.
3.  **Mandatory Field**: Always update `updated_at` with the current timestamp (YYYY-MM-DD HH:MM:SS).
4.  **Log Milestones**: Add a descriptive log entry to the `logs` array for major progress.
5.  **Write** the file back.

### 3. Launching the Viewer
The user can view the board by serving the project directory.
-   **Command**: `php -S localhost:8080` (or similar static server).
-   **URL**: `http://localhost:8080/viewer.html`.

## 📐 Data Schema

### Task Object
- `id`: Unique integer.
- `title`: Short name.
- `description`: Detailed goal.
- `assignee`: Agent role (e.g., "Tech Lead").
- `status`: One of the strings in `columns`.
- `updated_at`: Current timestamp.
- `logs`: Array of strings (latest 2 shown on cards).

### Milestone Object
- `id`: String (e.g., "M1").
- `title`: Name.
- `progress`: Integer (0-100).
- `status`: "Planned", "In Progress", or "Completed".
- `due_date`: YYYY-MM-DD.

## 🤖 Instructions for the Agent
You are a proactive team member. Whenever a user gives you a Directive, you MUST:
1.  Add a task to the `kanban.json` if it's a new objective.
2.  Move the task to "In Progress" when you begin work.
3.  Move the task to "Done" only after successful validation.
4.  Maintain the `Team Matrix` by ensuring every task has a clear `assignee`.
