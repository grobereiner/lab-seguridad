from flask import Flask, request
import pandas


app = Flask(__name__)

app.config['DEBUG'] = True


df = pandas.read_csv('sample_data.csv')
token = 'sdagkag'

@app.route('/get-token')
def get_token():
    auth_header = request.headers.get('Authorization')
    if auth_header and auth_header.startswith('Bearer '):
        bearer_token = auth_header.split(' ')[1]

        if bearer_token != token:
            return 'Invalid token', 401

        
        q = request.args.get('q')
        
        if not q:
            return 'Missing query', 401

        result_df = df.loc[df['username'] == q, ['company_name', 'payment']]


        return result_df.to_json(orient='records')
    else:
        return 'No Bearer token found in the request', 401
    
app.run('localhost', 5000)