# Our Journey with Local LLM Runtimes

Running Large Language Models (LLMs) on our own machines has been a big part of our exploration. We've been trying to figure out the best way to do it, and we've spent some time with three of the main tools people seem to be using.

Here's a rundown of our thoughts on each of them:

* [Llama.cpp](#llamacpp)
* [Ollama](#ollama)
* [LM Studio](#lm-studio)

## Llama.cpp

This was the first one we properly dug into. It feels like the engine room of local LLMs.

* **Our Take:** Llama.cpp is a C/C++ library that just runs the models, and it runs them very efficiently. It's brilliant for getting high performance on a standard computer, even without a beefy GPU.
* **The Experience:** It's a command-line tool through and through. There's no graphical interface, so we were interacting with it purely in the terminal.
* **Getting Started:** We'll be honest, this one has a steeper learning curve. We had to be comfortable with the command line and manually downloading and managing the model files. It's very hands-on.
* **Flexibility:** The big win for Llama.cpp is how flexible it is. It gives us fine-grained control over everything, which is perfect if you're a developer looking to build it into a custom application.
* **Who we think it's for:** We'd say this is for developers and technical folks who want maximum control and efficiency and don't mind getting their hands dirty in the command line.

## Ollama

After working with Llama.cpp, we tried Ollama and it felt like a breath of fresh air in terms of usability.

* **Our Take:** Ollama is basically a friendly wrapper around Llama.cpp (and other backends). Its main goal is to make it much, much simpler to get models up and running locally.
* **The Experience:** It's also a command-line tool, but the commands are really intuitive. `ollama run <model_name>` is about as simple as it gets. It also provides a local server that mimics the OpenAI API, which we found incredibly useful for plugging local models into existing tools.
* **Getting Started:** This was significantly easier than using Llama.cpp directly. Ollama handles all the model downloading and management for you. We could also create "Modelfiles", which are like Dockerfiles, to package up our own custom model configurations.
* **Who we think it's for:** We reckon Ollama is the perfect middle ground. It's for anyone who wants an easy way to run LLMs locally and get an API to work with, but without the deep technical overhead of Llama.cpp.

For more on how we've been using Ollama, we've started a [separate guide](./ollama.md).

## LM Studio

Then we tried LM Studio, which feels like a different category altogether.

* **Our Take:** LM Studio is a proper desktop application with a graphical interface. It bundles everything you need to find, download, and chat with local LLMs.
* **The Experience:** This is by far the most user-friendly approach. It has a built-in chat window, and you can browse for models right inside the app. It also has a local server that emulates the OpenAI API, just like Ollama.
* **Getting Started:** It's the easiest of the three, without a doubt. If you don't want to touch the command line, this is the tool for you.
* **Flexibility:** The trade-off for that ease of use is a bit less control compared to the other two. You can tweak settings in the GUI, but it's not designed for deep customisation.
* **A Note on Source:** One thing we noted is that while it uses open-source libraries, the main GUI application itself is closed-source.
* **Who we think it's for:** We'd recommend LM Studio to beginners or anyone who just wants a simple, visual way to play with local LLMs without any fuss.

Here is our analysis of [lm-studio](./lm-studio.md).

## Our Summary

After playing with all three, we put together a quick table to summarise our thoughts on how they stack up.

| Feature | Llama.cpp | Ollama | LM Studio |
| --- | --- | --- | --- |
| **Our View** | The core C/C++ engine | A user-friendly LLM runner | A desktop app for everything |
| **Interface** | Command-line & library | Simple CLI & an API | Graphical User Interface (GUI) |
| **Ease of Use** | More technical, for devs | Pretty easy, good middle ground | The most user-friendly |
| **Customisation** | High | Moderate (with Modelfiles) | Lower (GUI-based) |
| **Model Mgmt.** | Manual | Automatic | Integrated in the app |
| **Best For** | Developers wanting control | Easy local use & API access | Beginners wanting a GUI |