.PHONY: gen-frontend
gen-frontend:
cd .\app\frontend\ && cwgo server --type HTTP --idl ..\..\idl\frontend\home.proto --service frontend -module github.com/hourhl/Qmall/app/frontend -I ..\..\idl\