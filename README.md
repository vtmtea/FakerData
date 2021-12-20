# FakerData
Use golang generate faker data

Usage
```go
faker := faker.New()
//optional valid locale: ["en_US", "zh_CN"]
//default: zh_CN
faker.setLocale("en_US")
```

Output Random Name
```go
faker.Name() //游超
```

Output Random Phone Number
```go
faker.PhoneNumber() //17836850955
```