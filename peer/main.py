from src.rest_client.apiclient import APIClient
from src.rest_client.messages import Messages
from src.rest_client.setup import SetUp
from dotenv import load_dotenv
import os
import sys
load_dotenv()
#from peer.src.rest_client.verification import Verify

def query_procedure():
    print("Please write the name of the file you are currently searching for:")
    file2search = input()
    querydata = {"filename": file2search}
    print(querydata)
    return querydata

def upload_procedure():
    print("Please write the name of the file that you want to upload:")
    file2search = input()
    querydata = {"filename": file2search}
    return querydata

def conversion_procedure():
    print("Please write the conversion type you want to do. Choose one of the following options:")
    print(" ")
    print(" COP2USD - Colombian peso to American dollar.")
    print(" USD2COP - American dollar to Colombian peso.")
    print(" COP2EUR - Colombian peso to Euro.")
    print(" EUR2COP - Euro to Colombian peso.")
    print(" COP2GBP - Colombian peso to Great Britain pound.")
    print(" GBP2COP - Great Britain pound to Colombian peso.")
    print(" COP2JPY - Colombian peso to Japanese Yang.")
    print(" JPY2COP - Japanese Yang to Colombian peso.")
    print(" COP2AUD - Colombian peso to Australian dollar.")
    print(" AUD2COP - Australian dollar to Colombian peso.")
    print(" ")

    conversion_type = input("Conversion type: ")
    print(" ")

    conversions = ["COP2USD", "USD2COP", "COP2EUR", "EUR2COP",  "COP2GBP", "GBP2COP",  "COP2JPY", "JPY2COP", "COP2AUD", "AUD2COP"]

    while conversion_type not in conversions:
        conversion_type = input("Write a valid conversion type: ")
        print(" ")

    print("Now, please write the amount you want to convert:")
    print(" ")
    while True:
        try:
            amount = float(input("Amount: "))
            print(" ")
            break
        
        except:
            print("Invalid amount. Write a valid number.")
            print(" ")
            
    querydata = {"conversion": conversion_type, "amount": amount}        
    return querydata
    

if __name__ == "__main__":

    client, messagec, setup = APIClient(), Messages(), SetUp()

    username = input("Write your name: ")
    password = input("Write a password: ")
    user_url = os.getenv("EXPOSED_IP_PSERVER")
    user_port = os.getenv("PSERVER_PORT")

    LogIn_data = {"username": username, "password": password, "user_url": user_url+":"+user_port}
    
    error, login_response = setup.logIn(client.url_servidor, LogIn_data)
    if error:
        print(login_response['message'])
        sys.exit()
    else:
        print("Logged in successfully!!")
        print("Assigned security token:",login_response['token'])
        authToken = login_response['token']


    
    
    
    list = SetUp.get_files_in_folder()
    sent_index_files = setup.do_sendIndex(client.url_servidor, list, authToken)
    print(sent_index_files['message'])
    print("All files were successfully added to the index!!")

    Messages.introduction()
    while True:
        Messages.general_message()
        action = input()
        
        if action == "1":
            list = SetUp.get_files_in_folder()
            sent_index_files = setup.do_sendIndex(client.url_servidor, list, authToken)
            querydata = query_procedure()
            error, query_response = client.do_query(client.url_servidor, querydata, authToken)
            if error:
                print(query_response['message'])
            else:
                print(query_response)


        
        elif action == "2":
            list = SetUp.get_files_in_folder()
            sent_index_files = setup.do_sendIndex(client.url_servidor, list, authToken)
            querydata = query_procedure()
            error, query_response = client.do_download(client.url_servidor, querydata, authToken)
            if error:
                print(query_response['message'])
            else:
                sent_index_files = setup.do_sendIndex(client.url_servidor, [querydata['filename']], authToken)
                print(query_response)
    
        elif action == "3":
            list = SetUp.get_files_in_folder()
            sent_index_files = setup.do_sendIndex(client.url_servidor, list, authToken)
            querydata = upload_procedure()
            error, query_response = client.do_upload(client.url_servidor, querydata, authToken)
            if error:
                print(query_response['message'])
            else: 
                print(query_response)


        elif action == "4":
            list = SetUp.get_files_in_folder()
            sent_index_files = setup.do_sendIndex(client.url_servidor, list, authToken)
            querydata = conversion_procedure()
            error, query_response = client.do_conversion(client.url_servidor, querydata, authToken)
            if error: 
                print(query_response)
            else: 
                print(query_response)


        elif action == "5":
            error, logout_response = client.logOut(client.url_servidor, authToken)
            if error:
                print(logout_response['message'])
            else:
                print(logout_response['message'])
            break