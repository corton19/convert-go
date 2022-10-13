# Reads ORC files
# Used to test output of main.go
#%%
import pandas as pd
import pyorc
import numpy as np
import matplotlib.pyplot as plt

#%%
data = open('employee.orc','rb')
reader= pyorc.Reader(data)
columns = reader.schema.fields
#print(f'fields: {columns}')
# r in front of string: Tells the Python interpreter to treat backslashes as a literal (raw) character
# Normally, Python uses backslashes as escape characters

columns = [col_name for col_idx, col_name in sorted([(reader.schema.find_column_id(c),c) for c in columns])]

df = pd.DataFrame(reader,columns=columns)
print(df)

Name = []
Age = []
Country = []
Skills = []

for i in range(0,len(df)):
    Name.append(df.loc[i][0])
    Age.append(df.loc[i][1])
    Country.append(df.loc[i][2])
    Skills.append(df.loc[i][3])

Name = np.array(Name)
Age = np.array(Age)
Country = np.array(Country)
Skills = np.array(Skills)

# [EXAMPLES] for plotting purposes
df.plot(x="name", y="age", kind="bar")

# Convert to .csv
print("Starting CSV conversion...")
df.to_csv("employee.csv")
print("Finished.")

# %%
