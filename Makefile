# คำสั่งรันแบบปกติ


run:
	cd backendluckynumber && go run main.go

# คำสั่งทดสอบโค้ด
test:
	cd backendluckynumber && cd handler && go test -v

front:
	cd lucky-number && npm run dev

