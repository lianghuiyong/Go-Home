# 抓包记录

## 登陆

1、登陆初始化

    https://kyfw.12306.cn/otn/login/init

2、验证码

    https://kyfw.12306.cn/otn/passcodeNew/getPassCodeNew.do?module=login&rand=sjrand

    //验证码校验地址
    //POST:
    //    randCode //验证码
    //    rand     //sjrand
    //response:
    //    status:true
    //    httpstatus:200
    //    data:"Y"       //注意Y为正确的意思
    https://kyfw.12306.cn/otn/passcodeNew/checkRandCodeAnsyn


    //登陆验证码获取地址
    //GET
    https://kyfw.12306.cn/otn/passcodeNew/getPassCodeNew.do?module=login&rand=sjrand&0.27336162992543445

    //余票查询地址
    //GET:
    //    leftTicketDTO.train_date //乘车日期
    //    leftTicketDTO.from_station //出发城市
    //    leftTicketDTO.to_station  //开往城市
    //    purpose_codes  //乘车代码 ADULT 成人，0X00 学生
    https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=2014-01-28&leftTicketDTO.from_station=SHH&leftTicketDTO.to_station=XAY&purpose_codes=ADULT


    //校验用户登陆状态
    //POST
    https://kyfw.12306.cn/otn/login/checkUser

    //用户登陆
    //POST:
    //    loginUserDTO.user_name  //用户名
    //    userDTO.password //密码
    https://kyfw.12306.cn/otn/login/loginUserAsyn?random=1389261354627