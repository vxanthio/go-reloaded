Text Editing & Auto-Correction Tool — Analysis Document

Author: Βασιλική Ξανθιώτη
Date: 19/10/2025
Version: 1.0

1. Περιγραφή του προβλήματος

Το πρόγραμμα παίρνει ένα κείμενο και το διορθώνει ή το μορφοποιεί σύμφωνα με συγκεκριμένους κανόνες.

Δεν μπορεί να καταλάβει το νόημα του κειμένου, οπότε του δίνουμε ξεκάθαρους κανόνες για να φέρει το επιθυμητό αποτέλεσμα: ένα πιο σωστό και ευανάγνωστο κείμενο.

Το input είναι ένα αρχείο (file) με το κείμενο.

Το πρόγραμμα επεξεργάζεται το input, εφαρμόζει τους κανόνες και αποθηκεύει το αποτέλεσμα σε ένα νέο αρχείο (output).

2. Κανόνες που πρέπει να ακολουθηθούν

2.1 (hex)

Αν η προηγούμενη λέξη είναι δεκαεξαδικός αριθμός (hex), το πρόγραμμα την μετατρέπει σε δεκαδικό.

Παράδειγμα: 1E (hex) → 30

2.2 (bin)

Αν η προηγούμενη λέξη είναι δυαδικός αριθμός (bin), το πρόγραμμα την μετατρέπει σε δεκαδικό.

Παράδειγμα: 10 (bin) → 2

2.3 (up), (low), (cap)

(up) → κεφαλαία

(low) → μικρά

(cap) → πρώτη κεφαλαία γράμμα

Παράδειγμα: amazing (up) → AMAZING

2.4 (up, n), (low, n), (cap, n)

Ο αριθμός δείχνει πόσες προηγούμενες λέξεις θα μετατραπούν.

Παράδειγμα: so fun (up, 2) → SO FUN

2.5 Σημεία στίξης

Δεν υπάρχει κενό πριν από σημεία στίξης (.,!?;:) και υπάρχει κενό μόνο μετά, εκτός αν είναι στο τέλος πρότασης.

Παράδειγμα: wallet , as → wallet, as

2.6 Εξαιρέσεις στα σημεία στίξης

Για ... ή !?, τα σημεία στίξης μένουν ενωμένα και βάζει κενό μόνο μετά το τελευταίο.

Παράδειγμα: I was thinking ... we should go → I was thinking... we should go

2.7 Εισαγωγικά

Τα εισαγωγικά πρέπει να έχουν ζευγάρι.

Δεν υπάρχει κενό μετά το πρώτο ή πριν το δεύτερο εισαγωγικό, ακόμα και αν υπάρχουν πολλές λέξεις μέσα.

Παράδειγμα: ‘ fun ’ → ‘fun’

2.8 “a” → “an”

Αν η λέξη “a” προηγείται φωνήεντος, γίνεται “an”.

Παράδειγμα: a apple → an apple

3. Σύγκριση Pipeline και FSM αρχιτεκτονικής

Pipeline:

Διαχωρίζει το πρόβλημα σε στάδια (modules).

Κάθε στάδιο εκτελεί συγκεκριμένη λειτουργία και επικοινωνεί με το επόμενο.

Πλεονεκτήματα:

Μικρά και ελέγξιμα στάδια

Εύκολη προσθήκη/αφαίρεση λειτουργιών

Μειονέκτημα: Πολλαπλές περασιές ίσως μειώνουν την απόδοση

FSM (Finite State Machine):

Βασίζεται σε καταστάσεις και μεταβάσεις.

Πλεονέκτημα: Κατάλληλο για περίπλοκες αλληλεπιδράσεις

Μειονέκτημα: Πιο δύσκολη υλοποίηση

4. Επιλογή Αρχιτεκτονικής

Επιλέχθηκε η Pipeline, καθώς κάθε κανόνας μπορεί να εφαρμοστεί ανεξάρτητα.

Η προσθήκη νέων κανόνων γίνεται εύκολα με επιπλέον στάδιο στο pipeline.

5. Skeleton Plan (Pipeline Flow)

Stage 1:Διαβάζει το αρχείο εισόδου και αποθηκεύει το περιεχόμενο.

Stage 2:Χωρίζει το κείμενο σε λέξεις/tokens.

Stage 3:Ελέγχει κάθε token και εφαρμόζει τους κανόνες (hex, bin, up, low, cap, a→an, punctuation, quotes).

Stage 4:Αποθηκεύει το τελικό κείμενο σε νέο αρχείο.

6. Golden Test Set

Test 1: 1E files were deleted → 30 files were deleted

Test 2: it has been 10 (bin) since the last time i saw you → it has been 2 years since the last time i saw you

Test 3: last night was amazing (up)! → last night was AMAZING!

Test 4: did you go to the SUPERMARKET? (low) → did you go to the supermarket?

Test 5: I just came back from my vacations in italy (cap). → I just came back from my vacations in Italy.

Test 6: This is so fun (up, 2)! → This is SO FUN!

Test 7: Did you bring my wallet , as i asked you to do? → Did you bring my wallet, as i asked you to do?

Test 8: I was thinking ... it was very kind the way you step up to defend me. → I was thinking... it was very kind the way you step up to defend me.

Test 9: This dress is trully ' spectacular '! → This dress is trully 'spectacular'!

Test 10: I really appreciate a honest man! → I really appreciate an honest man!said,'Catch you on the flip side, buddy. Have a great night'."
Test 11:
Input:I was looking through my old photos and i couldn’t believe it , it has been 10 (bin) years since the last time we went for vacation. I would never deleted this 1E photos , we had such a good time (up, 2) in paris (cap) and seeing them remind me that. We should go to grab a coffee some time , it would be so fun (up) to CATCH (low) up. I was thinking ... back then we was in the same bookclub. Do you still read books? I am reading something but i dont want to spoil you , this book is trully ' amazing ' tho. I would love to hear a opinion from you. As the main character said, 'Catch you on the flip side, buddy. Have a great night'.
Output:I was looking through my old photos and i couldn’t believe it, it has been 2 years since the last time we went for vacation. I would never deleted this 30 photos, we had such a GOOD TIME in Paris and seeing them remind me that. We should go to grab a coffee some time, it would be so FUN to catch up. I was thinking... back then we was in the same bookclub. Do you still read books? I am reading something but i dont want to spoil you, this book is trully 'amazing' tho. I would love to hear an opinion from you. As the main character said, 'Catch you on the flip side, buddy. Have a great night'.