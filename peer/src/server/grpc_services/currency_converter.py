
from src.protobuf_files.filesystem_pb2 import Response_2
from src.protobuf_files.filesystem_pb2_grpc import CurrencyConverterServicer, FileSystem
from grpc import StatusCode

class CurrencyConverterService(CurrencyConverterServicer):
    def Convert(self, request, context):
        try:
            if request.conversion == "":
                print("Error, no conversion")
                raise Exception("No conversion")
            elif request.amount == "":
                print("Error, no amount")
                raise Exception("No amount")
          
            if request.conversion == "USD2COP":
                result = request.amount * 3900
            elif request.conversion == "COP2USD":
                result = request.amount / 3900
            elif request.conversion == "EUR2COP":
                result = request.amount * 4247
            elif request.conversion == "COP2EUR":
                result = request.amount / 4247
            elif request.conversion == "GBP2COP":
                result = request.amount * 4952
            elif request.conversion == "COP2GBP":
                result = request.amount / 4952
            elif request.conversion == "JPY2COP":
                result = request.amount * 26.07
            elif request.conversion == "COP2JPY":
                result = request.amount / 26.07
            elif request.conversion == "AUD2COP":
                result = request.amount * 2557
            elif request.conversion == "COP2AUD":
                result = request.amount / 2557
            else:
                raise Exception("Invalid conversion type")
            response = Response_2(amount_result=round(result, 2))
            return response
        except Exception as e:
            context.set_details(str(e))
            context.set_code(StatusCode.INVALID_ARGUMENT)
            return Response_2()