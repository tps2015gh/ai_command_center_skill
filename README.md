# AI Command Center: Professional Agent Coordination Skill

## 💡 The AI Agent's Perspective
This repository is **revolutionary for the Agent-Human workflow**. It solves the "context drift" problem by using a structured `kanban.json` as a shared project state. It allows any agent (Gemini, GPT, Claude) to maintain continuity, while providing you with a beautiful, real-time command center.

---

## 🚀 Quick Links
-   [**Quick Start Guide** (Setup in 60s)](quickstart.md)
-   [**Standard Operating Procedure** (Agent Instructions)](SKILL.md)

---

## 🎩 The "Magic Trick": No Buttons, No UI, No Drag-and-Drop!
This is the core philosophy: **The web viewer is read-only for humans.**

### How do I manage work?
**You ask your AI Agent to do it.**
-   **To Add**: *"Add task 'Fix CSS' for FE."*
-   **To Move**: *"Move Task #3 to Review."*
-   **To Remove**: *"Remove Task #5."*

The AI Agent surgically edits the `kanban.json` file, and the viewer reflects that change instantly. This ensures the AI remains the **Source of Action** while you remain the **Source of Intent**.

> **⚠️ Important Note**: This repository is designed to "wait" for a connection from **Gemini** or your preferred **AI Agent**. You do not need to fill out any forms or enter data manually. Gemini or the Agent will autonomously fill the project state, create tasks, and update milestones by themselves the moment you make a request in the prompt.

---

## 🔮 Why No-UI is the Future
1.  **Intent is the Interface**: We are moving from "point and click" to "state your goal."
2.  **Precise Addressability**: Every component (Task, Milestone) has a unique **ID**. This allows agents to target data with 100% precision, eliminating the ambiguity of traditional UIs.
3.  **Token Efficiency**: Machine-readable JSON updates use 90% fewer tokens than long status descriptions.

---

## 👥 The Default Elite Team
The skill initializes with a high-performance team using compact **Short IDs**:
-   **TL**: Tech Lead (Arch)
-   **SA**: System Analyst (Logic)
-   **FE**: Frontend (UI/UX)
-   **BE**: Backend (API/DB)
-   **SEC**: Security (Auth)
-   **OPS**: DevOps (Git/CI)
-   **FIX**: Elite Debugger (Deep investigation/Complex fixes)

---

## ⚖️ Licensing
Licensed under the **MIT License**. Use it for free, bundle it into products, or sell it—you have full commercial freedom.

---

## 📦 Multi-Platform Support
One-click serving scripts included for:
-   **Node.js** (`serve.js`)
-   **Python** (`serve.py`)
-   **Go** (`serve.go`)
-   **PHP** (`php -S`)
