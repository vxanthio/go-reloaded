# Golden Test Suite — go-reloaded

This file contains the official golden tests for validating the go-reloaded transformation pipeline.  
Each test includes Input, Expected Output, and Rules Covered.

---

## Test 1 — Hexadecimal Conversion

### Input
1E files were deleted

### Expected Output
30 files were deleted

### Rules Covered
- Hexadecimal number conversion

---

## Test 2 — Binary Conversion

### Input
it has been 10 (bin) since the last time i saw you

### Expected Output
it has been 2 years since the last time i saw you

### Rules Covered
- Binary number conversion
- Word replacement (“years”)

---

## Test 3 — Uppercase Transformation

### Input
last night was amazing (up)!

### Expected Output
last night was AMAZING!

### Rules Covered
- (up) → uppercase previous word
- punctuation hugging

---

## Test 4 — Lowercase Transformation

### Input
did you go to the SUPERMARKET? (low)

### Expected Output
did you go to the supermarket?

### Rules Covered
- (low) → lowercase previous word

---

## Test 5 — Capitalization

### Input
I just came back from my vacations in italy (cap).

### Expected Output
I just came back from my vacations in Italy.

### Rules Covered
- (cap) → capitalize previous word
- punctuation attachment

---

## Test 6 — Uppercase Range Transformation

### Input
This is so fun (up, 2)!

### Expected Output
This is SO FUN!

### Rules Covered
- (up, n) applied to previous 2 words
- punctuation hugging

---

## Test 7 — Comma Spacing Correction

### Input
Did you bring my wallet , as i asked you to do?

### Expected Output
Did you bring my wallet, as i asked you to do?

### Rules Covered
- punctuation spacing normalization (no space before comma, one after)

---

## Test 8 — Punctuation Groups

### Input
I was thinking ... it was very kind the way you step up to defend me.

### Expected Output
I was thinking... it was very kind the way you step up to defend me.

### Rules Covered
- grouped punctuation (...)
- spacing rules

---

## Test 9 — Quote Tightening

### Input
This dress is trully ' spectacular '!

### Expected Output
This dress is trully 'spectacular'!

### Rules Covered
- remove spaces inside quotes
- punctuation hugging

---

## Test 10 — Article Correction (a → an)

### Input
I really appreciate a honest man!

### Expected Output
I really appreciate an honest man!

### Rules Covered
- article correction before vowels/h

---

# Test 11 — Full Mixed-Rules Benchmark Paragraph

### Input
I opened my 1E (hex) old folders and i realized something ... i had exactly 1010 (bin) unread messages !? Then i found a photo of my friend   '   George   '  which made me smile , because we had such a nice (up, 2) trip to athens (cap) back in the day. A honest mistake happened though : someone wrote a very WEird sentence (low) right under it. I tried to fix it but i was thinking ... maybe i should write a, helpful note instead. As the main character used to say , ' take care ' !

### Expected Output
I opened my 30 old folders and i realized something... i had exactly 10 unread messages!? Then i found a photo of my friend 'George' which made me smile, because we had such a NICE TRIP to Athens back in the day. An honest mistake happened though: someone wrote a very weird sentence right under it. I tried to fix it but i was thinking... maybe i should write an helpful note instead. As the main character used to say, 'take care'!

### Rules Covered
- hex conversion  
- binary conversion  
- (up, n) range rules  
- (low)  
- (cap)  
- quote tightening  
- punctuation grouping (..., !?)  
- punctuation normalization  
- article rule a → an (including across punctuation)  
- spacing normalization  
- full pipeline integration  

