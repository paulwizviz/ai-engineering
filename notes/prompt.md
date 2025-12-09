# Our Thoughts on Prompts

Prompt seems to be one of the most important pieces of the puzzle when working with Large Language Models (LLMs). We've come to think of a prompt as simply the instruction we give to a model to get it to do something for us.

It can be as simple as:

* A question, like “What’s the capital of France?”
* A command, like “Summarise this paragraph for me.”
* Setting up a role, for example, “You are a helpful coding assistant…”
* Or even a template we want it to fill, like “Given the following context, generate a JSON response…”

## The Different Kinds of Prompts We've Used

We learnt that prompts come in a few different flavours, depending on what we're trying to achieve. Here are the main types we've been experimenting with:

* **Zero-shot:** This is when we just ask the model to do something without giving it any examples.
* **Few-shot:** Here, we give the model a few examples of what we want, right in the prompt, to help guide it.
* **Chain-of-Thought:** We've found this one fascinating. We encourage the model to 'think out loud' and explain its reasoning step-by-step, which can often lead to better answers for complex problems.
* **ReAct:** This is a more advanced technique we're exploring, where the prompt gets the model to combine its reasoning with actions, like using a tool.
* **System Prompt:** In the chat models we've used, this is a powerful way to set the assistant's personality or ground rules from the start (e.g., "You are a helpful tutor who always explains things simply.").

## The Idea of a 'Prompt Curator'

As we got deeper into this, we started talking about the role of a 'prompt curator'. It's not a formal job title we've seen much, but more a way of describing the work that goes into designing, testing, and refining prompts to get consistently good results from an LLM.

It feels like this role involves a few key tasks:

* **Crafting Prompts:** Designing the basic templates for what we want to achieve.
* **Testing Variants:** Just trying out different ways of asking for the same thing to see what works best.
* **Embedding Instructions:** Putting specific rules or knowledge directly into the prompt itself.
* **Prompt Chaining:** This is a cool one, where we design a sequence of prompts, and the output of one becomes the input for the next.
* **Organising a Prompt Library:** We're starting to think we need one of these—a shared place to keep all our best prompts.
* **Evaluating Prompt Quality:** Figuring out how to measure if a prompt is 'good' based on the accuracy, tone, or even cost of the output.

## Prompt Curator vs. an AI Agent

We had a bit of a debate about this, trying to get the distinction clear in our own heads. Here's where we've landed:

A **prompt curator**, as we see it, is the person (or maybe a system) who designs the questions. They're like a strategist, figuring out the best way to ask the AI model for something. They aren't typically autonomous.

An **AI agent**, on the other hand, is the autonomous system itself. It's the thing that perceives its environment, makes decisions, and acts on its own to achieve a goal.

We jotted down a quick comparison to keep it straight:

| Feature | AI Agent | Prompt Curator |
| --- | --- | --- |
| **Role** | An autonomous actor | The human (or tool) guiding the AI |
| **Intelligence** | Has its own reasoning & planning | Relies on the AI's intelligence |
| **Autonomy** | Acts independently | Not usually autonomous |
| **Example** | The software in a self-driving car | One of us writing a good prompt! |