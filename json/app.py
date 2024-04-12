from flask import Flask,render_template,url_for
from flask_cors import CORS
from database import Database

app=Flask(__name__)
cors=CORS(app,resources={r"/":{"origin":"*"}})


@app.route("/")
def index():
    data=Database()
    data=data.fetch()
    return render_template("index.html",data=data)


if __name__=="__main__":
    app.run(debug=True)

