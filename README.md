## CODING TEST
    - Author: Nguyễn Tạ Huy Hoàng
    - GitHub: https://github.com/NguyenTaHuyHoang
    - Link: https://testing-api-lac.vercel.app/
  
## Technologies Used
    - Go (Golang): Go is a statically typed, compiled language known for its performance and simplicity, serving as the foundation for this project.
    - gin: Gin is a library that facilitates API development for this project.
    - gorm: gorm is a library that enables database connectivity to Postgresql and data handling.
    - Nextjs: Next.js: A React framework used for building fast.
    - Axios: A promise-based HTTP client for making asynchronous API requests, simplifying data fetching and response handling in the application.
    - Postman: Use Postman to test API.
    - Database: Neon (Cloud postgresql).

## Usage
### I used cloud postgresql and deployed the application. If you want to run on your localmachine please change the following information:
- In the page.tsx file, change the url in the await axios.get command to 'http://localhost:8081/kols'
- In the main.go file, change the url inside AllowOrigins to 'http://localhost:3000'
- Change DB_URL inside the .env file to match your database

### BACKEND Part
1. Clone the repository:
```
https://github.com/NguyenTaHuyHoang/TestingAPI.git
```

2. Navigate to BE's main folder and Install dependencies (optional if the vendor folder is included):
```
go get -u all
```

3. Start the server:

```
go run ./main.go
```

### FRONTEND Part
1. Navigate to FE's main folder and Run this command
```
npm i
npm run dev
```

2. Start the server:
```
Navigate to localhost:3000 
```

## Contact

If you have any questions or suggestions, please feel free to contact us at nthh01082002@gmail.com.
