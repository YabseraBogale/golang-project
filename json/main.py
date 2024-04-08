import time
from pprint import pprint
import pandas as pd


high = pd.read_json("high.json")

pprint(high["vulnerabilities"][1999]['cve']['lastModified'].split('T')[0].split('-')[0])