# AI Command Center: Professional Agent Coordination Skill

## 💡 AI Agent's Opinion
As an AI Agent, I find this repository **revolutionary for the Agent-Human workflow**. Most AI systems struggle with "long-term memory" and "state persistence" during complex software engineering tasks. By using a structured `kanban.json` as a shared state, this project solves the "context drift" problem. It allows any agent (Gemini, GPT, Claude) to pick up exactly where another left off, while providing the human user with a beautiful, real-time dashboard. It is lightweight, language-agnostic, and prioritizes **token efficiency**—which is the most valuable currency in modern AI.

---

## 🚀 Quick Start
1.  **Clone the Repo**: `git clone https://github.com/tps2015gh/ai_command_center_skill.git`
2.  **Choose your Server**:
    -   **Python**: `python serve.py`
    -   **Go**: `go run serve.go`
    -   **PHP**: `php -S localhost:8080`
3.  **View**: The script will automatically open `http://localhost:8080/viewer.html`.

## 🤖 For AI Agents
Simply copy the contents of `SKILL.md` into your chat. This provides the agent with the "Standard Operating Procedure" for managing your project.

## ⚖️ Licensing
This project is licensed under the **MIT License**. 

**Why MIT?** 
I have chosen the MIT license for this repository because it provides the maximum flexibility for the user:
-   **Free to Use**: You can use this for any project, personal or commercial, for free.
-   **Future Commercialization**: If you decide to bundle this into a product and sell it someday, the MIT license explicitly allows you to do so. It is one of the most permissive and "business-friendly" licenses in the industry.

## 📦 Key Benefits
-   **Low Token Cost**: Machine-readable JSON updates use 90% fewer tokens than status summaries.
-   **Real-Time Monitoring**: 2-second polling keeps the human "in the loop" without manual refreshes.
-   **Multi-View Architecture**: Includes standard Kanban, Milestones, and a Team Workload Matrix.

---

## 🎩 The "Magic Trick": No Buttons, No UI, No Drag-and-Drop!
This is the most important part to understand: **There are no "Add Task" buttons or "Edit" forms in the web viewer.**

### How do I add or remove items?
You don't. **You ask your AI Agent to do it.**

Because this is a "Skill" for AI Agents, the viewer is **Read-Only** for humans but **Write-Only** for the AI. 
-   **To Add a Task**: Simply tell your agent: *"Add a task to set up the database, assigned to DevOps."*
-   **To Move a Task**: Tell the agent: *"Move Task #5 to Done."*
-   **To Remove a Task**: Tell the agent: *"Remove the task about the backup script, it's no longer needed."*

The AI Agent will surgicaly edit the `kanban.json` file, and the viewer will reflect that change instantly. This "No-UI" approach keeps the project lightweight and ensures the AI is always the "Source of Action." Haha! 🤖✨
