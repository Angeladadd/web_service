import requests

r = requests.post("http://localhost:8080", data="hhhhhhh")

print(r.text)