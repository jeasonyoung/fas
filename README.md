# fas
家庭账务管理系统(Family Accounting System)


# Test

curl -i -X POST -H "content-type:application/json" --data '{"head":{"version":1,"channel":100,"mac":"123456","token":"","time":1519391776000,"sign":"tesfdasfdasfdafdadfadfa"},"body":{"account":"admin","password":"123456"}}' http://127.0.0.1:8000/api/v1/common/signin
