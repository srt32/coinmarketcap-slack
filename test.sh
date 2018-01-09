up

curl `up token`
# equals: url is required

curl -d 'command=/coin&user_name=tj' `up token`
# equals: url is required

curl -d 'command=/coin&user_name=simon&text=spankchain' `up token`
# contains: Status 200
