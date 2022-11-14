import random
import string

from locust import HttpUser, task


def randomstring(length):
    return ''.join(random.choice(string.ascii_letters) for i in range(length))


class WOJUser(HttpUser):

    def on_start(self):
        self.username = randomstring(16)
        self.nickname = randomstring(16)
        self.password = randomstring(16)

        with self.client.post("/api/v1/user/create", data={
            "username": self.username,
            "nickname": self.nickname,
            "password": self.password
        }) as resp:
            j = resp.json()
            if j["code"] != 0:
                resp.failure("create user failed")
            else:
                self.token = j["body"]

    @task
    def view_problem(self):
        pid = []

        with self.client.post("/api/v1/problem/search") as resp:
            j = resp.json()
            if j["code"] != 0:
                resp.failure("search problem failed")
            else:
                for p in j["body"]:
                    pid.append(p["meta"]["ID"])

        for p in pid:
            with self.client.post("/api/v1/problem/details", data={"pid": pid}) as resp:
                if resp.json()["code"] != 0:
                    resp.failure("view problem failed")

    @task
    def submit_code(self):
        code = """
            #include<iostream>
            using namespace std;
            int main() {
                int a, b;
                cin >> a >> b;
                cout << a + b;
                return 0;
            }
        """

        with self.client.post("/api/v1/submission/create",
                              headers={"Authorization": "Bearer " + self.token},
                              data={"pid": 5, "language": "cpp", "code": code}
                              ) as resp:
            if resp.json()["code"] != 0:
                resp.failure("submit code failed")

    @task
    def view_submission(self):
        with self.client.post("/api/v1/submission/query", data={"pid": 5, "limit": 100}) as resp:
            if resp.json()["code"] != 0:
                resp.failure("view submission failed")
