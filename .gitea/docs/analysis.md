# go-reloaded — Technical Analysis Document

**Author:** Vasiliki Xanthióti  
**Role:** System Architect  
**Project:** go-reloaded (Text Processing & Auto-Correction Tool)  
**Version:** 1.2  
**Location:** /docs/analysis.md  

---

## Problem Description

The go-reloaded project is a command-line text-processing tool that reads an input file, applies a deterministic sequence of transformation rules, and produces a corrected, fully formatted output file.

The tool does not perform semantic analysis. All transformations are strictly mechanical and rule-driven.  
The system must:

- preserve the structure of the original text (line breaks, token ordering),
- apply rules in a predictable and fixed order,
- match the expected output exactly (golden tests),
- remain stable when encountering malformed tags or irregular input,
- maintain a clean, modular architecture suitable for audit and extension.

---

## Transformation Rules

### Number Conversions

#### (hex)
Converts the previous token from hexadecimal to decimal.  
Invalid hex → keep original word, remove tag.

**Example:**  
`1E (hex)` → `30`

#### (bin)
Converts the previous token from binary to decimal.  
Invalid binary → keep word, remove tag.

**Example:**  
`10 (bin)` → `2`

---

### Case Transformations

#### (up)
Uppercases the previous word.  
`nice (up)` → `NICE`

#### (low)
Lowercases the previous word.  
`LOUD (low)` → `loud`

#### (cap)
Capitalizes the previous word.  
`bridge (cap)` → `Bridge`

#### Range Variants: (up, n), (low, n), (cap, n)
Applies transformation to the previous **n** words (counting only word tokens, not punctuation or spaces).

**Example:**  
`so fun (up, 2)` → `SO FUN`

---

### Punctuation Rules

Punctuation must:

- attach directly to the previous word,
- have **no space before**,
- have **one space after** (unless followed by newline).

Applies to: `. , ! ? : ;`

**Example:**  
`wallet , as` → `wallet, as`

---

### Punctuation Groups

Groups such as:

- `...`
- `!?`
- `?!`

must be treated as atomic units:

- remain grouped,
- attach to previous word,
- include exactly one space after the group.

**Example:**  
`I was thinking ... we` → `I was thinking... we`

---

### Quotes

Single quotes `'...'` must:

- appear in matched pairs,
- contain **no inner spacing**,
- preserve external spacing,
- support multi-word regions.

**Example:**  
`' amazing '` → `'amazing'`

---

### Article Rule — a → an

Changes **a** to **an** if the next word begins with:

- a vowel (a, e, i, o, u),
- the letter h.

Case-insensitive, and works across punctuation:

**Example:**  
`a, honest` → `an, honest`

---

## Edge Cases & Error Handling

The system must remain predictable and robust under malformed input.

- **Invalid Numbers**  
  `ZZ (hex)` → `ZZ`
- **Tags with No Previous Word**  
  `(up) Hello` → remove tag
- **Malformed Tags**  
  `(up,)`, `(low, 0)`, `(hex`, `(bin, abc)` → ignored safely
- **Word Counting**  
  Punctuation and spaces do not count as words.
- **Article Across Punctuation**  
  `a, honest` → `an, honest`
- **Line Breaks**  
  Always preserved.

---

## Architecture Decision

### Options Considered

#### Pipeline Architecture
- Modular and sequential
- Easy to extend
- Clear rule ordering
- Isolated transformation stages
- Excellent auditability

#### FSM Architecture
- Efficient single-pass design
- Complex state transitions
- Harder to maintain
- Less transparent for auditors

### Chosen Architecture — Pipeline

The Pipeline architecture was selected because it:

- matches the deterministic nature of the rules,
- isolates transformation logic per module,
- provides predictable and reproducible results,
- simplifies debugging and auditing,
- scales easily as new rules are added.

FSM would introduce unnecessary complexity and reduce clarity.

---

## System Design (High-Level)

### Stage 1 — Input Reader
Reads file contents and returns raw string.  
**Location:** `/internal/inputreader/inputreader.go`

### Stage 2 — Tokenizer
Splits text into tokens:

- words
- spaces
- punctuation
- punctuation groups
- rule tags

Preserves original ordering and line breaks.  
**Location:** `/internal/tokenizer/tokenizer.go`

### Stage 3 — Rule Processor
Applies transformations in fixed order:

1. hex conversion  
2. bin conversion  
3. case transformations  
4. ranged case transformations  
5. quote tightening  
6. article correction  
7. spacing normalization  
8. punctuation attachment  

**Location:** `/internal/ruleprocessor/ruleprocessor.go`

### Stage 4 — Formatter
Reconstructs final string, ensuring punctuation spacing and line preservation.  
**Location:** `/internal/formatter/formatter.go`

---

## Data Flow

Input File
↓
Input Reader
↓
Tokenizer
↓
Rule Processor
↓
Formatter
↓
Output File

## Testing Strategy

Golden tests are located at:

.gitea/docs/golden-tests.md


Tests validate:

- correctness of each rule,
- correct ordering of rule application,
- punctuation/spacing normalization,
- exact match between output and expected result,
- stability under malformed input,
- pipeline-wide integration through a long paragraph.

The golden test suite includes:

- base rule tests,
- tricky custom edge cases,
- a long mixed-rule paragraph,
- benchmark audit paragraph.

---

## Benchmark Material (Audit-Oriented Test Case)

The following paragraph is included to evaluate full pipeline integration, rule overlap, and real-world text behavior. It is part of the golden test suite and used during peer audits.

I was looking through my old photos and i couldn’t believe it , it has been 10 (bin) years since the last time we went for vacation. I would never deleted this 1E photos , we had such a good time (up, 2) in paris (cap) and seeing them remind me that. We should go to grab a coffee some time , it would be so fun (up) to CATCH (low) up. I was thinking ... back then we was in the same bookclub. Do you still read books? I am reading something but i dont want to spoil you , this book is trully ' amazing ' tho. I would love to hear a opinion from you. As the main character said, 'Catch you on the flip side, buddy. Have a great night'.

---

## Success Criteria

The system is correct when:

- all golden tests pass,
- output matches expected text exactly,
- no crashes occur with malformed input,
- architecture remains clean and modular,
- documentation supports reproducibility and auditing,
- benchmark material yields stable and correct results.

---

## Conclusion

This Analysis Document defines:

- the complete transformation rules,
- edge case behavior,
- architectural rationale,
- pipeline structure,
- testing approach,
- and benchmark material.

It serves as the authoritative technical blueprint for the go-reloaded implementation.

