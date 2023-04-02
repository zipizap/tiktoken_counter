# tiktoken_counter
tiktoken counter in golang, for estimating the amount of  gpt tokens in a text

# Example
```
❯ ll
total 6836
drwxrwxr-x  3 ub ub    4096 abr  2 02:58 ./
drwxrwxr-x 11 ub ub    4096 abr  2 02:56 ../
drwxrwxr-x  8 ub ub    4096 abr  2 02:57 .git/
-rw-rw-r--  1 ub ub      59 abr  2 02:57 .gitignore
-rwxrwx---  1 ub ub     616 abr  2 02:57 go_build.sh*
-rwxrwx---  1 ub ub     513 abr  2 02:57 go_clean.sh*
-rw-rw-r--  1 ub ub     182 abr  2 02:57 go.mod
-rwxrwxr-x  1 ub ub     105 abr  2 02:57 go_mod_and_run.sh*
-rwxrwx---  1 ub ub     789 abr  2 02:57 go_modules_update.sh*
-rwxrwx---  1 ub ub     526 abr  2 02:57 go_run.sh*
-rw-rw-r--  1 ub ub     511 abr  2 02:57 go.sum
-rw-rw-r--  1 ub ub   35149 abr  2 02:56 LICENSE
-rwxrwx---  1 ub ub      23 abr  2 02:57 local_env.source*
-rw-rw-r--  1 ub ub    1702 abr  2 02:57 main.go
-rw-rw-r--  1 ub ub     390 abr  2 02:57 QUESTION.current.txt
-rw-rw-r--  1 ub ub    1707 abr  2 02:57 QUESTION.history.txt
-rw-rw-r--  1 ub ub      98 abr  2 02:58 README.md
-rwxrwxr-x  1 ub ub 6897094 abr  2 02:57 tiktoken_counter*
❯ ./tiktoken_counter
Usage:   tiktoken_counter <'gpt-3.5-turbo' or 'gpt-4'> <filename>
❯ ./tiktoken_counter gpt-3.5-turbo main.go
469
469 tokens        # << each time the file is modified, it prints the new number of tokens
474 tokens
469 tokens
^CExiting...


```
