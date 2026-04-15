# Sentence Stats

Sentence Stats is a Go tool for visualizing character and character pair frequencies in sentences. It generates histogram plots to help analyze the composition of text.

## Features

- **Character Frequency**: Visualizes the frequency of each character in the input text.
- **Character Pair Frequency**: Visualizes the frequency of character pairs (bigrams), ignoring order (e.g., "ab" and "ba" are counted together).
- **Sentence-based Analysis**: Processes input sentence by sentence (split by '.').

## Install

### GitHub Releases
Download binaries from: https://github.com/arran4/sentencestats/releases

### Homebrew
brew tap arran4/homebrew-tap
brew install sentencestats

### Go install
go install github.com/arran4/sentencestats/cmd/sentencestats@latest

### Native packages
- Debian/Ubuntu (`.deb`): see Releases assets
- RPM (`.rpm`): see Releases assets
- Alpine (`.apk`): see Releases assets
- Arch (`.pkg.tar.zst` or repo): see Releases assets


## Usage

The tools read from standard input and output a PNG file.

### Character Frequency

```bash
echo "This is an example. This is also a test. This is also a demo." | sentencestats characters -o characters-example.png
```

Output:

![](characters-example.png)

### Character Pair Frequency

```bash
echo "This is an example. This is also a test. This is also a demo." | sentencestats character-pairs -o character-pairs-example.png
```

Output:

![](character-pairs-example.png)

## Development

To run the tools from source:

```bash
go run ./cmd/sentencestats/ characters -o out.png < input.txt
go run ./cmd/sentencestats/ character-pairs -o out.png < input.txt
```

To run tests:

```bash
go test ./...
```
