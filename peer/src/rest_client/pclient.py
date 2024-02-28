import os
import cmd
import requests
from dotenv import load_dotenv
import sys
import json


# Load environment variables from .env file
load_dotenv()

class APIClient(cmd.Cmd):

    def __init__(self):
        self.ipl = os.getenv("IP_Listening")
        self.port = os.getenv("PORT")
        self.dir = os.getenv("DIR")
        self.url_servidor = os.getenv("URL_SERVIDOR_CENTRAL")
        self.help_messages = self.load_help_messages()


    def load_help_messages(self):
        with open('help.txt', 'r') as file:
            content = file.read()
            return content
            
    def logIn(self, url, LogIn_data):
        specurl = url + "api/v1/login"
        try:
            login_response = requests.post(specurl, json=LogIn_data)
            login_response.raise_for_status()
            return login_response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
            return None


    def do_sendIndex(self, url, sendIndex_data, authToken):
        specurl = url + "api/v1/sendIndex"
        headers = {
            "Content-Type": "application/json",
            "authToken": authToken
        }
        
        try:
            sendIndex_response = requests.post(specurl, json=sendIndex_data, headers=headers)
            sendIndex_response.raise_for_status()
            return sendIndex_response.json()
        
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
            return None

    
    def do_query(self, url, querydata, authToken):
        specurl = url+"api/v1/query?file="+querydata['filename']
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
            logout_response = requests.post(specurl, json=logout_data)
            logout_response.raise_for_status()
            return logout_response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error making request: {e}")
            return None
    

    def do_quit(self, line):
        return True

    def postcmd(self, stop, line):
        print() 
        return stop

if __name__ == "__main__":

    # APIClient().cmdloop()
    if len(sys.argv) != 4:
        print("Please run the client using the following format:")
        print("python pclient.py <Username> <Password> <User url>")
        sys.exit()
    username = sys.argv[1]
    password = sys.argv[2]
    user_url = sys.argv[3]
    
    LogIn_data = {"username": username,"password": password,"user_url": user_url}
    
    client = APIClient()
    login_response = client.logIn(client.url_servidor, LogIn_data)
    print("Logged in successfully!!")
    print(login_response)
    authToken = login_response['token']

    helpmes = client.load_help_messages()
    print(helpmes)
    
    
    indexdata = {"username": sys.argv[1],"files":['HarryPotter1234.txt']}
    sendIndex_response = client.do_sendIndex(client.url_servidor, indexdata, authToken)
    print(sendIndex_response)

    querydata = {"filename":'HarryPotter.txt'}
    query_response = client.do_query(client.url_servidor, querydata, authToken)
    print(query_response)

    logoutdata = {"username": username}
    logout_response = client.logOut(client.url_servidor, logoutdata, authToken)
    print(logout_response)