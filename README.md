# Memo

Memo is a simple todo list CLI program built to help you stay organized and on top of your tasks.

![Memo Main Menu](https://imgur.com/uWhciQX)

## Features

- **Task Management:** Easily add, delete, and update tasks from the command line.
- **Clean Interface:** Minimalist design for efficient task management without distractions.
- **Cross-Platform:** Works on any platform with Go support, including Windows, macOS, and Linux.

## Usage

### Using Pre-compiled Builds
1. Clone the repository:

```bash
git clone https://github.com/nelisa-dludla/memo.git
```

2. Navigate to the project directory:

```bash
cd memo
```

3. Choose the appropriate pre-compiled executable for your operating system:
- For Windows: `memo.exe`
- For Linux: `memo`

4. Run the executable:

```bash
./memo # On Linux
```

```bash
.\memo.exe # On Windows
```

### Building and Running from Source

1. Clone the repository:

```bash
git clone https://github.com/nelisa-dludla/memo.git
```

2. Navigate to the project directory:

```bash
cd memo
```

3. Build the project:

```bash
go build
```

4. Run the executable:

```bash
./memo
```

5. View Tasks:

Upon launching Memo, you'll be greeted with a list of your tasks, along with their due dates.

![Memo Main Menu](https://imgur.com/b1aFGcj)

6. Interact with Tasks:

- `Add Task`: Add a new task to your list.
- `Edit Task`: Modify the details of an existing task.
- `Delete Task`: Remove a task from your list.
- `Mark Task as Completed`: Mark a task as completed.

## Database

The project uses SQLite for database storage. The database file (memo.db) will be created in the project directory upon first run.

