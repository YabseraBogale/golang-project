import time
from pprint import pprint
import pandas as pd


def Data(data):
    year=data["vulnerabilities"][1999]['cve']['lastModified'].split('T')[0].split('-')[0]
    vulnerabilities=data["vulnerabilities"][199]['cve']['descriptions'][0]['value']
    attackComplexity=data["vulnerabilities"][199]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['attackComplexity']
    return {"year":year,"vulnerabilities":vulnerabilities,'attackComplexity':attackComplexity}
high = pd.read_json("high.json")

pprint(high["vulnerabilities"][199]['cve'])