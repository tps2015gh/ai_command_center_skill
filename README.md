# How to Use the AI Command Center

## For Gemini CLI (Local Installation)
1.  Copy the `ai_command_center_skill` folder to your project root.
2.  The AI Agent will automatically read `SKILL.md` and know how to update the `kanban.json`.
3.  To view the board:
    -   Run: `php -S localhost:8080` inside the folder.
    -   Open: `http://localhost:8080/viewer.html`.

## For Other AI Agents (ChatGPT, Claude, etc.)
1.  **Copy-Paste** the content of `SKILL.md` into the agent's System Prompt or the beginning of your chat session.
2.  **Provide context**: Tell the agent "You have access to a local file `kanban.json`. Use the schema defined in the skill instructions to track our progress."
3.  The agent will then generate the correct JSON blocks for you to save to your local file.

## Why use this?
-   **Low Token Cost**: The JSON format is extremely compact compared to long text updates.
-   **Structured History**: Keeps a permanent, machine-readable log of who did what and when.
-   **Visual Clarity**: Provides the human user with a professional dashboard without requiring the AI to "explain" status in every message.
