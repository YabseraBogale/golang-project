from flask import Flask,render_template,url_for
from database import Database
app=Flask(__name__)
data=Database()
@app.route("/")
def index():
    data=data.fetch()
    return render_template("index.html",data=data)


if __name__=="__main__":
    app.run(debug=True)

