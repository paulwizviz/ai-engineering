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

## Working Examples

* [n8n](https://github.com/paulwizviz/learn-n8n)
