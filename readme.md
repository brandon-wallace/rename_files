# Rename Files

## A command-line tool written in Go to clean up media filenames by removing unwanted characters and applying a consistent naming format.

### Features

- **Filename Normalization:** Converts filenames to lowercase and replaces spaces with underscores.
- **Emoji Removal:** Removes characters covered by the program's defined Unicode emoji ranges.
- **Bracket Cleanup:** Removes square-bracketed text, including the brackets.
- **Character Cleanup:** Removes commas and `#` characters, collapses repeated underscores, and trims underscores, hyphens, and periods from the start and end of the filename.
- **Recursive Scanning:** Searches the current directory and its subdirectories.
- **Dry Run:** Preview the proposed filenames without renaming files.
- **Duplicate Prevention:** Skips a rename if the target filename already exists.

### Prerequisites

- Go installed and available on your `PATH`.

## Installation

### Build

Save `rename_files.go` in your project directory, then build the executable:

```sh
go build -o rename_files rename_files.go
```

## Usage

Run the program from the directory containing the media files you want to process. It scans the current working directory (`.`) and all its subdirectories.

### Preview Changes

Use `-dry-run` to print the proposed filenames without changing any files:

```sh
./rename_files -dry-run
```

### Rename Files

After reviewing the preview, run the program without `-dry-run` to apply the changes:

```sh
./rename_files
```

## Flags

| Flag | Description | Default |
| --- | --- | --- |
| `-dry-run` | Preview proposed filenames without renaming files | `false` |

## Example

A filename such as:

```text
./Videos/My 🎬 Movie [1080p] - Sample.MP4
```

may be normalized to:

```text
./Videos/my_movie_sample.mp4
```

The exact result depends on the characters in the original filename and the cleanup rules implemented by the program.

When using `-dry-run`, the displayed count represents files that would be renamed, not files actually renamed.

## How It Works

1. **Scans:** Recursively walks the current directory and collects files with the supported extensions.
2. **Normalizes:** Removes recognized emoji characters, converts the filename to lowercase, removes square-bracketed text, replaces spaces with underscores, and removes selected punctuation.
3. **Preserves the extension:** Converts the file extension to lowercase.
4. **Checks for changes:** Skips files whose normalized path is unchanged.
5. **Prevents overwrites:** Skips a file if the proposed destination path already exists.
6. **Renames:** Renames the file unless dry-run mode is enabled.

## Notes

- Run the program from the directory tree you intend to process; it does not accept a source-directory flag.
- Files whose normalized target path already exists are skipped.
