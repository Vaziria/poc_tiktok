import asyncio
import traceback
import logging
from mitmproxy import options
from mitmproxy.tools import dump
from mitmproxy.http import HTTPFlow

from base64 import b64encode

logger = logging.getLogger(__name__)

class ResourceTampering:
    
    def load(self, loader):
        loader.add_option(
            name="validate_inbound_headers",
            typespec=bool,
            default=False,
            help="Add a count header to responses",
        )
    
    
    def response(context, flow: HTTPFlow):
        print(flow.request.url)
        
        if flow.request.url.find('goofy-va/pigeon-pc/chunk/vendors.b984d657.js') != -1:
            with open('./assets/vendors.b984d657.js', 'rb') as out:
                flow.response.content = out.read()
                
    
    # def websocket_message(context, flow: HTTPFlow):
    #     try:
    #         if flow.websocket:
    #             message = flow.websocket.messages[-1]
    #             print(message)
                
    #             with open('dump_websocket.txt', 'ab+') as out:
    #                 if message.from_client:
    #                     out.write('---->>>'.encode())
                    
    #                 else:
    #                     out.write('----<<<'.encode())
                        
    #                 out.write('\n'.encode())
                    
    #                 print(message.content, 'asdasdasd')
    #                 out.write(b64encode(message.content))
    #                 out.write('\n\n'.encode())
                    
    #     except Exception as e:
    #         print(e)
    #         traceback.print_exc()
            
    


async def start_proxy(host, port):
    opts = options.Options(listen_host=host, listen_port=port)

    master = dump.DumpMaster(
        opts,
        with_termlog=False,
        with_dumper=False,
    )
    
    master.addons.add(ResourceTampering())
    
    await master.run()
    return master


async def start_fix_proxy(host, port):
    opts = options.Options(listen_host=host, listen_port=port)

    master = dump.DumpMaster(
        opts,
        with_termlog=False,
        with_dumper=False,
    )
    
    master.addons.add(ResourceTampering())
    
    await master.run()

if __name__ == '__main__':
    asyncio.run(start_proxy('localhost', 6001))