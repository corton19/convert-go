# Reads ORC files
# Used to test output of main.go

import pyorc

file = "employee.orc"

with open(file, "rb") as data:
    reader = pyorc.Reader(data)
    for row in reader:
        print(row)
