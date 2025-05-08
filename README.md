# TimeWarrior

A command line time tracking tool for developers and freelance workers who need to track time worked on their client and personal projects.

## Features

- Track time spent on tasks with project organization
- Start, pause, resume, and complete timeslips
- Switch between tasks while maintaining time tracking
- Generate detailed reports by project
- Simple and intuitive command-line interface
- Project-based organization for better time management

## Installation

### From Source

```bash
git clone https://github.com/mrcook/time_warrior.git
cd time_warrior
go build -o tw ./tw
```

Move the binary to a location in your PATH:

```bash
mv tw /usr/local/bin/
```

## Usage

### Basic Commands

```bash
# Show current status and available commands
tw

# Start a new timeslip
tw start [Project.Task]    # Start a task (uses current project if not specified)
tw start Project.Task      # Start a task with specific project

# Pause the current timeslip
tw pause                   # Pause current task
tw p                      # Alias for pause

# Resume a paused timeslip
tw resume                  # Resume paused task

# Complete the current timeslip
tw done                    # Mark current task as completed

# Switch to a new task
tw switch [Task]           # Stop current task and start new one (uses current project)
tw switch Project.Task     # Switch to a task in a different project

# Delete the current timeslip
tw delete                  # Delete current task

# Adjust time on the current timeslip
tw adjust +30m             # Add 30 minutes
tw adjust -15m             # Subtract 15 minutes
```

### Project Management

```bash
# Set or show current project
tw project                 # Show current project
tw project ProjectName     # Set current project
tw pr                      # Alias for project command

# Generate a report
tw report                  # Show report for all projects
tw report ProjectName      # Show report for specific project
```

### Time Format

Time adjustments can be specified in various formats:
- `+30m` or `-30m` for minutes
- `+1h` or `-1h` for hours
- `+1h30m` for combined hours and minutes

## Examples

### Basic Workflow

```bash
# Set your current project
tw project alluvial

# Start working on a task
tw start coding

# Pause for a break
tw pause

# Resume work
tw resume

# Switch to a different task
tw switch testing

# Complete the task
tw done
```

### Project-Specific Workflow

```bash
# Work on multiple projects
tw project client1
tw start feature1
tw pause
tw project client2
tw start bugfix
tw done
tw project client1
tw resume
```

### Time Adjustment

```bash
# Add time to current task
tw adjust +30m

# Subtract time from current task
tw adjust -15m

# Add multiple hours and minutes
tw adjust +2h30m
```

### Reporting

```bash
# View all project reports
tw report

# View specific project report
tw report alluvial
```

## Configuration

TimeWarrior stores its data in the following locations:
- Data directory: `~/.time_warrior/`
- Pending timeslip: `~/.time_warrior/pending`
- Project setting: `~/.time_warrior/.project`

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
