import asyncio
import websockets
from base64 import b64decode, b64encode

ping = 'CJVOEOWIhuzUMBiSTiABOgJwYkKcAQjQDxCVThoFMC4zLjgiNklVaHl0QWRHNFNubGlNd09UTnU4ZkJKMDFaOEpJaVZXZDZQNzlBZmZYY2V3SXUyZXR2ZVNIMCgDMAA6DjEyYzkyOWE6bWFzdGVyQiSCfSEIhoKDnvWDydRjEAIaEzcxODEzMTA3OTc3ODY2OTM4OTRKEzcxNTc5NzA3Mjk0NzkzMDc1MjVaA3dlYpABAg=='


async def main():
    
    uri = "wss://oec-im-frontier-va.tiktokglobalshop.com/ws/v2?token=IUhytAdG4SnliMwOTNu8fBJ01Z8JIiVWd6P79AffXcewIu2etveSH0&aid=5341&fpid=92&device_id=7157970729479307525&access_key=5043e1721f1d2aa0ab2cdc3f41dc68d9&device_platform=web&version_code=10000&websocket_switch_region=ID"
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
    
    
    
    