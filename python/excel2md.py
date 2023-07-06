import pandas as pd 
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('input_file')
parser.add_argument('output_file')
args = parser.parse_args()

engine = 'python' 

if args.input_file.endswith('.csv'):
    df = pd.read_csv(args.input_file,engine=engine)  
elif args.input_file.endswith('.xlsx'):
    df = pd.read_excel(args.input_file)  
else:
    print('unsupported file format')
    exit()

md = df.to_markdown()
with open(args.output_file, 'w') as f:
    f.write(md)


### python excel2md.py [文件].xlsx或者 [文件].csv  [输出文件].md