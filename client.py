# localhost:8080にHTTPリクエストを送る
import requests

from proto import message_pb2

url = 'http://localhost:8080/api'

response = requests.get(url)

message_pb2 = message_pb2.Person()
print(type(message_pb2))

messsage = message_pb2.ParseFromString(response.content)
print(message_pb2)
print(messsage)

