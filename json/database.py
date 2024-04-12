from random import randint
import sqlite3
from pprint import pprint
"""
on old hardware
python for 100

real    0m0.078s
user    0m0.075s
sys     0m0.004s

golang for 100

real    0m0.467s
user    0m0.417s
sys     0m0.140s


python for 1000

real    0m0.239s
user    0m0.195s
sys     0m0.044s

golang for 1000

real    0m0.866s
user    0m0.662s
sys     0m0.250s

on new hard in the last 10 years

python3 for 100
real    0m0.342s
user    0m0.202s
sys     0m0.051s

"""
class Database():
    def __init__(self) -> None:
        self.connnection=sqlite3.Connection("data.db")
        self.pointer=self.connnection.cursor()

    def fetch(self):
        statment=f"select * from data where id={randint(0,4000)}"
        self.pointer.execute(statment)
        result=self.pointer.fetchone()
        return result
    
test=Database()
for i in range(0,100):
    pprint(test.fetch())
