import requests
import json

votes = "https://canon-4.api.network.umee.cc/umee/oracle/v1/validators/aggregate_votes"

def get_accepted_denoms():
    r = requests.get(url="https://canon-4.api.network.umee.cc/umee/oracle/v1/params")
    data = r.json()
    return data['params']['accept_list']

denoms_resp=get_accepted_denoms()
acceped_denoms = list()
for i in range(0,len(denoms_resp)):
    acceped_denoms.append(denoms_resp[i]['symbol_denom'].upper())
acceped_denoms = list(set(acceped_denoms))         

def get_validators():
    r = requests.get(url="https://canon-4.api.network.umee.cc/cosmos/staking/v1beta1/validators?status=BOND_STATUS_BONDED")
    data = r.json()
    return data['validators']

validators = get_validators()
valaddrs = list()

for i in range(0,len(validators)):
    valaddrs.append(validators[i]['operator_address'])

print("Accepted Denoms ",len(acceped_denoms),acceped_denoms)
print("Total Bonded validators ",len(valaddrs))

def get_req(url):
    r = requests.get(url=url)
    data = r.json()
    if len(data['aggregate_votes']) > 0:
        return data 
    return get_req(url)

def write_data(data):
    with open('data.json', 'w') as f:
        json.dump(data, f)

def get_moniker(valaddr):
    for i in range(0,len(validators)):
        if validators[i]['operator_address'] == valaddr:
            return validators[i]['description']['moniker']

# Get the total denoms exists in app
votes_resp = get_req(votes)
aggregate_votes = votes_resp['aggregate_votes']
votes_info = list()
voted_validators = list()
print("Total Validators ",len(aggregate_votes))
for i in range(0, len(aggregate_votes)):
    a=dict()
    a['voter'] = aggregate_votes[i]['voter']
    voted_validators.append( a['voter'])
    d = aggregate_votes[i]['exchange_rate_tuples']
    exchange_rate_tuples = []
    for k in range(0,len(d)):
        exchange_rate_tuples.append(d[k]['denom'])

    if len(exchange_rate_tuples) != len(acceped_denoms):
        print("❌" ,a['voter'], ">>", get_moniker(a['voter'])," is missing the votes for denoms")
        for j in range(0,len(acceped_denoms)):
            if exchange_rate_tuples.count(acceped_denoms[j])==0 :
                print("missing votes for denom : ",acceped_denoms[j])
                a['missing_votes'] =acceped_denoms[j]
    else:
        print("✅" ,a['voter'], ">>", get_moniker(a['voter']), "is submit the correct votes ",len(exchange_rate_tuples))
    print("*"*100)

print("*"*100)
print("Below validators does not submit the votes at all")
for i in range(0,len(valaddrs)):
    if voted_validators.count(valaddrs[i]) == 0 :
        print("❌" ,valaddrs[i], ">>", get_moniker(valaddrs[i])," is does not submit the votes for denoms")