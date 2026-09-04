# toGo - CLI for Task Management

![Gopher](gopher.png)

toGo is a simple and convenient CLI tool for managing notes and tasks. With it, you can save and delete notes and tasks, making it an ideal assistant for organizing your time.

## Features
- Add, delete, and view tasks
- Simple command-line interface

## Prerequisites
Make sure you have the following installed:
- `git`
- `bash`

## Installation

To install toGo, simply download the repository and run the `build.sh` script:

```bash
git clone https://github.com/gh0stErr0r/toGo.git
cd toGo
./build.sh
```

## Usage

Here are some examples of commands you can use with toGo:

- **Add a task:** 
  ```bash
  togo add "Buy milk"
  ```
- **Delete a task:** 
  ```bash
  togo del 1
  ```
- **View all tasks:** 
  ```bash
  togo list
  ```

## Error Handling
If you encounter issues with the database, you can try running:
```bash
sudo togo [command]
```
(Note: This command is intended to display the logo.)
