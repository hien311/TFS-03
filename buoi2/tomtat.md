## Các luật viết trong Go
### MixedCaps
- Trong Go thì sử dụng một quy tắc có tên gọi là mixedCaps (thường được gọi là camelCase) thay vì gạch dướng để viết tên nhiều từ. Quy tắc đó là chữ cái đầu sẽ viết thường, sau đó những chữ cái đầu tiên của mỗi từ sau đó sẽ được viết hoa. Nếu muốn sử dựng một hàm hay biến ở ngoài package, ký tự đầu tiên của  nó phải là chữ viết hgoa.
### Interface names
- Theo quy ước, một interface được đặt tên bằng tên phương thức cộng với hậu tố "er" hoặc được sửa đổi sao cho đúng với ngữ pháp.
- Tuy nhiên một interface có thể có nhiều phương thức => việc đặt tên theo quy ước sẽ không phải lúc nào cũng rõ ràng => việc chia interface ra nhiều interfaces khác. Tuy nhiên việc chia ra cũng còn tùy vào từng trường hợp.
### getters 
- Việc tự cung cấp getters và setters thường là phù hợp nhưng lại không cần thiếu
## Quy tắc chưa được viết trong Go
_ Có những quy tắc không được viết ra nhưng nó lại phổ biến trong cộng đồng _
### Shorter variable names
** Chương trình được viết ra để con người nhìn vào có thể hiểu được **
- Single-letter identifier: thường là biến được đặt trong một phạm vi hạn chế (ví dụ như là biến index của vòng lặp ...) 
- shorthand name: được khuyến nghị là nên rút ngắn nhất có thể nhưng vẫn đảm bảo để bất kì ai có thể hiểu được
### unique names
- Các từ như API, HTTP hoặc như ID, DB. Chúng ta sẽ giữ nó như vậy thay vì đưa về định dạng gốc (VD: userID thay vì userId)
### line length
- Tránh viết những dòng quá dài
