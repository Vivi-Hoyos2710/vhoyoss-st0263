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
            "authToken": authToken
        }
        try:
            query_response = requests.get(specurl, headers=headers)
            query_response.raise_for_status()
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
            json_response = json.loads(query_response.text)
            print("json: ",json_response)
            print("query:", querydata)
            client_grpc = Client_Remote()
            
            dresponse = client_grpc.download(f"{json_response['location']}", querydata['filename'])
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
            json_response = json.loads(query_response.text)
            print("json: ",json_response)
            client_grpc = Client_Remote()
            client_grpc.upload(f"{json_response['message']}", querydata['filename'])
            return query_response.json()
        
        
        except requests.exceptions.RequestException as e:
            jsonstring = json.loads(query_response.text)
            print(f"Error making request: {jsonstring['message']}")
            return None
        
    
    def do_conversion(self, url, querydata, authToken):
        specurl = url + "api/v1/getPeerUploading?filename=" + querydata['filename'] + "&user=" + querydata['username'] ### Acá es donde debo definir la URL
        headers = {
            "authToken": authToken
        }
        try:
            query_response = requests.get(specurl, headers=headers)
            query_response.raise_for_status()
            print("query:",query_response)
            json_response = json.loads(query_response.text)
            print("json: ",json_response)
            client_grpc = Client_Remote()
            c_response = client_grpc.currency_converter(f"{json_response['message']}", querydata['conversion'], querydata['amount'])
            return c_response
        
        
        except requests.exceptions.RequestException as e:
            jsonstring = json.loads(query_response.text)
            print(f"Error making request: {jsonstring['message']}")
            return None

    ### -----------------------------------------------------------------------------------------|
    #### Acá debo |(1)| crear las funciones para cada servicio donde organizado toda la URL,     |
    #### |(2)| envío la data necesaria y |(3)| hago la creación del cliente gRPC, del cual debo  |
    #### |(4)| llamar la función del servicio correspondiente y |(5)| recibir la respuesta.      |
    ### -----------------------------------------------------------------------------------------|




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