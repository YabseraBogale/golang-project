import sqlite3

connection=sqlite3.Connection("freelance-data-v2.db")

stakmap=["django", "flask", "fastapi", "javascript", "php", "wordpress", "java ", "java,", "spring boot",
		"spring", "node", "front end", "frontend", "backend", "back end", "fullstack", "react", "vue", "c#", ".net", "C++", "dotnet",
		"asp.net", "python", "andriod", "ios", "mobile", "mysql", "mongodb", "postgres", "flutter", "dart", "angularjs"]

pointer=connection.cursor()

pointer.execute("select message,id from Software")
result=pointer.fetchall()

d={}

for i in result:
    lst=[]
    count=0
    for j in stakmap:
        if str(i[0]).lower().find(j)>0:
            lst.append(j)
            count+=1
    if count>=2:
        d[i[1]]=lst
                

print(d)
print(len(d.keys()))

