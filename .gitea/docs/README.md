# go-reloaded

go-reloaded is a command-line text processing tool that reads an input file, applies a series of deterministic transformation rules, and writes a corrected, fully formatted output to a new file.

The program does not interpret meaning — all behavior is strictly mechanical and rule-based.  
Its architecture is modular, auditable, and designed according to Zone01 guidelines.

---

## Installation

Clone the repository:

```bash
git clone https://platform.zone01.gr/git/vxanthio/go-reloaded.git
```

Move into the project directory:

```bash
cd go-reloaded
```

Build (optional):

```bash
go build
```

---

## Usage

Run the program using:

```bash
go run main.go <input_file> <output_file>
```

Example:

```bash
go run main.go input.txt output.txt
```

---

## Project Structure

```
go-reloaded/
 ├─ main.go
 ├─ internal/
 │   ├─ inputreader.go
 │   ├─ tokenizer.go
 │   ├─ ruleprocessor.go
 │   └─ formatter.go
 └─ .gitea/
     ├─ AGENTS.md
     └─ docs/
         ├─ analysis.md
         ├─ user-guide.md
         └─ golden-tests.md
```

**internal/** contains all pipeline modules:

- `inputreader.go` — reads input files  
- `tokenizer.go` — produces tokens (words, punctuation, tags)  
- `ruleprocessor.go` — applies hex/bin, case rules, article correction, quotes, punctuation rules  
- `formatter.go` — reconstructs final output

---

## Documentation

Full documentation is available inside the `.gitea/docs` directory.

### 📘 Technical Analysis  
Detailed rule definitions, architecture decisions, and benchmark material:  
`.gitea/docs/analysis.md`

### 📘 User Guide  
Friendly explanation of each rule with examples:  
`.gitea/docs/user-guide.md`

### 📘 Golden Tests  
Authoritative input → output tests used for validation & audits:  
`.gitea/docs/golden-tests.md`

### 📘 AI Collaboration Guidelines  
Rules for AI agents working on this project:  
`.gitea/AGENTS.md`

---

## Tests

The project uses a **golden test suite** to ensure correctness.

📌 Location:  
`.gitea/docs/golden-tests.md`

Golden tests validate:

- exact character-for-character correctness  
- rule ordering  
- punctuation spacing  
- multi-rule interactions  
- audit benchmark paragraph  

You must match the expected output exactly.

---

## License

This project is part of the Zone01 curriculum.  
It is intended for educational and audit purposes.

