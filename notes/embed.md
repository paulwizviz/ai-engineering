# Embedding and Tokenisation

## Tokenisation and Embedding Overview

**Tokenisation** is how we convert raw text into discrete units (tokens) that computers can process. Think of it like breaking a sentence into LEGO bricks before you can build something with them.
Embeddings come next - they're how we represent those tokens as numbers (specifically, vectors) in a way that captures meaning and relationships.

As a developer, you're used to working with strings. But machine learning models can't work with "hello world" directly - they need numbers.

The simplest approach would be to split on whitespace:

```text
"Hello world" → ["Hello", "world"]
```

But that's naïve. What about punctuation? Rare words? Different languages?

Modern tokenisation (like BPE - Byte Pair Encoding, used in GPT models, or WordPiece, used in BERT) works more cleverly:

- Common words become single tokens: `"running" → ["running"]`
- Rare words get broken into subwords: `"antidisestablishmentarianism" → ["anti", "dis", "establishment", "arian", "ism"]`
- This handles unknown words gracefully and keeps vocabulary sizes manageable

Each token then gets a unique integer ID from the vocabulary.

**Embedding and tokenisation** are fundamental concepts in modern NLP and machine learning, so let's break them down properly.

Now you've got token IDs - integers. But here's the thing: the number 42 isn't inherently "closer" to 41 than to 99 in meaning. We need a better representation.

Embeddings convert discrete tokens into continuous vectors - typically arrays of 100-1000+ floating point numbers.

Think of it like plotting words in multi-dimensional space where similar meanings cluster together. For example:

- "king" and "queen" would be close
- "king" and "bicycle" would be far apart
- The vector for "king" - "man" + "woman" ≈ "queen" (that famous Word2Vec example)

These vectors are learned from data, not hand-crafted. The model discovers that words appearing in similar contexts should have similar embeddings.

## Word2Vec

Let's revisit the king, queen, man and woman example and illustrate embedding with a 2D vector.

We plot words in 2D space:

```text
king:   [0.9, 0.8]   (high on "royalty", high on "male")
queen:  [0.9, 0.2]   (high on "royalty", low on "male" / high on "female")
man:    [0.1, 0.8]   (low on "royalty", high on "male")
woman:  [0.1, 0.2]   (low on "royalty", low on "male" / high on "female")
```

Doing the maths:

```text
king - man + woman = ?

[0.9, 0.8] - [0.1, 0.8] + [0.1, 0.2]
= [0.8, 0.0] + [0.1, 0.2]
= [0.9, 0.2]

≈ queen! [0.9, 0.2]
```

When you subtract man from king, you're removing the "maleness" concept whilst keeping the "royalty" concept.

Then when you add woman, you're adding the "femaleness" concept back in.

So you end up with: royalty + female = queen

The embedding space learned these relationships from text. The model noticed patterns like:

- "king" and "man" appear in similar grammatical contexts (he, his, him)
- "king" and "queen" appear in similar semantic contexts (throne, crown, reign)
- The difference between king/queen is analogous to the difference between man/woman

So the vector arithmetic naturally captures these analogies.

## References

- [Embeddings & Vector Databases Explained](https://www.youtube.com/watch?v=rw1YfQQttfo)
- [A Beginner's Guide to Vector Embeddings](https://www.youtube.com/watch?v=NEreO2zlXDk)