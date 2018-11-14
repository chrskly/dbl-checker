# DBL Checker

Check domain names to see if they are listed in the spamhaus DBL list.

https://www.spamhaus.org/dbl/

Prints out domains which are in DBL as they are checked. Prints out percentage
of domains tested in the DBL.

## Usage

```
go build .
cat domains-to-test.txt | ./dbl-checker
```

Example

```
% cat << EOF > domains.txt
test
dbltest.com
example.com
google.com
EOF
% cat domains.txt | ./dbl-checker
test Spam domain [ 127.0.1.2 ]
dbltest.com Spam domain [ 127.0.1.2 ]
50% (2/4)
```
