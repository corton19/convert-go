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

# For 3 col
C1 = []
C2 = []
C3 = []

for i in range(0,len(df)):
    C1.append(df.loc[i][0])
    C2.append(df.loc[i][1])
    C3.append(df.loc[i][2])

C1 = np.array(Name)
C2 = np.array(Age)
C3 = np.array(Country)

###
# [EXAMPLES] for plotting purposes
df.plot(x=columns[0], y=columns[1], kind="bar")
