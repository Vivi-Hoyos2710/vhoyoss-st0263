import os
import cmd
import requests
from dotenv import load_dotenv

load_dotenv()

class APIClient(cmd.Cmd):

    def __init__(self):
        self.ipl = os.getenv("IP_Listening")
        self.port = os.getenv("PORT")
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
            return query_response.json()
        
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
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