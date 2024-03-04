import os
from src.protobuf_files.filesystem_pb2 import Response
from src.protobuf_files.filesystem_pb2_grpc import FileSystemServicer, FileSystem
from grpc import StatusCode

class FileSystemService(FileSystemServicer):
   def Upload(self, request, context):
      try:
          if request.name == "":
              print("Error, empty file")
              raise Exception("Empty file")
          response = Response(message=f"File '{request.name}' uploaded succesfully (ID: {request.id})")
          current_directory = os.path.dirname(os.path.abspath(__file__))
          files_directory = os.path.join(current_directory, '..','..', 'files')
          if not os.path.exists(files_directory):
              os.makedirs(files_directory)
          file_path = os.path.join(files_directory, request.name)
          file = open(file_path, 'w')
          file.close()


          return response
      except Exception as e:
          context.set_details(str(e))
          context.set_code(StatusCode.INVALID_ARGUMENT)
          return Response()
     
   def Download(self, request, context):
      try:
         if request.name == "":
               print("Error, empty file")
               raise Exception("Empty file")
         response = Response(message=f"File '{request.name}' downloaded succesfully (ID: {request.id})")
         return response
      except Exception as e:
          context.set_details(str(e))
          context.set_code(StatusCode.INVALID_ARGUMENT)
          return Response()