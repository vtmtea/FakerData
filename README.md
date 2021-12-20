# FakerData
Use golang generate faker data

Usage
```go
fakerEngine := faker.New()
//optional valid locale: ["en_US", "zh_CN"]
//default: zh_CN
fakerEngine.setLocale("en_US")
```

Output Random Name
```go
fakerEngine.Name() //游超
```

Output Random Phone Number
```go
fakerEngine.PhoneNumber() //17836850955
```