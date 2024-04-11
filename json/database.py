from random import randint
import sqlite3
from pprint import pprint

class Database():
    
    def __init__(self) -> None:
        self.connection=sqlite3.Connection("vulnerabilities.db")
        self.pointer=self.connection.cursor()
    
    def fetch(self):
        statment="select * from vulnerabilities"
        self.pointer.execute(statment)
        result=self.pointer.fetchall()
        return result[randint(0,len(result))]
    


test=Database()
pprint(test.fetch())