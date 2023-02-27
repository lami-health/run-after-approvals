# Run after approvals

GitHub Action that permit the pipeline run after some approvals

## Example

```sh
GITHUB_REPOSITORY=website GITHUB_PULL_REQUEST=720 GITHUB_TOKEN=<github-token> ./run-after-approvals
```

> Output: 2/2 Approvals - APPROVED

## Configuration Tips

On GitHub workflow the user can pass the current pull request number by using this:

```yaml
env:
  PR_NUMBER: ${{ github.event.number }}
```
