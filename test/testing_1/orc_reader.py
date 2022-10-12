# Reads ORC files
# Used to test output of main.go
#%%
import pandas as pd
import pyorc    # pandas?

data = open('test_sat.orc','rb')
reader= pyorc.Reader(data)
columns = reader.schema.fields
#print(f'fields: {columns}')
# r in front of string: Tells the Python interpreter to treat backslashes as a literal (raw) character
# Normally, Python uses backslashes as escape characters

columns = [col_name for col_idx, col_name in sorted([(reader.schema.find_column_id(c),c) for c in columns])]

df = pd.DataFrame(reader,columns=columns)
print(df)
    


#ss = df.to_numpy()
#print(ss)
# %%
