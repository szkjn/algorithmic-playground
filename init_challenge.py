"""
This script streamlines challenge initialization by creating directories and files for 
the three main languages used in this playground. 

Given a filename, it sets up :
- a Go environment with main.go and go.mod in Go/<filename>/
- a Python script in Python/scripts/
- a TypeScript file in TypeScript/ (converting the filename to camelCase)
"""

import os
import sys


def create_directory(path):
    if not os.path.exists(path):
        os.makedirs(path)


def create_file(path, content=""):
    with open(path, "w") as file:
        file.write(content)


def to_camel_case(snake_str):
    components = snake_str.split("_")
    return components[0] + "".join(x.title() for x in components[1:])


def main():
    if len(sys.argv) != 2:
        print("Usage: python init_challenge.py <filename>")
        sys.exit(1)

    filename = sys.argv[1]
    go_folder = os.path.join("Go", filename)
    python_folder = os.path.join("Python", "scripts")
    typescript_folder = os.path.join("TypeScript", "src")

    # Create Go folder and files
    create_directory(go_folder)
    create_file(os.path.join(go_folder, "main.go"))
    create_file(
        os.path.join(go_folder, "go.mod"),
        f"module github.com/szkjn/algorithmic-playground/Go/{filename}\n\ngo 1.21.4\n",
    )

    # Create Python file
    create_file(os.path.join(python_folder, f"{filename}.py"))

    # Create TypeScript file
    typescript_filename = to_camel_case(filename)
    create_file(os.path.join(typescript_folder, f"{typescript_filename}.ts"))

    print("Initialization complete.")


if __name__ == "__main__":
    main()
