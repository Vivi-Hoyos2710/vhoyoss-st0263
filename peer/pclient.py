import sys
from src.rest_client.apiclient import APIClient
from src.rest_client.messages import Messages
from src.rest_client.setup import SetUp
#from peer.src.rest_client.verification import Verify

def query_procedure():
    print("Please write the name of the file you are currently searching for:")
    file2search = input()
    querydata = {"filename": file2search}
    return querydata

def upload_procedure(username):
    print("Please write the name of the file you are currently searching for:")
    file2search = input()
    querydata = {"filename": file2search, "username":username}
    return querydata

def conversion_procedure():
    print("Please write the conversion type you want to do. Choose one of the following options:")
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

    print("Now, please write the amount you want to convert:")
    print(" ")

    amount = float(input("Amount: "))

    querydata = {"conversion": conversion_type, "amount": amount, "filename":"gatitos2.txt", "username":"Vivi"}
    return querydata
    

if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("Please run the client using the following format:")
        print("python main.py <Username> <Password> <User url>")
        sys.exit()

    client, messagec, setup = APIClient(), Messages(), SetUp()

    username,password,user_url = sys.argv[1],sys.argv[2],sys.argv[3]
    LogIn_data = {"username": username, "password": password, "user_url": user_url}
    
    login_response = setup.logIn(client.url_servidor, LogIn_data)
    
    print("Logged in successfully!!")
    print("Assigned security token:",login_response['token'])
    authToken = login_response['token']

    list = SetUp.get_files_in_folder()
    #print(list)
    sent_index_files = setup.do_sendIndex(client.url_servidor, list, username, authToken)
    print(sent_index_files['message'])
    print("All files were successfully added to the index!!")

    Messages.introduction()
    while True:
        Messages.general_message()
        action = input()
        
        if action == "1":
            querydata = query_procedure()
            query_response = client.do_query(client.url_servidor, querydata, authToken)
            print("Localtion:",query_response['location'])

        
        elif action == "2":
            querydata = query_procedure()
            query_response = client.do_download(client.url_servidor, querydata, authToken)
            print(query_response)

    
        elif action == "3":
            querydata = upload_procedure(username)
            query_response = client.do_upload(client.url_servidor, querydata, authToken)
            print(query_response)


        elif action == "4":
            querydata = conversion_procedure()
            query_response = client.do_conversion(client.url_servidor, querydata, authToken)
            print(query_response)


        elif action == "5":
            logoutdata = {"username": username}
            logout_response = client.logOut(client.url_servidor, logoutdata, authToken)
            print(logout_response['message'])
            break

        ### -------------------------------------------------------------------------------
        #### Debo (1) crear las dos opciones para los servicios adicionales, (2) crear las
        #### funciones para procesar la info necesaria. (3) Hacer llamado y enviar info a
        #### las funciones del cliente gRPC correspondientes. 
        ### -------------------------------------------------------------------------------