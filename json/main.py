import time
from pprint import pprint
import pandas as pd


def Data(data):
    year=data["vulnerabilities"][1999]['cve']['lastModified'].split('T')[0].split('-')[0]
    vulnerabilities=data["vulnerabilities"][199]['cve']['descriptions'][0]['value']
    attackComplexity=data["vulnerabilities"][199]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['attackComplexity']
    baseScore=data["vulnerabilities"][199]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['baseScore']
    return {"year":year,"vulnerabilities":vulnerabilities,'attackComplexity':attackComplexity,"score":baseScore}
high = pd.read_json("high.json")


maxLength=0
for i in range(0,2000):
    if len(high["vulnerabilities"][i]['cve']['descriptions'][0]['value'])>maxLength:
        maxLength=len(high["vulnerabilities"][199]['cve']['descriptions'][0]['value'])

print(maxLength)