# Model Context Protocol

The Model Context Protocol is an open protocol that standardises how applications provide context to Large Language Models (LLMs). It creates a universal way for AI assistants to connect to different data sources and tools through a client-server architecture.

## Terms

* **MCP Hosts**: Applications like Claude Desktop, IDEs, or AI tools that want to access data
* **MCP Clients**: Protocol clients inside the host application
* **MCP Servers**: Lightweight programs that each expose specific capabilities (data sources, tools, prompts)
* **Local Data Sources**: Your computer's files, databases, and services that MCP servers can access

## Core Components

MCP servers can provide three main types of resources:

* **Resources:** Data and content (files, database records, API responses)
* **Tools:** Functions the LLM can invoke to perform actions
* **Prompts:** Templated messages or instructions

## Specification

* [2025-11-25](https://modelcontextprotocol.io/specification/2025-11-25)