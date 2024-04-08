from flask import Flask
from pprint import pprint
import pandas as pd

high = pd.read_json("high.json")
app =Flask(__name__)

#print(high["vulnerabilities"][1999]['cve']['descriptions'][0]['value'])
#pprint(high["vulnerabilities"][1999]['cve'])


@app.route("/")
def index():
    return high["vulnerabilities"][1999]['cve']['descriptions'][0]['value']


if __name__=="__main__":
    app.run(debug=True)

