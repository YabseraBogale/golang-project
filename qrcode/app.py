from flask import Flask,url_for

app=FLask(__name__)


@spp.route('/')
def index():
    return "hello world"

if __name__=="__main__":
    app.run(debug=True)
