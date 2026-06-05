comparison_methods.go

1. Pass any two structs into CompareStructs.
2. diff.Diff from r3labs will create a list of changes
3. The for loop formats the changes in a user-friendly format, e.g. +{1}, {A->B}
4. DiffStrings uses difflib (based on Python's difflib) to show String differences character by character, e.g. "{Alic<ia->e>}" for "Alicia" and "Alice".
