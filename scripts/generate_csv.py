import csv
import requests

DATA_URL ='https://popcon.debian.org/by_inst'
text_data = ""

response = requests.get(DATA_URL)   
if response.status_code == 200:
    text_data = response.text
else:
    print("Failed to fetch data from the URL")

with open('data.csv', 'w', newline='') as file:
    writer = csv.writer(file)
    field = ["rank", "name", "installs", "maintainer"]
    writer.writerow(field)

    for line in text_data.split('\n') [11:-3]:
        data = line.split()
        rank = data[0]
        name = data[1]
        installs = data[2]
        maintainer = line.split('(')[-1].replace(")", "").strip()
        writer.writerow([rank, name, installs, maintainer])
        # print(line)


# create_table = """
	# create table if not exists packages (
	# 	rank 			integer not null,
	# 	name 			text not null,
	# 	installs 		integer not null,
	# 	maintainer 		text not null,
	# );
# """

