# Needs optimization to batch insert
# 39078



import requests
from dotenv import dotenv_values
from supabase import create_client, Client

config = dotenv_values(".env")
SUPABASE_URL = config["SUPABASE_URL"]
SUPABASE_KEY = config["SUPABASE_KEY"]

supabase_client: Client = create_client(SUPABASE_URL, SUPABASE_KEY)

DATA_URL ='https://popcon.debian.org/by_inst'
extracted_data = []
text_data = ""

response = requests.get(DATA_URL)   
if response.status_code == 200:
    text_data = response.text
else:
    print("Failed to fetch data from the URL")

data_lines = text_data.split('\n') 
for line in data_lines[11:-3]:
    data = line.split()
    rank = data[0]
    name = data[1]
    installs = data[2]
    maintainer = line.split('(')[-1].replace(")", "").strip()
    # extracted_data.append({
    #     "rank": rank,
    #     "name": name,
    #     "installs": installs,
    #     "maintainer": maintainer
    #     })
    data, count = supabase_client.table('packages').insert({
        "rank": rank,
        "name": name,
        "installs": installs,
        "maintainer": maintainer
        }).execute()
    print(data)
    print(count)

# for line in extracted_data[0:20]:
#     print(line)

# {"id": 1, "name": "Denmark"}
# data, count = supabase_client.table('packages').insert(extracted_data).execute()

# create_table = """
# 	create table if not exists packages (
# 		rank 			integer not null,
# 		name 			text not null,
# 		installs 		integer not null,
# 		maintainer 		text not null,
# 	);
# """

