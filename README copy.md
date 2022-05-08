# VersusSystems

## The Problem
All data that flows in or out of our API gets saved to a log. This is helpful when debugging, but unfortunately
means that all of our users' personal information gets stored in plain text. In order to respect our users' privacy,
we need to clean this up.
Your Task
To handle this, you must write a scrub function to remove all personal information from a piece of data before it
is logged. This function will take in an object and perform the following transformations:
replace all "name", "username", and "password" values with the string "******"
for all "email" fields, replace only the username (the part before the @)
Assume that personal data can be nested inside objects or arrays at any depth. The final result should behave
like this (the example below is in json, but you can assume the function will be called with data types native to
the language of your choice):
