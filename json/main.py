from random import randint
from pprint import pprint
import pandas as pd

import sqlite3

def Data(name,i):
      
    data= pd.read_json(f"{name}.json")
    year=data["vulnerabilities"][i]['cve']['lastModified'].split('T')[0].split('-')[0]
    vulnerabilities=data["vulnerabilities"][i]['cve']['descriptions'][0]['value']
    if str(data["vulnerabilities"][i]['cve']['metrics']).find("cvssMetricV31")<0:
        attackComplexity=data["vulnerabilities"][i]['cve']['metrics']['cvssMetricV30'][0]['cvssData']['attackComplexity']
        baseScore=data["vulnerabilities"][i]['cve']['metrics']['cvssMetricV30'][0]['cvssData']['baseScore']
    else:
        attackComplexity=data["vulnerabilities"][i]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['attackComplexity']
        baseScore=data["vulnerabilities"][i]['cve']['metrics']['cvssMetricV31'][0]['cvssData']['baseScore']
    return {"year":year,"vulnerabilities":vulnerabilities,'attackComplexity':attackComplexity,"score":baseScore}
   
    
def makeDabase(data):
    try:
        statment="""
            create table if not exists vulnerabilities(
                vulnerability varchar(332) not null,
                year int not null,
                attackcomplexity varchar(10) not null,
                baseScore varchar(10) not null
            );
        """
        connection=sqlite3.Connection("vulnerabilities.db")
        pointer=connection.cursor()
        pointer.execute(statment)
        connection.commit()
        statment="insert into vulnerabilities(vulnerability,year,attackcomplexity,baseScore) values(?,?,?,?)"
        pointer.execute(statment,(data["vulnerabilities"],data["year"],data["attackComplexity"],data["score"],))
        connection.commit()
        return i
    except Exception as e:
        return e


## maxlength of high.json descriptions is 308
## maxlength of critcal.json descriptions is 255


