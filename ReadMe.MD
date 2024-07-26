# Coding Challenge for a company


<h2>Solution</h2>

Please refer to Solution.docx for the detailed solution

<h2> Running the Sample </h2>

In mac:
./order-middleware

In Windows:
Run order-middleware.exe


Test:
--

GET
curl -u admin:password http://localhost:8080/orders?customer_id=10001


POST
curl -X POST http://localhost:8080/orders \
-u admin:password \
-H "Content-Type: application/json" \
-d '{
"customer_id": "10001",
"items": [
{
"item_number": "item-001",
"quantity": 2,
"price": 19.99
},
{
"item_number": "item-002",
"quantity": 1,
"price": 9.99
}
],
}'
