from src.server import server
from dotenv import load_dotenv
import os
load_dotenv()

if __name__ == "__main__":
    ipl = os.getenv("IP_Listening")
    port = os.getenv("PORT_2")
    server_test= server.Server(ipl,port)
    server_test.start_server()
    server_test.wait_for_termination()
    
