from src.server import server
from dotenv import load_dotenv
import os
load_dotenv()

def init_server():
    pserver_ip = os.getenv("IP_PSERVER")
    pserver_port = os.getenv("PSERVER_PORT")
    server_test= server.Server(pserver_ip,pserver_port)
    server_test.start_server()
    server_test.wait_for_termination()

if __name__ == '__main__':
    init_server()  
