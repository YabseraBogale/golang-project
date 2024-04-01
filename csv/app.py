from requests import get
from pprint import pprint

url="https://raw.githubusercontent.com/YabseraBogale/cvelistV5/main/cves/2015/0xxx/CVE-2015-0018.json"

data=get(url).json()

pprint(data)