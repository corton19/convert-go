# Reads ORC files - Example file

import pandas as pd
import pyorc
import numpy as np
import matplotlib.pyplot as plt

data = open('test.orc','rb')
reader= pyorc.Reader(data)
columns = reader.schema.fields

columns = [col_name for col_idx, col_name in sorted([(reader.schema.find_column_id(c),c) for c in columns])]

df = pd.DataFrame(reader,columns=columns)
print(df)

Name = []
Age = []
Country = []

for i in range(0,len(df)):
    Name.append(df.loc[i][0])
    Age.append(df.loc[i][1])
    Country.append(df.loc[i][2])

Name = np.array(Name)
Age = np.array(Age)
Country = np.array(Country)

###
# [EXAMPLES] for plotting purposes
df.plot(x=columns[0], y=columns[1], kind="bar")
