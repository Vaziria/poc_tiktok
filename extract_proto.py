import re
from typing import List

regex = re.compile(r'for \(.*?new c\.im_proto\.([a-zA-Z0-9]+).*?\{((.|\n)*?)\}')
refield = re.compile(r'case ([0-9]+)\:((.|\n)*?)break\;')
refieldname = re.compile(r'[a-z]{1}\.([a-zA-Z0-9\_]+)\s\=')
retype = re.compile(r'[a-z]{1}\.[a-zA-Z0-9\_]+\s\=\s[a-zA-Z0-9]\.([a-zA-Z0-9]+)\(\)\;')
reobjtype = re.compile(r'[a-z]{1}\.[a-zA-Z0-9\_]+\s\=\s[a-zA-Z]{1}\.im\_proto\.([a-zA-Z]+)\.decode')

rerepeatname = re.compile(r'[a-z]{1}\.([a-zA-Z0-9\_]+)\.push\((.*?)\);')
reempty = re.compile(r'[a-z]{1}\.([a-zA-Z0-9\_]+)\s===\s[a-z]{1}\.emptyObject\s\&\&')


with open('./assets/vendors.b984d657.js', 'r') as out:
    
    data = out.read()
    
    hasil = regex.findall(data)




def extract_field(text):
    
    fields = refield.findall(text)
    
    field_format = "{tipe} {key} = {index};"
    
    for field in fields:
        repeated = False
        
        key = refieldname.findall(field[1])
        tipe = retype.findall(field[1])
        
        if len(key) == 0 or len(tipe) == 0:
            
            objtype = reobjtype.findall(field[1])
            
            if len(objtype) == 0:
                
                repeatnames = rerepeatname.findall(field[1])
                
                if len(repeatnames) == 0:
                    
                    continue
                
                repeatnames = repeatnames[0]
                
                tipes = repeatnames[1] 
                if tipes.find('im_proto') != -1:
                    tipe = [tipes.split('.')[2]]                  
                else:
                    tipe = [tipes.split('.')[1].replace('()', '')]
                    
                key = repeatnames
                repeated = True
                
            
            else:
            
                tipe = objtype
        
          
        tipe = tipe[0]
        key = key[0]
        
        data = field_format.format(index=field[0], tipe=tipe, key=key)
        
        if repeated:
            data = 'repeated ' + data
        
        yield data



def parse_text(text: List[str]):
    
    name = text[0]
    
    fields = list(extract_field(text[1]))
    
    fields = '\n\t'.join(fields)
    fields = '\t' + fields
    
    buff_proto = f"message {name} " + "{\n\n" + fields + "\n\n}"
    
    
    return buff_proto



with open('protos/websockets.proto', 'w+') as out:
    out.write('syntax ="proto3"; \n\n\n')
    
    for text in hasil:
        
        data = parse_text(text)
        
        out.write(data)
        out.write('\n\n')
        print(data)
        print('--------')

    

    


    
    
    