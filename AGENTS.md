AGENTS.md — go-reloaded

(AI Collaboration & Architecture Guidelines)

🧭 Purpose

This document gives AI coding agents clear instructions for working on the go-reloaded project.

Humans read README.md

AI reads AGENTS.md

This file ensures that every AI assistant:

understands the architecture,

respects your pipeline,

writes consistent, idiomatic Go code,

and follows the project rules exactly.

🧩 Project Overview

Project Name: go-reloaded
Language: Go (Golang)
Architecture: Modular Pipeline
Goal: Transform an input text file according to strict rules (hex/bin conversion, case editing, punctuation, quotes, a→an).

Reference documents:

docs/analysis.md — full rule specification

docs/user-guide.md — project problem description

docs/golden-tests.md — golden test set

AGENTS.md — this file

🧱 Architecture Summary

The project uses a Pipeline Architecture — each stage is independent and takes the output of the previous one.

Input → Tokenize → Rule Processing → Formatting → Output

📦 Modules
Stage	Description	File
Input Reader	Reads input file, returns content	/input_reader/input_reader.go
Tokenizer	Splits text into tokens (words, punctuation, rule tags)	/tokenizer/tokenizer.go
Rule Processor	Applies transformations (hex, bin, up, low, cap, quotes, punctuation, article)	/rule_processor/rule_processor.go
Formatter	Joins processed tokens back into final string	/formatter/formatter.go
CLI	Argument parsing and pipeline orchestration	main.go
⚙️ Transformation Rules (Summary)

(hex) → convert previous word from hex to decimal

(bin) → convert previous word from binary to decimal

(up) / (low) / (cap) → change previous word’s case

(up, n) / (low, n) / (cap, n) → apply to previous n words

punctuation hugs previous word

punctuation groups (..., !?, ?!) stay together

quotes tighten: ' hello ' → 'hello'

a → an before vowel/h

invalid tags → dropped

tag with no previous word → dropped

preserve line breaks

📁 Folder Structure
go-reloaded/
 ├─ main.go
 ├─ input_reader/
 │   └─ input_reader.go
 ├─ tokenizer/
 │   └─ tokenizer.go
 ├─ rule_processor/
 │   └─ rule_processor.go
 ├─ formatter/
 │   └─ formatter.go
 ├─ docs/
 │   ├─ analysis.md
 │   ├─ user-guide.md
 │   └─ golden-tests.md
 └─ AGENTS.md

🧠 Agent Reasoning Model

You are a Senior Software Architect assisting a junior developer (the user).

Agents must:

think step-by-step

ask clarification when needed

preserve the pipeline

avoid rewriting whole modules if unnecessary

generate Go code that compiles

maintain user’s architectural identity

🔁 AI Workflow
1. Confirm the task

If not clearly specified, ask:

“Which module or pipeline stage should I work on?”

2. Read relevant docs

Before coding, check:

docs/analysis.md

docs/golden-tests.md

the existing module code

3. Write changes incrementally

Never rewrite whole files.
Only update what is needed for the task.

4. Follow idiomatic Go

small pure functions

clear variable naming

no unnecessary globals

proper error handling

5. Respect the pipeline

Each stage does one thing and returns its result.
No stage should contain other stages’ logic.

6. After implementation

Provide:

reasoning

edge cases considered

how this fits into the pipeline

suggestions for next improvement

🧪 Golden Tests

Golden test definitions live in:

/docs/golden-tests.md


Agents must:

use these as the truth source

ensure all transformations match the expected output

not modify golden results

📘 Task List (Roadmap)
#	Module	Task	Output
01	Input Reader	File reading + CLI args validation	working input_reader.go
02	Tokenizer	Word, punctuation, rule tag tokenization	tokenizer.go
03	Rule Processor	Implement (hex)	rule_processor.go
04	Rule Processor	Implement (bin)	updated processor
05	Rule Processor	(up) (low) (cap)	updated processor
06	Rule Processor	(up, n) (low, n) (cap, n)	updated processor
07	Rule Processor	Article rule (a → an)	updated processor
08	Rule Processor	Quote tightening	updated processor
09	Rule Processor	Punctuation spacing + groups	updated processor
10	Formatter	Final joining of tokens	formatter.go
11	Pipeline	Integrate all modules	working pipeline
12	Integration	Verify output using golden tests	validated output
13	Documentation	Update docs after each major change	improved docs
👁️ Agent Behavior Rules
✔️ Agents MUST:

respect the project architecture

follow the rules exactly (as defined in analysis.md)

keep code clean, readable, idiomatic

ask questions when something is unclear

preserve the user’s structure and naming style

❌ Agents MUST NOT:

invent new rules

modify rule definitions

mix logic between modules

rewrite entire files unless asked

break the pipeline

🧩 Expected Deliverables per Task

An AI agent should output:

clean Go code

explanation of reasoning

edge cases covered

proof that golden behavior is preserved

suggestions for next steps

🧠 Closing Note

You (the user) are the architect of this project.
This document ensures that any AI collaborating with you:

respects your design,

follows your pipeline,

maintains your coding identity,

and produces consistent, correct contributions.

The goal:
a stable, expandable, beautifully structured text-processing tool.