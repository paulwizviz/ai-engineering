# AI Agents

An AI agent is a software system that combines Large Language Models (LLMs) with external tools to autonomously complete tasks. Unlike standalone LLMs that only generate text, agents can:

- Interpret natural language requests
- Decide which actions to take from a predefined set
- Execute actions using deterministic tools (APIs, databases, etc.)
- Coordinate multi-step workflows

**Key distinction:** The LLM acts as an intelligent router that classifies user intent and selects appropriate actions, whilst deterministic code ensures reliable execution.

## Key Characteristics and Capabilities

- **Autonomy:** Agents can operate independently to complete multi-step workflows after being given a goal.
- **Tool Use:** They can interact with external software, APIs, and databases to perform tasks, such as updating a CRM or booking a flight.
- **Reasoning & Planning:** Agents use LLMs to decompose complex tasks into smaller steps by probabilistically selecting from predefined actions. The LLM classifies user intent and determines the next action based on learnt patterns, not formal logic.
- **Memory:** Agents maintain conversation state and context within a session, allowing them to track multi-turn interactions. This memory is typically temporary (session-based) and does not persist across different conversations.
- **Learning:** Most production agents do NOT learn or adapt during operation. Their behaviour is determined by:
  - The underlying LLM (which is fixed after training)
  - System prompts and tool definitions (which are static)
  - Conversation context (which resets each session)

**Note:** Some advanced systems may log interactions for offline analysis and prompt refinement, but this is not real-time learning.

## Components of an AI Agent

### LLM as Decision Engine

Agents use LLMs to:

1. **Classify user intent** - Map natural language to predefined action categories
2. **Extract parameters** - Identify relevant values (dates, locations, etc.)
3. **Select actions** - Choose which tool to invoke from a finite set
4. **Generate responses** - Communicate results in natural language

The LLM acts as a probabilistic classifier/router, mapping unstructured user input to structured tool calls. The tools themselves execute deterministically.

**Critical distinction:** The LLM doesn't "reason" or "decide" in a logical sense—it performs sophisticated pattern matching to select the most probable action given the context.

### Perception

Gathering data from user prompts or external environments.

### Action/Tools

Using software capabilities to execute tasks.

## Types of Modern AI Agents

### By Architecture

**1. ReAct Agents (Reasoning + Acting)**

- LLM generates: Thought → Action → Observation loop
- Example: "Thought: I need origin. Action: ask_origin. Observation: User said SFO"
- Most common pattern for tool-using agents

**2. Tool-Calling Agents**

- LLM directly invokes tools via structured API (e.g., MCP, function calling)
- No explicit "reasoning" steps shown to user
- Example: Flight booking agent

**3. Planning Agents**

- LLM creates complete plan upfront, then executes
- Example: "Step 1: Get origin. Step 2: Get destination. Step 3: Search. Step 4: Book"
- Less flexible but more predictable

**4. Conversational Agents**

- Primarily dialogue-driven, minimal tool use
- Focus on natural interaction over task completion
- Example: Customer service chatbot

### By Autonomy Level

**1. Single-Turn Agents**

- Execute one action per user request
- Example: "Search for flights" → returns results, done

**2. Multi-Turn Agents**

- Complete multi-step workflows within one conversation
- Example: Full booking flow (gather info → search → confirm → book)

**3. Fully Autonomous Agents**

- Operate with minimal human intervention
- Example: Background agent that monitors prices and alerts user
- Rare in production due to reliability concerns

### By Memory Scope

**1. Stateless Agents**

- No memory between conversations
- Each interaction is independent

**2. Session-Based Agents**

- Remember context within a single conversation
- State resets when conversation ends
- Most common for production systems

**3. Persistent Memory Agents**

- Maintain user preferences and history across sessions
- Example: "Book my usual flight" (remembers past bookings)
- Requires external storage (database, vector DB)

## Agent vs. LLM vs. Traditional Software

