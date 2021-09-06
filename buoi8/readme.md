#1 Về bài tập So sánh tốc độ tìm kiếm của ES và MySql
- Số row trong DB là 119905 dòng (Do em tải dữ liệu lên db bị chậm nên không up lên hết mà chỉ lấy khoảng đó để so sánh) trong ES là 124569
- Lúc em chạy để up dữ liệu lên thì máy em bị fulldisk do em dùng hdd, em không biết cái này có ảnh hưởng đến tốc độ không,
#2 Demo so sánh
case: "that great"
    thoi gian tim kiem "that great" voi mySql: 1.0830454s
    thoi gian tim kiem "that great" voi ES: 46.5438ms
case: "adsfjlalkdgjlakdfaldkfl"
    thoi gian tim kiem "adsfjlalkdgjlakdfaldkfl" voi mySql: 6.0852566s
    thoi gian tim kiem "adsfjlalkdgjlakdfaldkfl" voi ES: 76.077ms
case: "nice"
    thoi gian tim kiem "nice" voi mySql: 116.3036ms
    thoi gian tim kiem "nice" voi ES: 78.4595ms