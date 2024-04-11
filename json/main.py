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
   
    
def makeDabase():
    try:
        statment="""
            create table Data(
            id int not null primary key,
            vulnerability varchar(308) not null,
            year int not null,
            attackComplexity varchar(4) not null,
            baseScore int not null
        )
        """
        connection=sqlite3.Connection("data.db")
        pointer=connection.cursor()
        pointer.execute(statment)
        connection.commit()
        return "Done"    
    except Exception as e:
        return e

def Insert(Data,i):
    try:
        statment="insert into Data(id,vulnerability,year,attackComplexity,baseScore) values(?,?,?,?,?)"
        connection=sqlite3.Connection("data.db")
        pointer=connection.cursor()
        pointer.execute(statment,i,Data["vulnerabilities"],Data["year"],Data["attackComplexity"],Data["score"])
        connection.commit()
        return "Done"
    except Exception as e:
        return e
      
## maxlength of high.json descriptions is 308
## maxlength of critcal.json descriptions is 255

print(makeDabase())

