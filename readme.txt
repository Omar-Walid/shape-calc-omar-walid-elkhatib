# Shape Calculator Application

This application allows you to calculate properties (such as area and perimeter) for various geometric shapes using JSON input.

## How to Use

1. **Send a POST request** with your shape data in JSON format to the following destination address:

    **Destination Address:**  
    `http://localhost:8080/api/shape/calculate`

2. **JSON Input Format:**  
    The JSON should specify the shape type and its parameters.

## Examples of Valid JSON Data

### 1. Rectangle
##### {"shape":"rectangle",
       "length":10,
       "width":5}

### 2. Circle
##### {"shape":"circle",
       "radius":4.5}

### 3. Triangle
##### {"shape":"triangle",
       "base":6,
       "height":6,
       "side1":8.49,
       "side2":6,
       "side3":6}

                ##### IMPORTANT #####
In Triangle, All sides must be an actual triangle;
else it will return error : triangle does not equate, meaning
the provided sides' lengths cannot be a triangle.

Output will provide the type of shape, the Area in float value with a 2 digit approximate
and the perimeter in float value with a 2 digit approximate.