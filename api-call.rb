require 'net/http'
require 'json'

# Fetch data
url = URI("https://jsonplaceholder.typicode.com/posts")
response = Net::HTTP.get(url)
posts = JSON.parse(response)

# Filter posts for userId == 1
user_posts = posts.select { |post| post["userId"] == 1 }

# Print titles
user_posts.each do |post|
  puts post["title"]
end

