# Large Language Models

Large Language Models are a sophisticated subset of neural networks that operate on fundamentally probabilistic principles. Within the broader hierarchy of artificial intelligence, LLMs represent a specific evolutionary branch: they are deep learning systems built on the transformer architecture, distinguished by their massive scale (billions of parameters) and their specialisation in processing and generating human language. Unlike traditional neural networks designed for classification or pattern recognition tasks with discrete outputs, LLMs are generative models that produce text by continuously predicting probability distributions over possible next tokens. This architectural choice—combining the transformer's attention mechanisms with probabilistic text generation—defines both their remarkable capabilities and their inherent characteristics.

At their core, LLMs operate on a fundamentally probabilistic basis. Unlike traditional software that executes deterministic logic—where the same input always produces the same output—these neural networks generate text by repeatedly predicting probability distributions over thousands of possible next tokens. When you prompt an LLM with "The capital of France is", the model doesn't retrieve a fact from a database; instead, its neural layers calculate that "Paris" has perhaps a 95% probability of being the next token, "Lyon" might have 0.1%, and so forth across its entire vocabulary. This probabilistic mechanism, emerging from millions of weighted connections across dozens of transformer layers, underpins every word the model generates.
This statistical foundation has profound implications for how we understand and deploy these systems. The model's training process—consuming billions of words of text—teaches its neural network to recognise patterns, correlations, and structures in language. Through pure statistical learning via backpropagation and gradient descent, the network's billions of parameters adjust to discover that certain word sequences reliably follow others, that grammatical rules hold with high probability, and that certain reasoning patterns appear frequently in human text. Remarkably, at sufficient scale, these learned statistical patterns encoded in the neural network's weights begin to approximate logical reasoning, mathematical computation, and even creative thinking. The model doesn't "know" facts or "understand" concepts in a human sense; rather, its transformer layers have internalised probabilistic representations of how language behaves when humans discuss those concepts.

The probabilistic nature introduces both capabilities and limitations that practitioners must navigate carefully. On one hand, this approach grants LLMs remarkable flexibility—a single neural network can perform diverse tasks from translation to code generation simply by adjusting the prompt, because its attention mechanisms have learned the statistical patterns underlying all these domains. The sampling process, controlled by parameters like temperature, allows tuning between consistent, deterministic outputs and creative, varied responses. On the other hand, this probabilistic foundation means LLMs will occasionally generate plausible-sounding but entirely fabricated content (hallucinations), exhibit sensitivity to minor prompt variations, and produce non-deterministic results that require careful application design. Understanding that every word emerges from probability calculations flowing through billions of neural connections—not from reasoning or fact retrieval—is essential for building reliable LLM-powered applications and for appreciating both the power and limitations of these extraordinary neural networks.

## Where do LLMs Fit in AI Hireachy

The Hierarchy
Artificial Intelligence (broadest)
└─ Machine Learning (learns from data)
└─ Neural Networks (inspired by brain structure)
└─ Deep Learning (many-layered neural networks)
└─ Transformers (specific architecture using attention)
└─ Large Language Models (transformers trained on text at scale)

## How LLMs Work Under the Hood (The Essential Mental Model)

Think of an LLM as a sophisticated pattern-matching machine that predicts the most likely next piece of text based on patterns it learned from vast amounts of training data.

### The Core Architecture - Transformers

At the heart of modern LLMs is the transformer architecture (introduced in the famous "Attention Is All You Need" paper). Here's the key concept:

### Tokens

Text is broken into chunks called tokens (roughly ¾ of a word). "Hello world" might become ["Hello", " world"].

### Embeddings

Each token is converted into a vector (a list of numbers) that represents its meaning in high-dimensional space. Similar concepts end up close together mathematically.

### Attention Mechanism

This is the breakthrough. The model learns which tokens should "pay attention" to which other tokens. When processing "The cat sat on the mat because it was tired", the model learns that "it" should pay strong attention to "cat" rather than "mat".

### Layers

Modern LLMs have dozens of these attention layers stacked up (Claude Sonnet 4.5, for instance, has many layers). Each layer refines the understanding, building from simple patterns (grammar, syntax) to complex ones (reasoning, context, intent).

### The Forward Pass

When you send a prompt, it flows through all these layers, and the model outputs probabilities for what token should come next. It samples from these probabilities, adds that token to the sequence, and repeats.

### Training

LLMs are trained in stages:

1. **Pre-training:** The model reads massive amounts of internet text, books, code, learning to predict the next token. This is computationally expensive (millions of pounds worth of compute).
2. **Fine-tuning:** The model is refined on high-quality examples of desired behaviour.
3. **RLHF (Reinforcement Learning from Human Feedback):** Human reviewers rank outputs, teaching the model what "helpful, harmless, and honest" looks like.

## The Probabilistic Foundation

At each step, the LLM outputs a probability distribution over all possible next tokens. For example:

"The sky is" → {blue: 45%, clear: 12%, cloudy: 8%, ...}

It then samples from this distribution (with some randomness controlled by "temperature"). So yes, it's probabilistic predictions all the way down—there's no separate "reasoning engine" or symbolic logic system underneath.

However, it is more nuanced than pure randomness.

### Statistical Patterns Can Encode Logic

Through training on vast text, LLMs learn statistical correlations that approximate logical reasoning. When it completes "2 + 2 =" with "4", it's because that pattern appeared millions of times in training, not because it's doing arithmetic.

### Chain-of-Thought Emerges

Remarkably, larger models develop the ability to "reason" through problems when prompted to think step-by-step. This isn't true reasoning—it's learned patterns of how humans explain their thinking—but it produces more reliable outputs.

### Temperature Controls Determinism

- Temperature = 0: Nearly deterministic (always picks highest probability)
- Temperature = 1: More creative/random
- For production apps requiring consistency, use low temperature

### Context Shapes Everything

The probability distribution is entirely dependent on context. "Bank" after "river" vs "money" gets very different predictions.

## Types of Language Models

There are several types of models depending on how you categorise.

- **By architecture:** encoder-only (BERT), decoder-only (GPT), encoder-decoder (T5)
- **By training objective:** masked, autoregressive, prefix LM
- **By capability:** base models, instruction-tuned, chat-optimised

According to Chip Huyen, author of AI Engineering: Building Applications with Foundation Models, language models can be categorised by their training objectives into two main types: **masked language models** and **autoregressive language models**.

A **masked language model** is trained to predict missing tokens anywhere in a sequence, using bidirectional context from both before and after the missing tokens. For example, given "My favourite __ is blue", a masked language model predicts that the blank is likely "colour". BERT (Bidirectional Encoder Representations from Transformers) is the canonical example. These models were historically used for tasks like sentiment analysis, text classification, and named entity recognition that benefit from bidirectional context.

An **autoregressive language model** is trained to predict the next token in a sequence, using only the preceding tokens as context. It predicts what comes next in "My favourite colour is __". Modern LLMs like GPT-4, Claude, and Llama are autoregressive models. While originally designed for text generation, contemporary autoregressive LLMs—through scale and instruction-tuning—now dominate across virtually all language tasks, including classification and analysis, making them the prevailing architecture in current AI applications.

## References

- [How LLMs Actually Generate Text (Every Dev Should Know This)](https://www.youtube.com/watch?v=NKnZYvZA7w4).