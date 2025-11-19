# User Guide — go-reloaded  
**Location:** `.gitea/docs/user-guide.md`  
**Version:** 2.0  

---

## Overview

**go-reloaded** is a command-line text editing and auto-correction tool.  
It reads a text file, applies a series of deterministic transformation rules, and outputs a fully corrected version of the text.

The tool does **not** understand meaning or grammar.  
It performs strictly mechanical operations based on predefined rules.

This document explains:

- what the program does,  
- how each rule works,  
- why the chosen architecture is appropriate,  
- how the processing pipeline is structured.

---

## Problem Description

The program takes an **input file** containing text.  
It processes the content, applies transformation rules, and writes the final corrected text into a **new output file**.

The purpose is to transform low-quality or unformatted text into a clean and readable version using purely deterministic rules.

The program:

- does **not** understand semantics,  
- does **not** interpret context,  
- does **not** decide what “sounds better,”  
- only performs rule-driven transformations.

---

## Transformation Rules

Below is the full list of transformations.

### ## Number Conversions

### **(hex) — Hexadecimal → Decimal**
If the previous token is a valid hexadecimal number (base 16), convert it to decimal.  
Invalid hex results in keeping the original word and removing the tag.

**Example:**  
`1E (hex)` → `30`

---

### **(bin) — Binary → Decimal**
Convert the previous token from binary (base 2) to decimal.  
Invalid binary removes the tag only.

**Example:**  
`10 (bin)` → `2`

---

## Case Transformation Rules

### **(up)**  
Convert previous word to uppercase.  
`amazing (up)` → `AMAZING`

### **(low)**  
Convert previous word to lowercase.  
`LOUD (low)` → `loud`

### **(cap)**  
Capitalize previous word.  
`bridge (cap)` → `Bridge`

### **Range Transformations — (up, n), (low, n), (cap, n)**  
Apply transformation to previous **n** words (words only, no punctuation).

`so fun (up, 2)` → `SO FUN`

---

## Punctuation Rules

### ### Basic Spacing Rules

Punctuation marks (`. , ! ? : ;`) must:

- attach directly to the previous word,  
- have **no space before**,  
- have **one space after** (unless followed by newline).

Example:  
`wallet , as` → `wallet, as`

---

### ### Punctuation Groups

Certain punctuation combinations must be treated as single units:

```
...
!?
?!
```

Rules:

- remain grouped,  
- attach to previous word,  
- followed by exactly one space.

**Example:**  
`I was thinking ... we` → `I was thinking... we`

---

## Quote Rules

Single quotes `'...'` must:

- appear in pairs,  
- contain **no spaces inside**,  
- preserve spacing outside,  
- allow multi-word content.

**Example:**  
`' spectacular '` → `'spectacular'`

---

## Article Correction Rule

Convert **a** → **an** when the next word starts with:

- a vowel (`a, e, i, o, u`)  
- letter `h`  

Case-insensitive and works through punctuation.

**Example:**  
`a, honest` → `an, honest`

---

## Edge Cases & Error Handling

### Invalid numbers  
`ZZ (hex)` → `ZZ`

### Tags without previous word  
`(up) Hello` → remove tag

### Malformed tags  
Ignored safely:

```
(up,)
(low, 0)
(bin, xyz)
(hex
```

### Word counting in range rules  
Punctuation and spaces do **not** count.

### Line breaks  
Must be preserved exactly.

---

## Architectural Comparison

The go-reloaded project can theoretically be implemented using:

---

### ## Pipeline Architecture (Chosen)

**Characteristics:**

- Each stage performs one well-defined transformation.
- Output of one stage flows into the next.
- Highly modular and testable.
- Easy to add/remove rules.
- Clear separation of concerns.

**Advantages:**

- predictable execution order  
- clean debugging  
- independent testing per module  
- easy auditor verification  
- simple extensibility  

**Disadvantages:**

- multiple passes may be slightly slower (not relevant here)

---

### ## FSM Architecture (Rejected)

**Characteristics:**

- One large machine controlling states & transitions
- Single-pass logic
- Rule handling deeply interconnected

**Advantages:**

- fast, memory-efficient

**Disadvantages:**

- hard to read  
- extremely hard to maintain  
- not suitable for many loosely related rules  
- difficult for peer auditing  
- rule interactions become tangled  

---

### **Final Decision:** Pipeline Architecture

It is more readable, modular, and aligns with Zone01 audit expectations.  
Every rule is isolated and testable, ensuring deterministic and reproducible behavior.

---

## Pipeline Processing Flow

The full processing flow is:

```
Input File
   ↓
Stage 1 — Input Reader
   ↓
Stage 2 — Tokenizer
   ↓
Stage 3 — Rule Processor
   ↓
Stage 4 — Formatter
   ↓
Output File
```

### Stage Responsibilities:

#### **Input Reader**
Reads file, returns raw text.

#### **Tokenizer**
Splits into tokens:
- words  
- punctuation  
- punctuation groups  
- rule tags  
- spaces  

Preserves structure.

#### **Rule Processor**
Applies all transformations in deterministic order:
- hex → decimal  
- bin → decimal  
- case rules  
- ranged case rules  
- quotes  
- article corrections  
- punctuation rules  
- spacing normalization  

#### **Formatter**
Joins tokens, reconstructs final text, preserves newlines.

---

## Example Transformations

### Example 1 — Hex + Case
Input:  
`We saw 1E (hex) birds. It was amazing (up)!`

Output:  
`We saw 30 birds. It was AMAZING!`

---

### Example 2 — Article + Range
Input:  
`a honest friend said it was so fun (up, 2)`

Output:  
`an honest friend said it was SO FUN`

---

### Example 3 — Punctuation Groups + Quotes
Input:  
`I was thinking ... this is ' cool ' tho`

Output:  
`I was thinking... this is 'cool' tho`

---

## Using the Program

### Running the tool

```
go run main.go <input_file> <output_file>
```

Example:

```
go run main.go samples/input.txt samples/output.txt
```

---

## Relation to Golden Tests

Your solution is **correct only if** the output matches:

```
.gitea/docs/golden-tests.md
```

Golden tests:

- define the required output  
- cover every rule  
- include tricky edge cases  
- include a long benchmark paragraph  
- are used during peer audits  

Your program must match them **character by character**.

---

## Summary

The user guide has explained:

- the purpose of the tool  
- all supported rules  
- rule examples  
- edge case handling  
- architectural comparison  
- pipeline flow  
- how to run the program  
- its connection to golden tests  

go-reloaded is a deterministic, auditable, rule-driven text transformation system implemented with a clean Pipeline Architecture.
