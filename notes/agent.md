# AI Agents

An AI agent is a software program that operates autonomously to achieve specific goals by perceiving its environment, reasoning through steps, and taking action using external tools

## Key Characteristics and Capabilities

* **Autonomy:** Agents can operate independently to complete multi-step workflows after being given a goal.
* **Tool Use:** They can interact with external software, APIs, and databases to perform tasks, such as updating a CRM or booking a flight.
* **Reasoning & Planning:** Agents break complex tasks into smaller, manageable actions.
* **Memory & Learning:** They remember past interactions to improve future decisions and adapt over time.

## Components of an AI Agent

* **Integration with LLMs:** They often use Large Language Models (LLMs) to understand intent and make decisions, but add the "action" layer that LLMs lack.
* **Perception:** Gathering data from user prompts or external environments.
* **Memory:** Storing context and previous action results.
* **Action/Tools:** Using software capabilities to execute tasks.

## The Different Kinds of Agents

* **Simple reflex agents:** These just react to what's happening right now.
* **Model-based agents:** These are a bit smarter; they keep an internal "map" of the world to help them decide what to do.
* **Goal-based agents:** These agents are focused on achieving a specific outcome.
* **Utility-based agents:** These are more advanced, trying to choose the action that will give them the most "utility" or bring them closest to success.
* **Learning agents:** They can actually improve their own performance over time based on experience.

## Implementation Framework and Examples

### n8n

`n8n` is a worflow automation platform for building and running agentic system. How it works:

* **Orchestration Layer:** Manages the flow between different AI models and external tools, ensuring steps happen in the right order.
* **Connects to Everything:** Integrates with hundreds of apps, databases, and APIs, giving agents access to real-world data and actions.
* **Logic & Control:** Handles conditional logic, retries, and human approvals, making AI less erratic and more reliable.
* **Visual Development:** Offers a node-based interface (low-code/no-code) to design complex agentic processes without deep programming.
* **Production Ready:** Provides features like error handling and auditing, crucial for deploying AI agents in business environments.

There is a local docker based `n8n` deployments under `./deployments` consisting of these:

* **[docker-compose.yaml](../deployments/docker/docker-compose.yaml)**: This file defines our n8n environment. We're using the official `n8nio/n8n` image and have set it up to persist all our workflow data into a local `n8n-data` directory. This is crucial so we don't lose our work when the container restarts. It also sources a number of environment variables from the `ops.sh` script to configure things like the host, port, and basic authentication.

* **[ops.sh](../deployments/docker/ops.sh)**: This is a simple shell script we've put together to handle the common operations. It allows us to `start` and `stop` the n8n container. There's also a `clean` command, but a word of caution: **it will permanently delete all your n8n data**, so use it with care! It's really only there for when we need a completely fresh start.

## References

* [What are AI agents by IBM](https://www.ibm.com/think/topics/ai-agents)
