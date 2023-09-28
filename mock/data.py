import pandas as pd
import random

# Generate 100 rows of sample data
data = {
    'username': [f'user{random.randint(1,10)}' for i in range(1, 101)],
    'company_name': [f'Company {chr(65 + (i % 26))}' for i in range(1, 101)],
    'payment': [random.randint(500, 2000) for _ in range(1, 101)]
}

# Create a DataFrame
df = pd.DataFrame(data)
df.to_csv('sample_data.csv', index=False)

# Display the first few rows of the DataFrame
print(df.head())
