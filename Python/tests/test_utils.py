import os
import json

PROJ_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "../.."))


def load_test_data(filename):
    json_path = os.path.join(PROJ_PATH, "test_data", filename)
    with open(json_path, "r") as file:
        return json.load(file)
