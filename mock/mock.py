from flask import Flask, request
import pandas

app = Flask(__name__)

app.config['DEBUG'] = True

df = pandas.read_csv('sample_data.csv')

@app.route('/get-token')
def get_token():
    auth_header = request.headers.get('Authorization')
    if auth_header and auth_header.startswith('Bearer '):
        bearer_token = auth_header.split(' ')[1]

        result_df = df.loc[df['username'] == bearer_token, ['company_name', 'payment']]


        return result_df.to_json(orient='records')
        return f'Bearer token: {bearer_token}', 200
    else:
        return 'No Bearer token found in the request', 400
    
app.run('localhost', 5000)