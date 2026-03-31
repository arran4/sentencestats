# Let's add `version: 1` or something to `.golangci.yml`
# The error says "unsupported version of the configuration: \"\""
# Which means it found `version: ""` or it defaulted to `""` because it was absent.
# If I add `version: 2` or `version: 1` or whatever is standard in 2026.
# Let's check the old `.goreleaser.yml` to see if it had `version: 2`. Yes it did.
