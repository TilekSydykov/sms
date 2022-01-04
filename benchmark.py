import requests
import time

for i in range(1000):
    start = time.time()
    r = requests.get("http://localhost:8090/book/1")
    end = time.time()
    print(end - start)
