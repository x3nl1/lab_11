import requests
import os

BASE = os.getenv("BASE_URL", "http://localhost:8080")

def calculate(values):
    return requests.post(f"{BASE}/calculate", json={"values": values}).json()

def get_token():
    return requests.get(f"{BASE}/token").json()["token"]

def protected(token):
    return requests.get(
        f"{BASE}/protected",
        headers={"Authorization": f"Bearer {token}"}
    ).status_code