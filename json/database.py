import sqlite3

class Database():
    
    def __init__(self) -> None:
        self.connection=sqlite3.Connection("vulnerabilities.db")
        self.pointer=self.connection.cursor()
    
    def fetch(self):
        statment="select * from vulnerabilities"
        self.pointer.execute(statment)
    
    