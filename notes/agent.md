# Our Thoughts on AI Agents

The idea of "AI Agents" is something we've found ourselves coming back to again and again. It seems to be a key concept in building more advanced AI applications. After a lot of reading and discussion, we've started to think of an AI agent as a system that can act on its own to get things done.

It's called an "agent" because it acts on our behalf, with a certain degree of autonomy. We've boiled its key characteristics down to these four things:

1. **Perception:** It has to be able to sense its environment, whether that's through APIs, input data, or something else.
2. **Reasoning:** It needs to process that information and make a decision.
3. **Action:** It has to be able to do something that affects its environment, like calling an API or answering a question.
4. **Autonomy:** This is the big one. It has to be able to operate without a human holding its hand at every step.

## The Different Kinds of Agents

We've learned that not all agents are created equal. They seem to fall into a few categories, from the simple to the very complex:

* **Simple reflex agents:** These just react to what's happening right now.
* **Model-based agents:** These are a bit smarter; they keep an internal "map" of the world to help them decide what to do.
* **Goal-based agents:** These agents are focused on achieving a specific outcome.
* **Utility-based agents:** These are more advanced, trying to choose the action that will give them the most "utility" or bring them closest to success.
* **Learning agents:** These are the most exciting to us. They can actually improve their own performance over time based on experience.

## How We Can Build Agents: SDKs

We quickly realised that building an agent from scratch is a huge task. There are Software Development Kits (SDKs) out there that give us a head start.

We see these SDKs as more than just a simple wrapper around an LLM. They help us build the entire agent, including its memory, its ability to use tools, and its reasoning loop.

Here are the key features we've been looking for in an Agent SDK:

| Feature | What it means to us |
| --- | --- |
| **Prompt Management** | Helps us create and manage good, consistent prompts. |
| **Memory** | Gives the agent a way to remember things, both short-term and long-term. |
| **Tool/Plugin Interfaces** | Lets the agent connect to the outside world, like searching the web or reading files. |
| **Agent Loop Orchestration** | Manages the core "think, decide, act" loop that the agent runs on. |
| **State Management** | Keeps track of what the agent has done and what it has learned. |
| **Multi-agent support** | Some frameworks even let us create teams of agents that can work together. |
| **Execution Environment** | Provides a safe way for the agent to actually perform actions, like running code. |

### Python SDKs We've Looked At

The Python ecosystem is full of these, but here are a few of the main ones we've been exploring:

| SDK | Our Quick Take |
| --- | --- |
| **LangChain** | The most popular one. It's a huge library for chaining LLM calls and building agents. |
| **AutoGen** | A framework from Microsoft that's focused on getting multiple agents to collaborate. |
| **CrewAI** | A newer one we're watching that focuses on role-based agent collaboration. |
| **Semantic Kernel** | Another one from Microsoft, which lets you define agent "skills". |
| **Haystack** | A great framework for building pipelines, especially for RAG (Retrieval-Augmented Generation). |

## Our Working Examples

We're always trying to build things to learn better. Here's an example project where we've been playing with some of these concepts:

* [n8n Workflow Automation](https://github.com/paulwizviz/learn-n8n)