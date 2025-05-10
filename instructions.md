git add .
git rm -r kraken
git commit -m "fix: expose CLI entrypoint at root for go install"
git push origin main

git tag -d v0.1.5
git push origin :refs/tags/v0.1.5
git tag v0.1.5
git push origin v0.1.5
