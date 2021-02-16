# Go-Update Test

## Creating a release

First step is to tag a release. You need to make sure the git repo is up to date before you do this. 
To do that you need to run the following commands.

```shell
git add .
git commit -m "commit message"
git tag -a v1.0.1 -m "tag message" # replace version with your version
git push origin v1.0.1
```

Once that is done you are ready to release which can be done by running the following command

```shell
goreleaser release
```