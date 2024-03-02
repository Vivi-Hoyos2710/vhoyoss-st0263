import os
import cmd
import requests
from dotenv import load_dotenv
import sys
import os
import json

# Get the current directory of the script
current_dir = os.path.dirname(os.path.abspath(__file__))

# Get the parent directory of the current directory (i.e., the root directory of your project)
parent_dir = os.path.dirname(current_dir)

# Add the parent directory to the Python path
sys.path.append(parent_dir)

# Now you should be able to import modules from both grpc_client and rest_client directories
from grpc_client.client import Client_Remote



load_dotenv()

class APIClient(cmd.Cmd):

    def __init__(self):
        self.ipl = os.getenv("IP_Talking")
        self.port = os.getenv("PORT")
        self.port2 = os.getenv("PORT_2")
        self.dir = os.getenv("DIR")
        self.url_servidor = os.getenv("URL_SERVIDOR_CENTRAL")

    def do_query(self, url, querydata, authToken):
        specurl = url + "api/v1/query?file=" + querydata['filename']
        headers = {
            "authToken": authToken
        }
        try:
            query_response = requests.get(specurl, headers=headers)
            query_response.raise_for_status()
            print("query:",query_response)
            #client_grpc = Client_Remote()
            #client_grpc.download(f"{self.ipl}:{self.port2}", querydata['filename'])
            return query_response.json()
        
        
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
            return None
        
    def do_download(self, url, querydata, authToken):
        specurl = url + "api/v1/query?file=" + querydata['filename']
        headers = {
            "authToken": authToken
        }
        try:
            query_response = requests.get(specurl, headers=headers)
            query_response.raise_for_status()
            #print(query_response)
            client_grpc = Client_Remote()
            dresponse = client_grpc.download(f"{self.ipl}:{self.port2}", querydata['filename'])
            return dresponse
        
        
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
            return None
        
    def do_upload(self, url, querydata, authToken):
        specurl = url + "api/v1/getPeerUploading?filename=" + querydata['filename'] + "&user=" + querydata['username']
        headers = {
            "authToken": authToken
        }
        try:
            query_response = requests.get(specurl, headers=headers)
            query_response.raise_for_status()
            print("query:",query_response)
            client_grpc = Client_Remote()
            client_grpc.upload(f"{self.ipl}:{self.port2}", querydata['filename'])
            return query_response.json()
        
        
        except requests.exceptions.RequestException as e:
            jsonstring = json.loads(query_response.text)
            print(f"Error making request: {jsonstring['message']}")
            return None

    def logOut(self, url, logout_data, authToken):
        specurl = url + "api/v1/logout"
        headers = {
            "authToken": authToken
        }
        try:
            logout_response = requests.post(specurl, json=logout_data, headers=headers)
            logout_response.raise_for_status()
            return logout_response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
            return None