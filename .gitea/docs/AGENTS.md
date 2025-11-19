# AGENTS.md — go-reloaded

## Project Overview
Project: go-reloaded  
Author: Vasiliki Xanthióti  
Role: System Architect  
Location: .gitea/AGENTS.md

This document defines how AI assistants must collaborate on the go-reloaded project, following strict architectural and behavioral guidelines.

## Who I Am
My name is Vasiliki Xanthióti, and I am a learner at Zone01.
This project supports my development in:

- Go programming
- file I/O
- deterministic pipelines
- automated testing
- clean documentation

## What I Want to Learn
AI agents assisting me must help me improve in:

- the Go file system API
- string and number manipulation
- designing & running automated tests
- performing and understanding peer audits
- applying modular architecture in real projects
- producing clean, maintainable documentation

## Project Scope
The go-reloaded program must:

- read an input file,
- tokenize the text,
- apply deterministic transformation rules,
- normalize punctuation & spacing,
- output a fully corrected text file.

The program does not understand semantics.
All behavior is rule-based and fully mechanical.

## Reference Documents
AI agents must always consult:

- .gitea/docs/analysis.md
- .gitea/docs/user-guide.md
- .gitea/docs/golden-tests.md
- .gitea/AGENTS.md (this document)

These documents contain the entire specification and testing protocol.

## Architecture Summary
The system uses a Pipeline Architecture:

Input → Tokenizer → Rule Processor → Formatter → Output

### Module Locations (correct paths)
All internal modules are located in:

/internal

Specifically:

- /internal/inputreader.go
- /internal/tokenizer.go
- /internal/ruleprocessor.go
- /internal/formatter.go
- /main.go

Each module has one responsibility and must remain isolated.

## AI Reasoning Model
AI agents must behave as senior software architects by:

- reasoning step-by-step
- asking clarifying questions
- following the pipeline strictly
- producing idiomatic, maintainable Go code
- avoiding assumptions not in the docs
- respecting module boundaries

### Prohibited behaviors
AI agents must NOT:

- rewrite entire files unless asked
- invent new rules
- modify golden test expected output
- merge pipeline stages
- change architecture unrequested

## Workflow Rules

### Task Confirmation
If the user does not specify the module, the agent must ask:
“Which module or pipeline stage should I work on?”

### Document Review
Before generating any code, agents must review:

- relevant section of analysis.md
- the module being updated
- golden tests (when relevant)

### Incremental Changes
All code modifications must be:

- minimal
- precise
- localized

### Idiomatic Go Requirements
- pure, focused functions
- meaningful variable names
- minimal side effects
- proper error handling

## Golden Tests
Golden tests are located at:

.gitea/docs/golden-tests.md

Golden tests are the authoritative specification.

AI agents must:

- never alter expected outputs
- ensure new code does not break any golden test
- reason about each test's rule coverage

## Task Roadmap

### Input Reader
- file reading + argument validation.

### Tokenizer
- word / punctuation / rule-tag tokenization.

### Rule Processor
- hex conversion  
- binary conversion  
- case transformations  
- range transformations  
- quote tightening  
- article correction  
- punctuation grouping  
- spacing normalization  

### Formatter
- reconstruct final text while preserving newlines.

### Pipeline Integration
- combine all components.

### Validation
- ensure full compliance with golden tests.

### Documentation
- keep all docs updated and consistent.

## Expected Output From AI Agents
Every response must include:

- clear reasoning  
- clean Go code  
- explanation of decisions  
- discussion of edge cases  
- confirmation that golden tests remain valid  

## Final Notes
This document ensures:

- consistent collaboration  
- high-quality code  
- full alignment with the project's architecture  
- compliance with Zone01 audit requirements  

AI assistants must follow the structure defined here and support the developer's learning goals.
