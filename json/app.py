from flask import Flask,render_template,url_for
from flask_cors import cross_origin
from database import Database

app=Flask(__name__)



@app.route("/power")
@cross_origin()
def index():
    data=Database()
    data=data.fetch()
    return render_template("index.html",data=data)


if __name__=="__main__":
    app.run(debug=True)

