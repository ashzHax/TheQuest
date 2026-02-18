# TheQuest
A website to input code and get answers regard to the code. Like a "quest".


# Backend Build
cd backend/
go build -o app
./app

# Frontend Build
cd frontend/
npm run build
rm ../backend/dist -rf
cp dist ../backend/