import time
from pprint import pprint
import pandas as pd


def Data(data):
    year=data["vulnerabilities"][1999]['cve']['lastModified'].split('T')[0].split('-')[0]
    vulnerabilities=data["vulnerabilities"][199]['cve']['descriptions'][0]['value']
    attackComplexity=data["vulnerabilities"][199]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['attackComplexity']
    baseScore=data["vulnerabilities"][199]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['baseScore']
    return {"year":year,"vulnerabilities":vulnerabilities,'attackComplexity':attackComplexity,"score":baseScore}

## maxlength of high.json descriptions is 308
## maxlength of critcal.json descriptions is 255
## maxlength of windows.json descriptions is 332

high= pd.read_json("windows.json")


maxLength=308
for i in range(0,2000):
    if len(high["vulnerabilities"][i]['cve']['descriptions'][0]['value'])>maxLength:
        maxLength=len(high["vulnerabilities"][199]['cve']['descriptions'][0]['value'])

print(maxLength)