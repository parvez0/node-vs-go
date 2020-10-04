import random
from locust import HttpUser, task, between

class Node_vs_Go(HttpUser):
    wait_time = between(1, 15)

    @task
    def do_get_request(self):
        self.client.get("/")

    @task
    def do_post_request(self):
        data = {"sleepTime": random.randint(50, 15000)}
        # for i in servers:
        resp = self.client.post("/", json=data, headers={"content-type": "application/json"})
        print(resp.json())

