from random import randint
from pprint import pprint
import pandas as pd

import sqlite3

def Data(data):
    year=data["vulnerabilities"][1999]['cve']['lastModified'].split('T')[0].split('-')[0]
    vulnerabilities=data["vulnerabilities"][199]['cve']['descriptions'][0]['value']
    attackComplexity=data["vulnerabilities"][199]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['attackComplexity']
    baseScore=data["vulnerabilities"][199]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['baseScore']
    return {"year":year,"vulnerabilities":vulnerabilities,'attackComplexity':attackComplexity,"score":baseScore}

    
def makeDabase():
    statment="""
        create table if not exists vulnerabilities(
            id int not null primary key,
            vulnerability varchar(332) not null,
            year varchar(4) not null,
            attackComplexity varchar(10) not null,
            baseScore varchar(10) not null
        );
    """
    connection=sqlite3.Connection("vulnerabilities.db")
    pointer=connection.cursor()
    pointer.execute(statment)




## maxlength of high.json descriptions is 308
## maxlength of critcal.json descriptions is 255
## maxlength of windows.json descriptions is 332

high= pd.read_json("high.json")


