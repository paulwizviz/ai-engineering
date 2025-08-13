# AI Agent

An AI agent is a system that perceives its environment, processes that information, and takes actions to achieve specific goals. It’s called an “agent” because it acts on behalf of something or someone, guided by some level of autonomy or intelligence.

## Key Characteristics

1. **Perception** – It gathers data from its environment (e.g. via sensors, input data, APIs).
2. **Reasoning/Decision-Making** – It processes the data using algorithms, models, or rules.
3. **Action** – It takes actions that affect the environment (e.g. moving a robot, answering a question, executing a trade).
4. **Autonomy** – It operates without direct human intervention, at least for parts of its task.

## Types of AI Agents

* **Simple reflex agents** – Act only based on current perceptions.
* **Model-based agents** – Maintain an internal model of the world to inform actions.
* **Goal-based agents** – Aim to achieve a specific outcome.
* **Utility-based agents** – Choose actions that maximise a utility function (a measure of ‘happiness’ or success).
* **Learning agents** – Improve their performance over time using experience/data.

## SDK

An Agent SDK (Software Development Kit) is a toolkit or framework that provides abstractions, components, and utilities to help developers easily build and deploy LLM-powered agents.

It’s not just a wrapper around the LLM — it helps build the whole lifecycle of an agent, including memory, tool use, reasoning, execution, and interaction.

### Characteristics

| Feature | Description |
| --- | --- |
| Prompt Management | Standardised prompts, dynamic prompt templates, role setting. |
| Memory | Short-term (conversational) or long-term (vector DB, file-based) memory. |
| Tool/Plugin Interfaces | Integrations with APIs, web search, file systems, code interpreters. |
| Agent Loop Orchestration | Handles planning-reasoning-action loops like AutoGPT or ReAct. |
| State Management | Tracks steps taken, decisions made, history, etc. |
| Multi-agent support | Enables teams of agents to collaborate or delegate. |
| Execution Environment | Executes actions from LLM, like running code or calling APIs. |

### Python Examples

| SDK | Description |
| --- | --- |
| LangChain | Build LLM chains, agents, tools, memory. Most popular. |
| AutoGen (by Microsoft) | Framework for multi-agent conversation and collaboration. |
| CrewAI | Role-based agent collaboration framework. |
| Semantic Kernel (by Microsoft) | Agentic SDK with skills, memory, planning in .NET and Python. |
| Haystack | Agent-oriented framework for RAG, pipelines, and orchestration. |

## Ai

These are a class of prebuilt but configurable AI-agents to help developers integrate applications with LLM.

### n8n - A workflow automation platform

n8n is an open-source workflow automation platform. Think of it as a digital assembly line where you can connect different applications and services to automate repetitive tasks and build complex data pipelines. It's built on a visual, node-based editor, where each "node" represents a specific action, trigger, or data transformation.

The key differentiators for n8n are:

* **Low-Code/No-Code Flexibility:** While you can build complex workflows with a user-friendly drag-and-drop interface and pre-built nodes, n8n also gives you the option to use custom JavaScript or Python code for more advanced tasks. This hybrid approach is perfect for both technical and non-technical users.

* **Self-Hosting:** This is a huge advantage. N8n is source-available and can be deployed on your own infrastructure (on-premise with Docker or Kubernetes, for example). This gives you full control over your data, a high degree of security, and freedom from vendor lock-in. A managed cloud-hosted version is also available for those who prefer it.

* **AI Integration:** n8n is well-equipped for building AI-powered workflows. It has built-in AI nodes that can connect to various large language models (LLMs) like OpenAI, Gemini, and Claude, allowing you to create sophisticated AI agents for tasks like content generation, data analysis, and customer support.

* **Extensive Integrations:** With over 400 pre-built integrations and a thriving community creating even more, n8n can connect to a wide array of popular apps, services, and databases. If a native node doesn't exist, you can always use the HTTP Request node to connect to any API.

### Local deployments

* [Docker compose](../deployments/n8n/docker-compose.yaml)
