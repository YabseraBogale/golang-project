from flask import Flask
from pprint import pprint
import pandas as pd

high = pd.read_json("high.json")
critcal=pd.read_json("critcal.json")
windows=pd.read_json("windows.json")
app =Flask(__name__)

#print(high["vulnerabilities"][1999]['cve']['descriptions'][0]['value'])
#pprint(high["vulnerabilities"][1999]['cve'])


@app.route("/")
def index():
    return windows["vulnerabilities"][199]['cve']['descriptions'][0]['value']


if __name__=="__main__":
    app.run(debug=True)

