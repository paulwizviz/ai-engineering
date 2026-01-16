# Language Model

We learn from reviewing the book "AI Engineering: Building Applications with Foundation Model" that a language model encodes statistical information about text base data.

We also learnt that a language model is one type of AI model and that it is optimised for a particular use case, namely processing languages or text-based dataset. The following tables summarise our understanding.

| Feature | LLMs | Other Models |
| ----- | ---- | ---- |
| Input type | Text | Images, audio, numbers, or narrow data |
| Capabilities | General-purpose reasoning + language | Usually single-task |
| Training data | Huge text corpora | Task-specific datasets |
| Flexibility | Very flexible | Limited, highly specialised |
| Examples | ChatGPT, Claude | Image classifiers, speech recognizers, recommender systems, Machine Learning (Linear regression, etc) |

## Token

The basic unit of a language model is token. A token can be a character, a word, or a part of a word, depending on the model. We learnt that a phrase like: `I can't walk anymore`. The word `can't` is broken into two tokens: `can` and `t`. The phrase is composed of 5 tokens.

The process of breaking words into token is tokenization. A set of token is known is known as a vocabulary. Different models has different vocabulary sizes. A model's vocabulary size determines the number of distinct words, it can construct or "understand". The tokenization method and vocabulary size are decided by model developers.

## Types of Languge Models

According to Chip Huyen author of AI Engineering: Building Applications with Foundation Models, there are two main types of language models: `masked language models` and `autoregressive language models`.

A **masked language model** is trained to predict missing tokens anywhere in a sequence, using the context from both before and after the missing tokens. For example, given the context, “My favorite __ is blue”, a masked language model should predict that the blank is likely “color”. Masked language models are commonly used for non-generative tasks such as sentiment analysis and text classification.

An **autoregressive language model** is trained to predict the next token in a sequence, using only the preceding tokens. It predicts what comes next in “My favorite color is __.” An **autoregressive language models** are the models of choice for text generation.