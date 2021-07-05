# mock-data-generator

# Plan
This repo is to generate mock data

# Input
There might me a input ommand/files which tells how many rows, what properties are going to generate.    
For example, 
```
row 500

columns: 

id: integer (increased)
name: string
age: integer
address: string
some_id: uuid

```

It might support various column types(uuid, name, random text, id, etc)

# Output
Output can be json or csv which has the given conditions

# Requirement
1. Using Golang
2. Work with concurrency
3. TBD
