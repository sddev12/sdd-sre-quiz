# Change Default Branch to Main

## Current Status
The repository currently has `feat/initial-spec` set as the default branch.

## Required Action
The default branch needs to be changed to `main`.

## Steps to Change Default Branch

### Option 1: Using GitHub Web Interface (Recommended)
1. Go to the repository: https://github.com/sddev12/sdd-sre-quiz
2. Click on **Settings** tab
3. In the left sidebar, click on **Branches**
4. Under "Default branch", click the switch icon next to the current default branch
5. Select `main` from the dropdown menu
6. Click **Update** to confirm the change
7. Confirm the change in the dialog that appears

### Option 2: Using GitHub CLI
If you have the GitHub CLI (`gh`) installed and authenticated:

```bash
gh repo edit sddev12/sdd-sre-quiz --default-branch main
```

### Option 3: Using GitHub API
If you have a personal access token with `repo` scope:

```bash
curl -X PATCH \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  https://api.github.com/repos/sddev12/sdd-sre-quiz \
  -d '{"default_branch":"main"}'
```

## Why This Change?
- The `main` branch is the standard default branch name in Git and GitHub
- It follows modern Git conventions
- The `main` branch contains the latest merged features and is the most up-to-date stable branch

## Branches Overview
- `main`: Latest stable code with merged features (should be default)
- `feat/initial-spec`: Feature branch with specifications (currently default)
- `feat/frontend-scaffold`: Feature branch with frontend scaffold
- `copilot/make-main-default-branch`: Branch for this change documentation

## Notes
- Changing the default branch does not delete or modify any branches
- It only affects which branch is shown by default and which branch new PRs target by default
- Users who have the repository cloned will need to update their local default branch reference

## After Changing Default Branch
Once the default branch is changed to `main`, you can update your local repository reference:

```bash
git remote set-head origin main
```

This will update your local tracking reference for the remote's default branch.
