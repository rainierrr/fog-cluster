from locust import HttpUser, task, between


class QuickstartUser(HttpUser):
    wait_time = between(1, 1)

    @task
    def get_contents(self):
        self.client.get("https://www.google.com/")
