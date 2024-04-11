import sqlite3

class Database():
    def __init__(self) -> None:
        self.connection=sqlite3.Connection("vulnerabilities.db")
        self.pointer=self.connection.cursor()
    
    def fetch(self):
        statment="select * from vulnerabilities"
        self.pointer.execute(statment)
    
    def changeType(self):
        statment="ALTER TABLE table_name MODIFY COLUMN baseScore float;"
        try:
            self.pointer.execute(statment)
            self.connection.commit()
            return "changed"
        except Exception as e:
            return e
        

test=Database()
print(test.changeType())