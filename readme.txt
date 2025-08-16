Shape Calculator Application

This application calculates area and perimeter for various geometric shapes using JSON input.

Technologies Used

Go (Golang)

Gin web framework

Go testing package for unit tests

Usage

Send a POST request with shape data in JSON format to:

http://localhost:8080/api/shape/calculate


The JSON must specify the shape type and its parameters.

Example Requests
Rectangle
{
  "shape": "rectangle",
  "length": 10,
  "width": 5
}

Circle
{
  "shape": "circle",
  "radius": 4.5
}

Triangle
{
  "shape": "triangle",
  "base": 6,
  "height": 6,
  "side1": 8.49,
  "side2": 6,
  "side3": 6
}


⚠️ Important (Triangles):

All three sides must form a valid triangle.

If not, the API will return an error:

{ "error": "Triangle sides do not form a valid triangle." }

Output

The response includes:

Shape type

Area (float, rounded to 2 decimal places)

Perimeter (float, rounded to 2 decimal places)

How to Run Locally

Clone the repository:

git clone https://github.com/your-username/shape-calculator.git
cd shape-calculator


Install dependencies:

go mod tidy


Run the application:

go run main.go


Run tests:

go test ./...

Example cURL Requests
Rectangle
curl -X POST http://localhost:8080/api/shape/calculate \
  -H "Content-Type: application/json" \
  -d '{"shape":"rectangle","length":10,"width":5}'


✅ Example Response:

{
  "shape": "rectangle",
  "area": 50.00,
  "perimeter": 30.00
}

Circle
curl -X POST http://localhost:8080/api/shape/calculate \
  -H "Content-Type: application/json" \
  -d '{"shape":"circle","radius":4.5}'


✅ Example Response:

{
  "shape": "circle",
  "area": 63.62,
  "perimeter": 28.27
}