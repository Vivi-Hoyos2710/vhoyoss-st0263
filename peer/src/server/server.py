from concurrent import futures
import grpc
from src.protobuf_files.filesystem_pb2_grpc import add_FileSystemServicer_to_server
from src.protobuf_files.filesystem_pb2_grpc import add_CurrencyConverterServicer_to_server
from src.server.grpc_services.file_system import FileSystemService
from src.server.grpc_services.currency_converter import CurrencyConverterService

class Server:
    def __init__(self, host: str,port: str) -> None:
        self.host = host
        self.port = port
        
    def start_server(self)->None:
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        #add service to server
        add_FileSystemServicer_to_server(FileSystemService(), self.server)
        add_CurrencyConverterServicer_to_server(CurrencyConverterService(), self.server)

        self.server.add_insecure_port(f"{self.host}:{self.port}")
        self.server.start()
        print(f"Server started on host {self.host} and port {self.port}", flush=True)

    def wait_for_termination(self) -> None:
        if self.server:
            self.server.wait_for_termination()
            print(f"Server stopped", flush=True)
        else:
            print("Server not started", flush=True)