| Aspect | Traditional Software | Standalone LLM | AI Agent |
| ------ | -------------------- | -------------- | -------- |
| **Input** | Structured (forms, API calls) | Natural language | Natural language |
| **Processing** | Deterministic logic | Text generation | LLM + deterministic tools |
| **Output** | Structured data | Text responses | Actions + text |
| **Flexibility** | Requires code changes | Handles variation | Handles variation + takes action |
| **Reliability** | 100% predictable | Probabilistic | Hybrid (LLM routes, tools execute) |
| **Example** | Booking form with dropdowns | ChatGPT answering questions | Flight booking agent that searches AND books |

## Critical Design Principle: Bounded Action Space

**Key Insight:** Agents work reliably because they operate within a **finite, predefined set of actions**, not infinite possibilities.

**The Pattern:**

```text
Infinite user inputs → LLM (classifier) → Finite actions → Deterministic execution
```

**Example:**

- Possible user inputs: Unlimited ways to say "I want to fly to Tokyo"
- Agent actions: 8 predefined operations (ask_origin, search_flights, etc.)
- Tool execution: Deterministic code with validation

**Why this matters:** The agent's LLM doesn't "reason" freely. It classifies input to select from your predefined menu of safe, tested operations.

## Decision Framework: Agent vs. Traditional Software

**Use AI Agents when:**

- Natural language input is required
- User intent varies widely but maps to finite operations
- Multi-step coordination with judgement calls
- Requirements change frequently (easier to update prompts than code)

**Use Traditional Software when:**

- Performance critical (sub-second response times)
- Absolute determinism required (financial transactions, medical dosing)
- High-volume, low-complexity operations (form submissions)
- No ambiguity in inputs/outputs

**Best: Hybrid Architecture**

- Agents for: Intent understanding, workflow orchestration, user communication
- Traditional for: Database queries, calculations, transaction processing, validation

Example: Flight booking agent uses LLM to understand "I want to fly to Tokyo" but PostgreSQL to search flights and process payments.

## AI Agent vs. Agentic AI

**AI Agent** (noun): A specific software system that uses an LLM to interpret requests, select actions from a predefined set, and execute multi-step workflows.

- Example: Flight booking agent, customer service agent
- Concrete, deployable application

**Agentic AI** (adjective): A design quality describing systems with agent-like characteristics such as autonomous decision-making, goal-directed behaviour, and adaptive responses.

- Example: "Our system has agentic capabilities"
- Describes HOW a system works, not WHAT it is

**Key distinction:**

- AI Agent = the thing you build
- Agentic AI = the quality it has

**Analogy:**

- "Self-driving car" (AI Agent) vs. "autonomous driving capability" (Agentic AI)

## Implementation Framework and Examples

### n8n

`n8n` is a worflow automation platform for building and running agentic system. How it works:

- **Orchestration Layer:** Manages the flow between different AI models and external tools, ensuring steps happen in the right order.
- **Connects to Everything:** Integrates with hundreds of apps, databases, and APIs, giving agents access to real-world data and actions.
- **Logic & Control:** Handles conditional logic, retries, and human approvals, making AI less erratic and more reliable.
- **Visual Development:** Offers a node-based interface (low-code/no-code) to design complex agentic processes without deep programming.
- **Production Ready:** Provides features like error handling and auditing, crucial for deploying AI agents in business environments.

There is a local docker based `n8n` deployments under `./deployments` consisting of these:

- **[docker-compose.yaml](../deployments/docker/docker-compose.yaml)**: This file defines our n8n environment. We're using the official `n8nio/n8n` image and have set it up to persist all our workflow data into a local `n8n-data` directory. This is crucial so we don't lose our work when the container restarts. It also sources a number of environment variables from the `ops.sh` script to configure things like the host, port, and basic authentication.
- **[ops.sh](../deployments/docker/ops.sh)**: This is a simple shell script we've put together to handle the common operations. It allows us to `start` and `stop` the n8n container. There's also a `clean` command, but a word of caution: **it will permanently delete all your n8n data**, so use it with care! It's really only there for when we need a completely fresh start.

## References

- [What are AI agents by IBM](https://www.ibm.com/think/topics/ai-agents)
- [RAG Explained For Beginners](https://www.youtube.com/watch?v=_HQ2H_0Ayy0)
