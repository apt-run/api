



import requests

url = 'https://popcon.debian.org/by_inst'

response = requests.get(url)
if response.status_code == 200:
    text_data = response.text

    data_lines = text_data.split('\n') 
    count = 0

    for line in data_lines[8:-3]:
        # if count == 20: break
        # count = count + 1

        data = line.split()
        # rank = data[0]
        # name = data[1]
        # installs = data[2]
        # maintainer = line.split('(')[-1].replace(")", "") 
        print(data)





else:
    print("Failed to fetch data from the URL")

