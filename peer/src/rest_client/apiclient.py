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
        self.pserver_ip = os.getenv("IP_PSERVER")
        self.pserver_port = os.getenv("PSERVER_PORT")
        self.dir = os.getenv("DIR")
        self.url_servidor = os.getenv("URL_SERVIDOR_CENTRAL")

    def do_query(self, url, querydata, authToken):
        specurl = url + "api/v1/query?file=" + querydata['filename']
        headers = {
            "Authorization": authToken
        }
        try:
            query_response = requests.get(specurl, headers=headers)
            query_response.raise_for_status()
            return False, query_response.json()
        
        
        except requests.exceptions.RequestException as e:
            return True, query_response.json()
        
    def do_download(self, url, querydata, authToken):
        error, query_response = self.do_query(url, querydata, authToken)
        if error:
            return query_response
        #specurl = url + "api/v1/query?file=" + querydata['filename']
        #headers = {
        #    "Authorization": authToken
        #}
        try:
            #query_response = requests.get(specurl, headers=headers)
            client_grpc = Client_Remote()
            
            dresponse = client_grpc.download(f"{query_response['location']}", querydata['filename'])
            return False, dresponse
        
        
        except requests.exceptions.RequestException as e:
            return True, query_response
        
    def do_upload(self, url, querydata, authToken):
        specurl = url + "api/v1/getPeerUploading"
        headers = {
            "Authorization": authToken
        }
        try:
            query_response = requests.get(specurl, headers=headers)
            query_response.raise_for_status()
            json_response = json.loads(query_response.text)
            print("json: ",json_response)
            client_grpc = Client_Remote()
            upresponse = client_grpc.upload(f"{json_response['location']}", querydata['filename'])
            return False, upresponse
        
        
        except requests.exceptions.RequestException as e:
            return True, query_response
        
    
    def do_conversion(self, url, querydata, authToken):
        specurl = url + "api/v1/getPeerUploading" 
        headers = {
            "Authorization": authToken
        }
        try:
            query_response = requests.get(specurl, headers=headers)
            query_response.raise_for_status()
            print("query:",query_response)
            json_response = json.loads(query_response.text)
            print("json: ",json_response)
            client_grpc = Client_Remote()
            c_response = client_grpc.currency_converter(f"{json_response['location']}", querydata['conversion'], querydata['amount'])
            return False, c_response
        
        
        except requests.exceptions.RequestException as e:
            return True, query_response




    def logOut(self, url, authToken):
        specurl = url + "api/v1/logout"
        headers = {
            "Authorization": authToken
        }
        try:
            logout_response = requests.post(specurl, headers=headers)
            logout_response.raise_for_status()
            return True, logout_response.json()
        except requests.exceptions.RequestException as e:
            return False, None