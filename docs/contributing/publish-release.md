# Publishing a new release

This document explains how to publish a new release.

## Steps

First of all, update the `VERSION` file with the desired release version.
For instance:
```
0.13.0
```

Then run: `make prepare-release`

Once done create a new release branch and submit a Pull Request:

```
VERSION=$(cat VERSION)
git checkout -b release/v$VERSION
git commit -am "Prepare release $VERSION"
git push origin/release/v$VERSION
gh pr create --title "Prepare release v$VERSION" --base main --head release/v$VERSION
```

Once the PR has been approved and merged, publish a new tag from the `main` branch

```
git tag v0.13.0
git push origin --tags
```

Finish by creating a new release on Github.