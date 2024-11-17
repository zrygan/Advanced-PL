require 'json'
require 'http'

# Sample data for your pie chart using the special hash
special = {
  "Category A" => 30,
  "Category B" => 15,
  "Category C" => 25,
  "Category D" => 10,
  "Category E" => 20
}

# Extract labels and data from the hash
labels = special.keys
data = special.values

# Define the chart configuration for a pie chart
chart_config = {
  type: 'pie',
  data: {
    labels: labels,
    datasets: [{
      data: data,
      backgroundColor: [
        '#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0', '#9966FF'
      ],
      borderColor: '#ffffff',
      borderWidth: 2
    }]
  },
  options: {
    plugins: {
      legend: {
        display: true,
        position: 'right'
      }
    }
  }
}

# Set up the request parameters for the QuickChart API
request_payload = {
  chart: chart_config.to_json,  # Convert the chart configuration to JSON string
  width: 800,                   # Set the width of the output image
  height: 600,                  # Set the height of the output image
  devicePixelRatio: 2,          # Use a higher device pixel ratio for better quality
  backgroundColor: 'white',     # Set background color to white
  format: 'png',                # Specify the output format as PNG
  version: '4'                  # Use Chart.js version 4 for latest features
}

# Convert the setup to JSON format for the API request
json_payload = request_payload.to_json

# Make the POST request to the QuickChart API
chart_url = 'https://quickchart.io/chart'
response = HTTP.post(chart_url, headers: { 'Content-Type' => 'application/json' }, body: json_payload)

# Check if the request was successful and save the image
if response.status.success?
  File.open('ruby-pie-chart.png', 'wb') do |file|
    file.write(response.body)
    puts "Pie chart saved as 'ruby-pie-chart.png'"
  end
else
  puts "Error: #{response.status} - #{response.body.to_s}"
end
