from flask import Flask

from requests import get


app=Flask(__name__)





@app.route('')
def index():
    url="https://raw.githubusercontent.com/YabseraBogale/cvelistV5/main/cves/2015/0xxx/CVE-2015-0018.json"
    data=get(url).json()
    return f"{data}"

if __name__=="__main__":
    app.run(debug=True)