import asyncio
import websockets
from base64 import b64decode, b64encode

from src.tpayload import Request

ping = Request()
ping.SerializeToString()



ping = 'CJJOEOavpejVMBiSTiABOgJwYkKGAQjIARCSThoFMC4zLjgiNkRwOHdjc1BhTjlud1FnU1diTUg2elY1NFkzZUZsRWo5NHBSenhybXBVd0JWbEJ6VkZuMXRoNCgDMAA6DjEyYzkyOWE6bWFzdGVyQg7CDAsI4dKv8tCd/AIQMkoTNzE1Nzk3MDcyOTQ3OTMwNzUyNVoDd2VikAEC'


async def main():
    
    uri = "wss://oec-im-frontier-va.tiktokglobalshop.com/ws/v2?token=Dp8wcsPaN9nwQgSWbMH6zV54Y3eFlEj94pRzxrmpUwBVlBzVFn1th4&aid=5341&fpid=92&device_id=7157970729479307525&access_key=5043e1721f1d2aa0ab2cdc3f41dc68d9&device_platform=web&version_code=10000&websocket_switch_region=ID"
    async with websockets.connect(uri) as websocket:
        
        async def ping_loop():
            while True:
                await asyncio.sleep(10)
                await websocket.send(b64decode(ping))
            
        ping_task = asyncio.create_task(ping_loop())
        
        
        while True:
            
            
            pong = await websocket.recv()
            print(b64encode(pong))
            print('-----------------------------------------')
        
        
        await ping_task
        
        
if __name__ == "__main__":
    asyncio.run(main())
    
    
    
     