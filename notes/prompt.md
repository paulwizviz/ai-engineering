# Prompt

A prompt is the input text or instruction that you give to a language model (LLM) to guide its output. Examples of prompts are:

* **A question**, e.g., “What’s the capital of France?”
* **A command**, e.g., “Summarise this paragraph.”
* **A role setup**, e.g., “You are a helpful coding assistant…”
* **A template**, e.g., “Given the following context, generate a JSON response…”

## Types of Prompts

| Type | Description |
| --- | --- |
| Zero-shot | The model is asked to do a task without examples. |
| Few-shot | Prompt includes a few examples to guide the model. |
| Chain-of-Thought | Prompt encourages the model to explain reasoning step-by-step. |
| ReAct | Prompt combines reasoning and actions (e.g., tool use). |
| System Prompt | In chat models, sets the assistant's behaviour (e.g., "You are a helpful tutor.").

## Prompt Curator

A prompt curator is someone — or sometimes a system — that designs, organises, and refines prompts to get high-quality, reliable, or specialised outputs from a language model (LLM).

| Task |  Description |
| --- | --- |
| Crafting Prompts | Designs prompt templates for specific goals (e.g., summarisation, Q&A, agent behaviour). |
| Testing Variants | Experiments with different phrasings, formats, roles, or temperature settings. |
| Embedding Instructions | Injects task-specific knowledge, formatting rules, or logic into prompts. |
| Prompt Chaining | Designs sequences of prompts where the output of one becomes the input of another (e.g., ReAct, Chain-of-Thought). |
| Organising a Prompt Library | Maintains and curates a set of prompts for different use cases or users. |
| Evaluating Prompt Quality | Measures prompt performance by output accuracy, tone, speed, or cost. |

## Prompt Curator vs AI Agent

A prompt curator:

* Is not an agent in the traditional AI sense.
* It is typically a human (or sometimes a rule-based system) that designs, selects, or refines prompts (input queries or instructions) to guide an AI model (like ChatGPT) to produce desired outputs.
* Think of it more like a designer or strategist, ensuring that the AI model is asked the right questions in the right way.

An AI agent is:

* An autonomous system that perceives, reasons, and acts to achieve a goal.
* It often operates in an environment over time and can adapt, learn, and make decisions.

| Feature | AI Agent | Prompt Curator |
| --- | --- | --- |
| Role | Autonomous actor | Human (or tool) guiding AI use |
| Intelligence | Has reasoning, planning, and memory | Relies on external AIâ€™s intelligence |
| Autonomy | Acts independently | Usually not autonomous |
| Example | Self-driving car software | Person writing effective prompts |
