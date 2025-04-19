# SaGaCLI

SaGaCLI is a simple yet powerful command-line interface for leveraging Large Language Models (LLMs). It can use OpenAI's GPT-3.5/GPT-4, Anthropic's Claude models, or Google's Gemini models to provide translation, summarization, explanation, and search capabilities for text.

## Features

- **Multiple LLM Support**: Use OpenAI (GPT-3.5/GPT-4), Claude (Anthropic), or Gemini (Google) models
- **Diverse Functions**: Translation, summarization, explanation, search, custom messages, and more
- **Flexible Input**: Accept text from file paths or standard input
- **Multilingual Support**: Specify any target language for output

## Installation

### Prerequisites

- Go 1.16 or higher

### Method 1: Binary Installation (Recommended)

```bash
# Install Go if not installed
$ brew install go  # macOS
# or
$ sudo apt install golang-go  # Ubuntu/Debian

# Install SaGaCLI
$ go install github.com/sa-giga/saga-cli@latest

# Add SaGaCLI to your path
$ export PATH=$PATH:$(go env GOPATH)/bin
```

### Method 2: Building from Source

```bash
# Clone the repository
$ git clone https://github.com/sa-giga/saga-cli.git
$ cd saga-cli

# Build and install
$ make build  # Creates binary at bin/saga
$ make install  # Installs to your system
```

## Detailed Configuration

SaGaCLI can use OpenAI's GPT-3.5/GPT-4, Anthropic's Claude models, or Google's Gemini models. Set the appropriate environment variables based on the service you want to use.

### Using OpenAI API

```bash
# Required settings
export OPENAI_API_KEY=sk-...  # Your OpenAI API key
export OPENAI_API_TYPE=openai  # API type (default)

# Optional settings
export OPENAI_API_BASE_URL=https://api.openai.com/v1  # Base URL (default)
export OPENAI_API_VERSION=2023-05-15  # API version
export OPENAI_API_MODEL=gpt-3.5-turbo  # Model to use (default: gpt-3.5-turbo)

# For Azure OpenAI Service
export OPENAI_API_TYPE=azure
export OPENAI_API_BASE_URL=https://your-resource-name.openai.azure.com
export OPENAI_API_KEY=your-azure-api-key
export OPENAI_API_VERSION=2023-05-15
export OPENAI_API_MODEL=your-deployment-name
```

### Using Anthropic Claude API

```bash
# Required settings
export CLAUDE_API_KEY=sk-ant-...  # Your Anthropic API key
export OPENAI_API_TYPE=claude  # Set API type to claude

# Optional settings
export CLAUDE_API_MODEL=claude-3-haiku-20240307  # Model to use (default)
# Available models: claude-3-opus-20240229, claude-3-sonnet-20240229, claude-3-haiku-20240307, etc.
```

### Using Google Gemini API

```bash
# Required settings
export GEMINI_API_KEY=your-gemini-api-key  # Your Gemini API key
export OPENAI_API_TYPE=gemini  # Set API type to gemini

# Optional settings
export GEMINI_API_MODEL=gemini-1.5-pro  # Model to use (default)
# Also available: gemini-1.5-flash, etc.
```

You can check your current settings by running the `saga env` command.

## Detailed Usage

### Basic Usage

```bash
# Process file content
$ saga [options] < input_file

# Or
$ cat input_file | saga [options]

# Direct text input from standard input (end with Ctrl+D)
$ saga [options]
Enter your text...
[Ctrl+D]
```

### Main Options

```bash
--translation  # Translation mode
--summary      # Summarization mode
--explanation  # Explanation mode
--search       # Search mode
--message      # Custom message mode (send specific instructions)
--lang [language_code]  # Specify output language (e.g., en, ja, fr, zh, etc.)
```

### Usage Examples

```bash
# Translate text to Japanese
$ echo "Hello world" | saga --translation --lang ja

# Summarize a long document in English
$ cat document.txt | saga --summary --lang en

# Explain program code
$ cat code.py | saga --explanation --lang en

# Search for information on a specific topic
$ echo "Quantum computing basics" | saga --search --lang en

# Extract name field from JSON data
$ cat document.json | saga --message "Extract the name value" --lang en

# Calculate average of age column in CSV data
$ cat document.csv | saga --message "Calculate the average of the age column" --lang en

# Translate a document to Japanese then summarize it
$ cat english_doc.txt | saga --translation --lang ja | saga --summary --lang ja
```

### View Examples

```bash
# View various usage examples
$ saga examples
```

## Troubleshooting

### API Key Not Set Error

```
Error: OPENAI_API_KEY or CLAUDE_API_KEY or GEMINI_API_KEY is not set
```

If you see this error, check that your environment variables are correctly set.

### Model Selection

If processing is slow or you need higher quality results, select a more powerful model via environment variables:

```bash
# For OpenAI
export OPENAI_API_MODEL=gpt-4

# For Claude
export CLAUDE_API_MODEL=claude-3-opus-20240229

# For Gemini
export GEMINI_API_MODEL=gemini-1.5-pro
```

## Getting Help

```bash
# Display command help
$ saga --help

# Display version information
$ saga --version

# Check environment settings
$ saga env
```

## License

MIT License
