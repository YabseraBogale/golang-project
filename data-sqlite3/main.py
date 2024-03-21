import sqlite3
from pprint import pprint
connection=sqlite3.Connection("freelance-data-v2.db")

pointer=connection.cursor()

pointer.execute("select message from Software where id=34400")
result=pointer.fetchall()
print(result[0][0])
