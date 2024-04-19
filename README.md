# CRAPI (Automated Exploits)

Automate exploits for the CRAPI application, available at [http://{you_know_ip}:8888](http://{you_know_ip}:8888).

## Rules for Working Together

1. **Create Your Own Branches:** Create branches for your work with descriptive names (e.g., `challenge1`, `feature-x`, `bugfix-y`).
   
2. **Merge Changes from Main:** Before starting your work, ensure your branch is up-to-date by merging changes from the `main` branch.
   
   ```bash
   git checkout main
   git pull origin main
   git checkout {your-branch-name}
   git merge main


3. # Assuming you are on your branch and want to merge to main
    ```bash
        git checkout main
        git pull origin main
        git checkout {your-branch-name}
        git merge main

    # Resolve any conflicts, commit changes, then push
        git push origin {your-branch-name}
