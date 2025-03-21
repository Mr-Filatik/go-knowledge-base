# Workflows

This file describes all workflows used in the project and the rules for working with them:

## Check Merge Source

The workflow in `check-merge-source.yml` checks from which branch the `Pull Request` to the `main` branch originates and allows it only for the `work` branch.

Complete Condition:

```bash
if [ "${{ github.event.pull_request.head.ref }}" != "work" ]; then
  echo "Merge is not allowed from branch other than 'work'."
  exit 1
fi
```

For some reason (should work without it) the check was running on a branch other than main, so I added an additional condition:

```yaml
if: github.event.pull_request.base.ref == 'main'
```

## Go Build

So far it's not working the way I'd like it to, it launches 2 jobs at a time on both push and pull request.

## Go Test

So far it's not working the way I'd like it to, it launches 2 jobs at a time on both push and pull request.

# How to use

Additional useful data:

## How to set a merge limit in GitHub

To disallow merge on failing tests, you need to create a rule in GitHub. How to do it:
1. Go to `Settings` -> `Branches` -> `Add branch ruleset` (or `Settings` -> `Rules` -> `Rulesets` -> `New ruleset` -> `New branch ruleset`);
2. Specify `Ruleset Name`;
3. Select `Targets`;
4. Select an item `Require status checks to pass`;
5. Add the necessary checks through `Add checks`.
