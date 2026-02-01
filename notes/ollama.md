# Using Ollama

Ollama is a tool written in Go that simplifies running and managing large language models (LLMs) locally from your terminal.

This guide covers the basic steps for installing Ollama, managing models, and interacting with them.

## Installation and Setup (macOS)

There are two common methods for installing Ollama on macOS.

### Option A: Install via Homebrew (Recommended)

This method installs the Ollama command-line interface (CLI) and sets up a background service.

1. Run the following command in your terminal:

    ```shell
    brew install ollama
    ```

2. The Ollama service will start automatically. You do not need to run `ollama serve` manually.

### Option B: Manual Installation

You can also install the application by downloading it directly.

1. Navigate to the [Ollama website](https://ollama.ai/) and download the macOS application.
2. This will install both the command-line tools and a macOS menu bar application for managing the service.

## Managing Models

Ollama handles the storage and organisation of your models. On macOS and Linux, models are kept in the `~/.ollama/models` directory.

### Finding Models

You can find a library of available models on the [Ollama website](https://ollama.com/library).

### Basic Model Commands

Use the following commands in your terminal to manage your local models.

* **To download a model:**

  ```shell
  ollama pull <model_name>
  ```

  *Example: `ollama pull llama2`*

* **To list all downloaded models:**

  ```shell
  ollama list
  ```

* **To remove a model:**

  ```shell
  ollama rm <model_name>
  ```

  *Example: `ollama rm llama2:latest`*

## Interacting with a Model

Ollama provides two primary ways to interact with an LLM.

### Text Completion Mode

This mode is for generating a single response to a prompt.

* **Command:** `ollama run <model_name> "<your_prompt>"`
* **Use Case:** Best for simple, one-off tasks like summarising text or answering a question.

### Conversational (Chat) Mode

This mode allows for a back-and-forth conversation, where the model remembers the context of previous messages.

* **Command:** `ollama run <model_name>` (without a prompt) or `ollama chat <model_name>`
* **Use Case:** Ideal for chatbot-style interactions or more complex queries that require conversational context.

## Customising Models with a `Modelfile`

A `Modelfile` is a powerful configuration file that allows you to customise how Ollama runs a model. You can use it to set a specific persona, adjust generation parameters, and more.

### Key `Modelfile` Instructions

* `FROM`: Specifies the base model to use (e.g., `FROM llama2`). This is the starting point for your customisation.
* `SYSTEM`: Defines the model's persona or role. This instruction shapes the model's tone and behaviour.
  * *Example: `SYSTEM "You are a helpful and friendly travel agent.""`*
* `PARAMETER`: Adjusts text generation settings.
  * `temperature<value>`: Controls creativity. Higher values (e.g., 0.8) are more random; lower values (e.g., 0.2) are more deterministic.
  * `repeat_penalty<value>`: Helps prevent the model from repeating itself.

### Example `Modelfile`

Here is an example of a simple `Modelfile` that creates a sarcastic movie critic persona.

```shell
# Sarcastic Critic Modelfile
FROM llama2

# Set the persona for the model
SYSTEM """
You are a sarcastic movie critic. Your reviews should be witty, cynical,
and grudgingly give praise only when a film is truly exceptional.
"""

# Set the creativity to be a bit more predictable
PARAMETER temperature 0.4
```

For more detailed information, refer to the [official `Modelfile` documentation](https://docs.ollama.com/modelfile).

### Steps to Customise a Model

STEP 1: Save it as a file (e.g. Modelfile)

STEP 2: `ollama create <choose-a-model-name> -f <location of the file e.g. ./Modelfile>` this will create a model based on the one specified in Modelfile.

STEP 3: `ollama run <choose-a-model-name>`

STEP 4: Start using the model!

## Working Examples

The following examples demonstrates interaction with ollama via curl.

* [Example 1](../examples/ollama/ex1.sh) - This calls ollama expecting no streaming response.
* [Example 2](../examples/ollama/ex2.sh) - This calls ollama expecting streaming response.
